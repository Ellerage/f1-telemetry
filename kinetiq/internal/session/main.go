package sessionstorage

import model "f1-telemetry/internal/model/service"

type SessionStorage struct {
	CurrentLapNum      uint8
	CurrentLapTimeInMS uint32
	LapDistance        float32
	TrackLength        uint16
	SessionType        uint8

	TrackId int8

	LastLapTimeInMS uint32
}

func NewSessionStorage() *SessionStorage {
	return &SessionStorage{}
}

func (s *SessionStorage) GetCurrentLapNum() uint8 {
	return s.CurrentLapNum
}
func (s *SessionStorage) GetCurrentLapTimeInMS() uint32 {
	return s.CurrentLapTimeInMS
}
func (s *SessionStorage) GetLapDistance() float32 {
	return s.LapDistance
}
func (s *SessionStorage) GetTrackLength() uint16 {
	return s.TrackLength
}
func (s *SessionStorage) GetSessionType() uint8 {
	return s.SessionType
}
func (s *SessionStorage) GetTrackId() int8 {
	return s.TrackId
}
func (s *SessionStorage) GetLastLapTimeInMS() uint32 {
	return s.LastLapTimeInMS
}

func (s *SessionStorage) Update(toUpdate model.SessionStorageUpdate) {
	if toUpdate.CurrentLapNum != nil {
		s.CurrentLapNum = *toUpdate.CurrentLapNum
	}
	if toUpdate.CurrentLapTimeInMS != nil {
		s.CurrentLapTimeInMS = *toUpdate.CurrentLapTimeInMS
	}
	if toUpdate.LapDistance != nil {
		s.LapDistance = *toUpdate.LapDistance
	}
	if toUpdate.TrackLength != nil {
		s.TrackLength = *toUpdate.TrackLength
	}
	if toUpdate.SessionType != nil {
		s.SessionType = *toUpdate.SessionType
	}
	if toUpdate.TrackId != nil {
		s.TrackId = *toUpdate.TrackId
	}
	if toUpdate.LastLapTimeInMS != nil {
		s.LastLapTimeInMS = *toUpdate.LastLapTimeInMS
	}
}
