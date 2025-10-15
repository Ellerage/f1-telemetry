package model

type TelemetryRow struct {
	// Header
	PacketId       uint8
	SessionTime    float32
	SessionUID     uint64
	PlayerCarIndex uint8

	// Car telemetry
	Speed    uint16
	Throttle float32
	Brake    float32
	Steer    float32

	// Lap data
	CurrentLapNum      uint8
	CurrentLapTimeInMS uint32
	LapDistance        float32 // current lap in m

	// Session
	TrackLength uint16
	SessionType uint8
}

var TelemetryRowColumns = []string{
	"PacketId",
	"SessionTime",
	"SessionUID",
	"PlayerCarIndex",
	"Speed",
	"Throttle",
	"Brake",
	"Steer",
	"CurrentLapNum",
	"CurrentLapTimeInMS",
	"LapDistance",
	"TrackLength",
	"SessionType",
}
