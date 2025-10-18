package main

import (
	"f1-telemetry/internal/db"
	filemanager "f1-telemetry/internal/file_manager"
	model "f1-telemetry/internal/model/csv"
	"f1-telemetry/internal/obs"
	"f1-telemetry/internal/repository"
	"f1-telemetry/internal/service"
	sessionstorage "f1-telemetry/internal/session"
	telemetryhandler "f1-telemetry/internal/telemetry_handler"
	"log/slog"
	"net"
)

const (
	TELEMETRY_FILE_NAME = "telemetry.csv"
	LAPS_FILE_NAME      = "laps.csv"
	// OBS
	OBS_BUFFER_SECONDS = 120
	OBS_PASSWORD       = "123456"
	OBS_PORT           = 4455
	OBS_ADDR           = "localhost"

	// Telemetry file manger
	TELEMETRY_FM_BUFFER_ROWS = 100

	// UDP Server
	UDP_PORT = 20778
)

func main() {
	dbProvider := db.NewDBProvider()
	conn, closeDBConn, err := dbProvider.DBConnect()
	defer closeDBConn()

	if err != nil {
		panic(err)
	}

	lapsFM := filemanager.NewFileManager(model.LapRowColumn)
	if _, err := lapsFM.OpenFile(LAPS_FILE_NAME); err != nil {
		slog.Error(err.Error())
	}

	telemetryFM := filemanager.NewFileManager(model.TelemetryRowColumns)
	bufferedTelemetryFM := filemanager.NewBufferFileManger(filemanager.BufferFileMangerParams{
		FileManager: telemetryFM,
		BufferSize:  TELEMETRY_FM_BUFFER_ROWS,
	})
	if _, err := telemetryFM.OpenFile(TELEMETRY_FILE_NAME); err != nil {
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

	obsApi := obs.NewOBSService(OBS_ADDR, OBS_PORT, OBS_PASSWORD, OBS_BUFFER_SECONDS)
	if err := obsApi.Connect(); err != nil {
		slog.Error(err.Error())
		return
	}

	if err := obsApi.StartReplayBuffer(); err != nil {
		slog.Error(err.Error())
		return
	}

	telemetryServer := telemetryhandler.NewTelemetryServer(telemetryhandler.TelemetryUDPServerParams{
		Addr: net.UDPAddr{
			IP:   net.ParseIP("0.0.0.0"),
			Port: UDP_PORT,
		},
		LapService:       lapService,
		TelemetryService: telemetryService,
		SessionStorage:   sessionstorage.NewSessionStorage(),
		ObsApi:           obsApi,
	})

	closeTSConn := telemetryServer.CreateConnection()
	defer closeTSConn()

	go telemetryServer.RegisterHandler()

	// Hotkey handler for quick commands
	// hkHandler := hotkeyhandler.NewHotkeyHandler(
	// 	hotkeyhandler.HotkeyHandlerParams{SessionStorage: sesstionStorage},
	// )
	// hkHandler.RegisterHotkeyListener()

	// select {}
}
