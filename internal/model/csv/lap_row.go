package model

type LapRow struct {
	// Header
	SessionUID     uint64
	PlayerCarIndex uint8

	// Lap data
	CurrentLapNum uint8

	Sector1Minutes uint8  // Sector 1 whole minute part
	Sector1MS      uint16 // Sector 1 time milliseconds part

	Sector2Minutes uint8  // Sector 2 whole minute part
	Sector2MS      uint16 // Sector 2 time milliseconds part

	Sector3Minutes uint8  // Sector 3 whole minute part
	Sector3MS      uint16 // Sector 3 time milliseconds part

	TotalMinutes uint8
	TotalMS      uint16

	CurrentLapInvalid uint8 // Current lap invalid - 0 = valid, 1 = invalid

	// Session
	SessionType uint8
	TrackId     int8
}

var LapRowColumn = []string{
	"SessionUID",
	"PlayerCarIndex",
	"CurrentLapNum",
	"Sector1Minutes",
	"Sector1MS",
	"Sector2Minutes",
	"Sector2MS",
	"Sector3Minutes",
	"Sector3MS",
	"TotalMinutes",
	"TotalMS",
	"CurrentLapInvalid",
	"SessionType",
	"TrackId",
}

type LapFilters struct {
	SessionUID  *uint64
	TrackId     *int8
	SessionType *uint8
}
