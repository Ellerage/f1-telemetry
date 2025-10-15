package model

type TelemetryRow struct {
	// Header
	PacketId    uint8
	SessionTime float32

	// Car telemetry
	Speed    uint16
	Throttle float32
	Steer    float32
	Brake    float32
	// Gear     int8

	// Lap data
	CurrentLapNum      uint8
	CurrentLapTimeInMS uint32
	LapDistance        float32
}

var TelemetryRowColumns = []string{"PacketId", "SessionTime", "Speed", "Throttle", "Steer", "Brake", "CurrentLapNum", "CurrentLapTimeInMS", "LapDistance"}
