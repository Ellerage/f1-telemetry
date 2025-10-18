package telemetryhandler

import (
	modelcsv "f1-telemetry/internal/model/csv"
	packets "f1-telemetry/internal/model/packets"
	modelservice "f1-telemetry/internal/model/service"
	"f1-telemetry/internal/utils"
	"fmt"
	"log/slog"
	"net"
)

type LapService interface {
	Create(toCreate modelcsv.LapRow) error
}

type TelemetryService interface {
	Create(toCreate modelcsv.TelemetryRow) error
}

type SessionStorage interface {
	Update(toUpdate modelservice.SessionStorageUpdate)
	GetCurrentLapNum() uint8
	GetCurrentLapTimeInMS() uint32
	GetLapDistance() float32
	GetTrackLength() uint16
	GetSessionType() uint8
	GetTrackId() int8
	GetLastLapTimeInMS() uint32
}

type TelemetryUDPServerParams struct {
	Addr             net.UDPAddr
	LapService       LapService
	TelemetryService TelemetryService
	SessionStorage   SessionStorage
}

type TelemetryUDPServer struct {
	conn             *net.UDPConn
	addr             net.UDPAddr
	lapService       LapService
	telemetryService TelemetryService
	sessionStorage   SessionStorage
}

func NewTelemetryServer(params TelemetryUDPServerParams) *TelemetryUDPServer {
	return &TelemetryUDPServer{
		addr:             params.Addr,
		lapService:       params.LapService,
		telemetryService: params.TelemetryService,
		sessionStorage:   params.SessionStorage,
	}
}

func (ts *TelemetryUDPServer) CreateConnection() func() error {
	conn, err := net.ListenUDP("udp", &ts.addr)
	if err != nil {
		panic(err)
	}

	ts.conn = conn

	return conn.Close
}

// Blocking operation
func (ts *TelemetryUDPServer) RegisterHandler() {
	buf := make([]byte, 1500)

	for {
		n, _, err := ts.conn.ReadFromUDP(buf)
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		data := buf[:n]

		pkt, err, ok := ParsePacket(data)
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		if !ok {
			continue
		}

		switch p := pkt.(type) {
		case packets.SessionPacket:
			fmt.Println("[Session] Packet received")

			ts.sessionStorage.Update(modelservice.SessionStorageUpdate{
				SessionType: &p.SessionType,
				TrackId:     &p.TrackId,
				TrackLength: &p.TrackLength,
			})

		case packets.LapPacket:
			fmt.Println("[LapData] Packet received")

			if ts.sessionStorage.GetCurrentLapNum() != p.LapData[p.Header.PlayerCarIndex].CurrentLapNum {
				Sector3Minutes, Sector3MS := utils.GetSector3(utils.GetSector3Params{
					LastLapTimeInMS:        p.LapData[p.Header.PlayerCarIndex].LastLapTimeInMS,
					Sector1TimeMSPart:      p.LapData[p.Header.PlayerCarIndex].Sector1TimeMSPart,
					Sector1TimeMinutesPart: p.LapData[p.Header.PlayerCarIndex].Sector1TimeMinutesPart,
					Sector2TimeMSPart:      p.LapData[p.Header.PlayerCarIndex].Sector2TimeMSPart,
					Sector2TimeMinutesPart: p.LapData[p.Header.PlayerCarIndex].Sector2TimeMinutesPart,
				})

				playerData := p.ToLapRow(packets.LapRowSessionData{
					Sector3Minutes: Sector3Minutes, // count
					Sector3MS:      Sector3MS,
					Total:          p.LapData[p.Header.PlayerCarIndex].LastLapTimeInMS,
					SessionType:    ts.sessionStorage.GetSessionType(),
					TrackId:        ts.sessionStorage.GetTrackId(),
				})

				ts.lapService.Create(playerData)
			}

			ts.sessionStorage.Update(modelservice.SessionStorageUpdate{
				CurrentLapNum:      &p.LapData[p.Header.PlayerCarIndex].CurrentLapNum,
				CurrentLapTimeInMS: &p.LapData[p.Header.PlayerCarIndex].CurrentLapTimeInMS,
				LapDistance:        &p.LapData[p.Header.PlayerCarIndex].LapDistance,
				LastLapTimeInMS:    &p.LapData[p.Header.PlayerCarIndex].LastLapTimeInMS,
			})

		case packets.CarTelemetryPacket:
			fmt.Println("[CarTelemetryPacket] Packet received")

			ts.telemetryService.Create(p.ToLapRow(packets.TelemetryRowSessionData{
				CurrentLapNum:      ts.sessionStorage.GetCurrentLapNum(),
				CurrentLapTimeInMS: ts.sessionStorage.GetCurrentLapTimeInMS(),
				LapDistance:        ts.sessionStorage.GetLapDistance(),
				TrackLength:        ts.sessionStorage.GetTrackLength(),
				SessionType:        ts.sessionStorage.GetSessionType(),
				TrackId:            ts.sessionStorage.GetTrackId(),
			}))
		default:
			fmt.Println("[Unknown] Packet received")
		}
	}
}
