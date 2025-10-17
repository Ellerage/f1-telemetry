package model

import "strconv"

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

func (r *TelemetryRow) FormatToRow() []string {
	return []string{
		strconv.FormatUint(uint64(r.PacketId), 10),
		strconv.FormatUint(uint64(r.SessionTime), 10),
		strconv.FormatUint(uint64(r.SessionUID), 10),
		strconv.FormatUint(uint64(r.PlayerCarIndex), 10),
		strconv.FormatUint(uint64(r.Speed), 10),
		strconv.FormatUint(uint64(r.Throttle), 10),
		strconv.FormatUint(uint64(r.Brake), 10),
		strconv.FormatUint(uint64(r.Steer), 10),
		strconv.FormatUint(uint64(r.CurrentLapNum), 10),
		strconv.FormatUint(uint64(r.CurrentLapTimeInMS), 10),
		strconv.FormatUint(uint64(r.LapDistance), 10),
		strconv.FormatUint(uint64(r.TrackLength), 10),
		strconv.FormatUint(uint64(r.SessionType), 10),
	}
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
