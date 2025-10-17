package repository

import (
	"database/sql"
	model "f1-telemetry/internal/model/csv"
	"fmt"
	"strings"
)

type LapRepository struct {
	db *sql.DB
}

func NewLapRepository(db *sql.DB) *LapRepository {
	return &LapRepository{db: db}
}

func (r *LapRepository) GetAll(filters model.LapFilters) ([]model.LapRow, error) {

	var query strings.Builder
	query.WriteString("SELECT * FROM read_csv_auto('laps.csv')")

	var conditions []string

	if filters.SessionUID != nil {
		conditions = append(conditions, fmt.Sprintf("SessionUID = %d", *filters.SessionUID))
	}

	if filters.TrackId != nil {
		conditions = append(conditions, fmt.Sprintf("TrackId = %d", *filters.TrackId))
	}

	if filters.SessionType != nil {
		conditions = append(conditions, fmt.Sprintf("SessionType = %d", *filters.SessionType))
	}

	if len(conditions) > 0 {
		query.WriteString(" WHERE ")
		query.WriteString(strings.Join(conditions, " AND "))
	}

	rows, err := r.db.Query(query.String())

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var result []model.LapRow

	var Lap model.LapRow
	for rows.Next() {
		rows.Scan(
			&Lap.SessionUID,
			&Lap.PlayerCarIndex,
			&Lap.CurrentLapNum,
			&Lap.Sector1Minutes,
			&Lap.Sector1MS,
			&Lap.Sector2Minutes,
			&Lap.Sector2MS,
			&Lap.Sector3Minutes,
			&Lap.Sector3MS,
			&Lap.TotalMinutes,
			&Lap.TotalMS,
			&Lap.CurrentLapInvalid,
			&Lap.SessionType,
			&Lap.TrackId,
		)

		result = append(result, Lap)
	}

	return result, nil
}

func (r *LapRepository) GetBySessionIdLapNum(sessionId uint64, lapNum uint8) (model.LapRow, error) {
	row := r.db.QueryRow(
		"SELECT * FROM read_csv_auto('laps.csv') WHERE SessionUID = ? AND CurrentLapNum = ?",
		sessionId, lapNum,
	)

	var lap model.LapRow
	row.Scan(
		&lap.SessionUID,
		&lap.PlayerCarIndex,
		&lap.CurrentLapNum,
		&lap.Sector1Minutes,
		&lap.Sector1MS,
		&lap.Sector2Minutes,
		&lap.Sector2MS,
		&lap.Sector3Minutes,
		&lap.Sector3MS,
		&lap.TotalMinutes,
		&lap.TotalMS,
		&lap.CurrentLapInvalid,
		&lap.SessionType,
		&lap.TrackId,
	)

	return lap, row.Err()
}
