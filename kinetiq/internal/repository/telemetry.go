package repository

import (
	"database/sql"
	model "f1-telemetry/internal/model/csv"
)

type TelemetryBFileManager interface {
	AddRow([]string) error
}

type TelemetryRepositoryParams struct {
	DB         *sql.DB
	FileManger TelemetryBFileManager
}

type TelemetryRepository struct {
	db         *sql.DB
	fileManger TelemetryBFileManager
}

func NewTelemetryRepository(params TelemetryRepositoryParams) *TelemetryRepository {
	return &TelemetryRepository{db: params.DB, fileManger: params.FileManger}
}

func (repo *TelemetryRepository) Create(toCreate model.TelemetryRow) error {
	return repo.fileManger.AddRow(toCreate.FormatToRow())
}
