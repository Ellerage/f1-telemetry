package telemetryhandler

import (
	model "f1-telemetry/internal/model/csv"
	packets "f1-telemetry/internal/model/packets"
	"fmt"
	"log/slog"
	"net"
)

type LapService interface {
	Create(toCreate model.LapRow) error
}

type TelemetryUDPServerParams struct {
	Addr       net.UDPAddr
	LapService LapService
}

type TelemetryUDPServer struct {
	conn       *net.UDPConn
	addr       net.UDPAddr
	lapService LapService
}

func NewTelemetryServer(params TelemetryUDPServerParams) *TelemetryUDPServer {
	return &TelemetryUDPServer{
		addr:       params.Addr,
		lapService: params.LapService,
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
		case packets.LapPacket:
			fmt.Println("[LapData] Packet received")
			ts.lapService.Create(p.ToLapRow())
		case packets.CarTelemetryPacket:
			fmt.Println("[CarTelemetryPacket] Packet received")
		default:
			fmt.Println("[Unknown] Packet received")
		}
	}
}
