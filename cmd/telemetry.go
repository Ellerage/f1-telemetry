package main

import (
	hotkeyhandler "f1-telemetry/internal/hotkey_handler"
	sessionstorage "f1-telemetry/internal/session"
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

	// Storage with current token
	sesstionStorage := sessionstorage.NewSessionStorage()

	// Hotkey handler for quick commands
	hkHandler := hotkeyhandler.NewHotkeyHandler(
		hotkeyhandler.HotkeyHandlerParams{SessionStorage: sesstionStorage},
	)
	hkHandler.RegisterHotkeyListener()

	// Blocking operations
	go telemetryServer.RegisterHandler()

	select {}
}
