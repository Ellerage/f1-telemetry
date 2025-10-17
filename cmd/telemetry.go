package main

import (
	"f1-telemetry/internal/db"
	filemanager "f1-telemetry/internal/file_manager"
	hotkeyhandler "f1-telemetry/internal/hotkey_handler"
	"f1-telemetry/internal/repository"
	"f1-telemetry/internal/service"
	sessionstorage "f1-telemetry/internal/session"
	telemetryhandler "f1-telemetry/internal/telemetry_handler"
	"log/slog"
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
	if _, err := lapsFM.OpenFile("laps.csv"); err != nil {
		slog.Error(err.Error())
	}

	telemetryFM := filemanager.NewFileManager()
	bufferedTelemetryFM := filemanager.NewBufferFileManger(filemanager.BufferFileMangerParams{
		FileManager: telemetryFM,
		BufferSize:  100,
	})
	if _, err := telemetryFM.OpenFile("telemetry.csv"); err != nil {
		slog.Error(err.Error())
	}

	lapService := service.NewLapService(service.LapServiceParams{
		LapRepository: repository.NewLapRepository(repository.LapRepositoryParams{
			DB:         conn,
			FileManger: lapsFM,
		}),
		LapFileManager: lapsFM,
	})

	repository.NewTelemetryRepository(repository.TelemetryRepositoryParams{
		DB:         conn,
		FileManger: bufferedTelemetryFM,
	})

	telemetryService := service.NewTelemetryService(service.TelemetryServiceParams{
		TelemetryRepository: repository.NewTelemetryRepository(repository.TelemetryRepositoryParams{
			DB:         conn,
			FileManger: bufferedTelemetryFM,
		}),
		TelemetryFileManager: lapsFM,
	})

	telemetryServer := telemetryhandler.NewTelemetryServer(telemetryhandler.TelemetryUDPServerParams{
		Addr: net.UDPAddr{
			IP:   net.ParseIP("0.0.0.0"),
			Port: 20778,
		},
		LapService:       lapService,
		TelemetryService: telemetryService,
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
