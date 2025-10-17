package service

import (
	model "f1-telemetry/internal/model/csv"
)

type TelemetryRepository interface {
	GetAll() ([]model.LapRow, error)
}

type TelemetryServiceParams struct {
	TelemetryRepository  TelemetryRepository
	TelemetryFileManager FileManager
}

type TelemetryService struct {
	telemetryRepository  TelemetryRepository
	telemetryFileManager FileManager
}

func NewTelemetryService(params TelemetryServiceParams) *LapService {
	return &LapService{lapRepository: params.TelemetryRepository, lapFileManager: params.TelemetryFileManager}
}
