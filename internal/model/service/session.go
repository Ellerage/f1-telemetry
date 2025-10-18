package model

type SessionStorageUpdate struct {
	CurrentLapNum      *uint8
	CurrentLapTimeInMS *uint32
	LapDistance        *float32
	TrackLength        *uint16
	SessionType        *uint8

	TrackId *int8

	LastLapTimeInMS *uint32
}
