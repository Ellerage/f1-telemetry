package telemetryhandler

import (
	packets "f1-telemetry/internal/model/packets"
	"fmt"
	"log/slog"
	"net"
)

type FileManager interface {
}

type TelemetryUDPServer struct {
	conn        *net.UDPConn
	Addr        net.UDPAddr
	FileManager FileManager
}

func NewTelemetryServer(params TelemetryUDPServer) *TelemetryUDPServer {
	return &TelemetryUDPServer{
		Addr:        params.Addr,
		FileManager: params.FileManager,
	}
}

func (ts *TelemetryUDPServer) CreateConnection() func() error {
	conn, err := net.ListenUDP("udp", &ts.Addr)
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

		switch pkt.(type) {
		case packets.SessionPacket:
			fmt.Println("[Session] Packet received")
		case packets.LapPacket:
			fmt.Println("[LapData] Packet received")
		case packets.CarTelemetryPacket:
			fmt.Println("[CarTelemetryPacket] Packet received")
		default:
			fmt.Println("[Unknown] Packet received")
		}
	}
}
