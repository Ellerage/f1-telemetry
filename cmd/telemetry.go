package main

import (
	"f1-telemetry/internal/db"
	filemanager "f1-telemetry/internal/file_manager"
	hotkeyhandler "f1-telemetry/internal/hotkey_handler"
	"f1-telemetry/internal/repository"
	"f1-telemetry/internal/service"
	sessionstorage "f1-telemetry/internal/session"
	telemetryhandler "f1-telemetry/internal/telemetry_handler"
	"net"
)

func main() {
	dbProvider := db.NewDBProvider()
	conn, closeDBConn, err := dbProvider.DBConnect()
	defer closeDBConn()

	if err != nil {
		panic(err)
	}

	lapsFM := filemanager.NewFileManager()
	lapsFM.OpenFile("laps.csv")

	telemetryFM := filemanager.NewFileManager()
	telemetryFM.OpenFile("telemetry.csv")

	lapService := service.NewLapService(service.LapServiceParams{
		LapRepository:  repository.NewLapRepository(conn),
		LapFileManager: lapsFM,
	})

	telemetryService := service.NewTelemetryService(service.TelemetryServiceParams{
		TelemetryRepository:  repository.NewLapRepository(conn),
		TelemetryFileManager: lapsFM,
	})

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

	go telemetryServer.RegisterHandler()

	select {}
}
