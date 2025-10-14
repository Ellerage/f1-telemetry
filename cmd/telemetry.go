package main

import (
	telemetryhandler "f1-telemetry/internal/telemetry_handler"
	"net"
)

func main() {
	telemetryServer := telemetryhandler.NewTelemetryServer(telemetryhandler.TelemetryUDPServer{
		Addr: net.UDPAddr{
			IP:   net.ParseIP("0.0.0.0"),
			Port: 20778,
		},
	})

	closeTSConn := telemetryServer.CreateConnection()
	defer closeTSConn()

	// Blocking operations
	telemetryServer.RegisterHandler()
}
