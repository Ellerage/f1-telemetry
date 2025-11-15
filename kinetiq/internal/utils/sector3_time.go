package utils

type GetSector3Params struct {
	LastLapTimeInMS        uint32 // Last lap time in milliseconds
	Sector1TimeMSPart      uint16 // Sector 1 time milliseconds part
	Sector1TimeMinutesPart uint8  // Sector 1 whole minute part
	Sector2TimeMSPart      uint16 // Sector 2 time milliseconds part
	Sector2TimeMinutesPart uint8  // Sector 2 whole minute part
}

func GetSector3(params GetSector3Params) (uint8, uint16) {
	sector1Total := uint32(params.Sector1TimeMinutesPart)*60_000 + uint32(params.Sector1TimeMSPart)
	sector2Total := uint32(params.Sector2TimeMinutesPart)*60_000 + uint32(params.Sector2TimeMSPart)

	sector3Total := int64(params.LastLapTimeInMS) - int64(sector1Total) - int64(sector2Total)
	if sector3Total < 0 {
		sector3Total = 0
	}

	minutes := uint8(sector3Total / 60_000)
	ms := uint16(sector3Total % 60_000)

	return minutes, ms
}
