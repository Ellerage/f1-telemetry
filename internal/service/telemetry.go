package service

import (
	model "f1-telemetry/internal/model/csv"
)

type TelemetryRepository interface {
	GetAll(model.LapFilters) ([]model.LapRow, error)
}

type TelemetryServiceParams struct {
	TelemetryRepository  TelemetryRepository
	TelemetryFileManager FileManager
}

type TelemetryService struct {
	telemetryRepository  TelemetryRepository
	telemetryFileManager FileManager
}

func NewTelemetryService(params TelemetryServiceParams) *TelemetryService {
	return &TelemetryService{telemetryRepository: params.TelemetryRepository, telemetryFileManager: params.TelemetryFileManager}
}
