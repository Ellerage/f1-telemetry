package main

import (
	"f1-telemetry/internal/config"
	"f1-telemetry/internal/db"
	filemanager "f1-telemetry/internal/file_manager"
	model "f1-telemetry/internal/model/csv"
	"f1-telemetry/internal/obs"
	"f1-telemetry/internal/repository"
	apiserver "f1-telemetry/internal/server"
	"f1-telemetry/internal/service"
	sessionstorage "f1-telemetry/internal/session"
	telemetryhandler "f1-telemetry/internal/telemetry_handler"
	"log/slog"
	"net"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	dbProvider := db.NewDBProvider()
	conn, closeDBConn, err := dbProvider.DBConnect()
	defer closeDBConn()

	if err != nil {
		panic(err)
	}

	lapsFM := filemanager.NewFileManager(model.LapRowColumn)
	if _, err := lapsFM.OpenFile(cfg.LapsFileName); err != nil {
		slog.Error(err.Error())
	}

	telemetryFM := filemanager.NewFileManager(model.TelemetryRowColumns)
	bufferedTelemetryFM := filemanager.NewBufferFileManger(filemanager.BufferFileMangerParams{
		FileManager: telemetryFM,
		BufferSize:  int(cfg.TelemetryFMBufferRows),
	})
	if _, err := telemetryFM.OpenFile(cfg.TelemetryFileName); err != nil {
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

	obsApi := obs.NewOBSService(cfg.ObsAddr, int(cfg.ObsPort), cfg.ObsPassword, int(cfg.ObsBufferSeconds), cfg.UseObs == 1)
	if err := obsApi.Connect(); err != nil {
		slog.Error(err.Error())
	}

	if err := obsApi.StartReplayBuffer(); err != nil {
		slog.Error(err.Error())
	}

	telemetryServer := telemetryhandler.NewTelemetryServer(telemetryhandler.TelemetryUDPServerParams{
		Addr: net.UDPAddr{
			IP:   net.ParseIP("0.0.0.0"),
			Port: int(cfg.Port),
		},
		LapService:       lapService,
		TelemetryService: telemetryService,
		SessionStorage:   sessionstorage.NewSessionStorage(),
		ObsApi:           obsApi,
	})

	closeTSConn := telemetryServer.CreateConnection()
	defer closeTSConn()

	apiServer := apiserver.NewApiServer(apiserver.APIServerParams{
		LapService: lapService,
		Config:     cfg,
	})

	go apiServer.RegisterRouter()

	go telemetryServer.RegisterHandler()

	select {}
}
