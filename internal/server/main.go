package apiserver

import (
	"encoding/json"
	"f1-telemetry/internal/config"
	model "f1-telemetry/internal/model/csv"
	"net/http"
	"strconv"
)

type LapService interface {
	GetAll(filters model.LapFilters) ([]model.LapRow, error)
}

type Config interface {
	ChangeConfig(payload config.ConfigUpdate) config.Config
}

type APIServerParams struct {
	LapService LapService
	Config     Config
}

type APIServer struct {
	lapService LapService
	config     Config
}

func NewApiServer(params APIServerParams) *APIServer {
	return &APIServer{lapService: params.LapService, config: params.Config}
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

	http.HandleFunc("/api/config", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(s.config)
			return
		}

		if r.Method == http.MethodPost {
			var config config.ConfigUpdate
			err := json.NewDecoder(r.Body).Decode(&config)
			if err != nil {
				http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			newConfig := s.config.ChangeConfig(config)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newConfig)
			return
		}

	})

	http.ListenAndServe(":8080", nil)
}
