package model

type CarTelemetryData struct {
	MSpeed                   uint16     // Speed of car in kilometres per hour
	MThrottle                float32    // Amount of throttle applied (0.0 to 1.0)
	MSteer                   float32    // Steering (-1.0 (full lock left) to 1.0 (full lock right))
	MBrake                   float32    // Amount of brake applied (0.0 to 1.0)
	MClutch                  uint8      // Amount of clutch applied (0 to 100)
	MGear                    int8       // Gear selected (1-8, N=0, R=-1)
	MEngineRPM               uint16     // Engine RPM
	MDrs                     uint8      // 0 = off, 1 = on
	MRevLightsPercent        uint8      // Rev lights indicator (percentage)
	MRevLightsBitValue       uint16     // Rev lights (bit 0 = leftmost LED, bit 14 = rightmost LED)
	MBrakesTemperature       [4]uint16  // Brakes temperature (celsius)
	MTyresSurfaceTemperature [4]uint8   // Tyres surface temperature (celsius)
	MTyresInnerTemperature   [4]uint8   // Tyres inner temperature (celsius)
	MEngineTemperature       uint16     // Engine temperature (celsius)
	MTyresPressure           [4]float32 // Tyres pressure (PSI)
	MSurfaceType             [4]uint8   // Driving surface, see appendices
}

type CarTelemetryPacket struct {
	Header           PacketHeader // Header
	CarTelemetryData [22]CarTelemetryData
	MfdPanelIndex    uint8 // Index of MFD panel open - 255 = MFD closed
	// Single player, race – 0 = Car setup, 1 = Pits
	// 2 = Damage, 3 = Engine, 4 = Temperatures
	// May vary depending on game mode
	MfdPanelIndexSecondaryPlayer uint8 // See above
	SuggestedGear                int8  // Suggested gear for the player (1-8) // 0 if no gear suggested
}
