package apiserver

import (
	"encoding/json"
	model "f1-telemetry/internal/model/csv"
	"net/http"
	"strconv"
)

type LapService interface {
	GetAll(filters model.LapFilters) ([]model.LapRow, error)
}

type APIServerParams struct {
	LapService LapService
}

type APIServer struct {
	lapService LapService
}

func NewApiServer(params APIServerParams) *APIServer {
	return &APIServer{lapService: params.LapService}
}

func (s *APIServer) RegisterRouter() {
	http.HandleFunc("/api/laps", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		var filters model.LapFilters

		if v := q.Get("session_uid"); v != "" {
			if uid, err := strconv.ParseUint(v, 10, 64); err == nil {
				filters.SessionUID = &uid
			}
		}

		if v := q.Get("track_id"); v != "" {
			if id64, err := strconv.ParseInt(v, 10, 8); err == nil {
				id := int8(id64)
				filters.TrackId = &id
			}
		}

		if v := q.Get("session_type"); v != "" {
			if t64, err := strconv.ParseUint(v, 10, 8); err == nil {
				t := uint8(t64)
				filters.SessionType = &t
			}
		}

		laps, err := s.lapService.GetAll(filters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(laps)
	})

	http.ListenAndServe(":8080", nil)
}
