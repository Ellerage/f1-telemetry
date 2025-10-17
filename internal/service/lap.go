package service

import (
	model "f1-telemetry/internal/model/csv"
)

type LapRepository interface {
	GetAll(filters model.LapFilters) ([]model.LapRow, error)
	GetBySessionIdLapNum(sessionId uint64, lapNum uint8) (model.LapRow, error)
	Create(toCreate model.LapRow) error
}

type LapServiceParams struct {
	LapRepository  LapRepository
	LapFileManager FileManager
}

type LapService struct {
	lapRepository  LapRepository
	lapFileManager FileManager
}

func NewLapService(params LapServiceParams) *LapService {
	return &LapService{lapRepository: params.LapRepository, lapFileManager: params.LapFileManager}
}

func (s *LapService) GetAll(filters model.LapFilters) ([]model.LapRow, error) {
	return s.lapRepository.GetAll(filters)
}

func (s *LapService) GetBySessionIdLapNum(sessionId uint64, lapNum uint8) (model.LapRow, error) {
	return s.lapRepository.GetBySessionIdLapNum(sessionId, lapNum)
}

func (s *LapService) Create(toCreate model.LapRow) error {
	return s.lapRepository.Create(toCreate)
}
