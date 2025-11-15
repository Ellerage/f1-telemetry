package model

// 1 Hz update
type MarshalZoneData struct {
	ZoneStart float32 // Fraction (0..1) of way through the lap the marshal zone starts
	ZoneFlag  int8    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow
}

type WeatherForecastData struct {
	SessionType            uint8 // 0 = unknown, see appendix
	TimeOffset             uint8 // Time in minutes the forecast is for
	Weather                uint8 // Weather - 0 = clear, 1 = light cloud, 2 = overcast, 3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature       int8  // Track temp. in degrees Celsius
	TrackTemperatureChange int8  // Track temp. change – 0 = up, 1 = down, 2 = no change
	AirTemperature         int8  // Air temp. in degrees celsius
	AirTemperatureChange   int8  // Air temp. change – 0 = up, 1 = down, 2 = no change
	RainPercentage         uint8 // Percentage chance of rain (0-100)
}

type SessionPacket struct {
	Header                          PacketHeader            // Packet header
	Weather                         uint8                   // Weather
	TrackTemperature                int8                    // Track temperature
	AirTemperature                  int8                    // Air temperature
	TotalLaps                       uint8                   // Total laps
	TrackLength                     uint16                  // Track length
	SessionType                     uint8                   // Session type
	TrackId                         int8                    // Track ID
	Formula                         uint8                   // Formula
	SessionTimeLeft                 uint16                  // Session time left
	SessionDuration                 uint16                  // Session duration
	PitSpeedLimit                   uint8                   // Pit speed limit
	GamePaused                      uint8                   // Game paused
	IsSpectating                    uint8                   // Is spectating
	SpectatorCarIndex               uint8                   // Spectator car index
	SliProNativeSupport             uint8                   // SLI Pro native support
	NumMarshalZones                 uint8                   // Number of marshal zones
	MarshalZones                    [21]MarshalZoneData     // Marshal zones
	SafetyCarStatus                 uint8                   // Safety car status
	NetworkGame                     uint8                   // Network game
	NumWeatherForecastSamples       uint8                   // Number of weather forecast samples
	WeatherForecastSamples          [64]WeatherForecastData // Weather forecast samples
	ForecastAccuracy                uint8                   // Forecast accuracy
	AIDifficulty                    uint8                   // AI difficulty
	SeasonLinkIdentifier            uint32                  // Season link identifier
	WeekendLinkIdentifier           uint32                  // Weekend link identifier
	SessionLinkIdentifier           uint32                  // Session link identifier
	PitStopWindowIdealLap           uint8                   // Pit stop window ideal lap
	PitStopWindowLatestLap          uint8                   // Pit stop window latest lap
	PitStopRejoinPosition           uint8                   // Pit stop rejoin position
	SteeringAssist                  uint8                   // Steering assist
	BrakingAssist                   uint8                   // Braking assist
	GearboxAssist                   uint8                   // Gearbox assist
	PitAssist                       uint8                   // Pit assist
	PitReleaseAssist                uint8                   // Pit release assist
	ERSAssist                       uint8                   // ERS assist
	DRSAssist                       uint8                   // DRS assist
	DynamicRacingLine               uint8                   // Dynamic racing line
	DynamicRacingLineType           uint8                   // Dynamic racing line type
	GameMode                        uint8                   // Game mode
	RuleSet                         uint8                   // Rule set
	TimeOfDay                       uint32                  // Time of day
	SessionLength                   uint8                   // Session length
	SpeedUnitsLeadPlayer            uint8                   // Speed units lead player
	TemperatureUnitsLeadPlayer      uint8                   // Temperature units lead player
	SpeedUnitsSecondaryPlayer       uint8                   // Speed units secondary player
	TemperatureUnitsSecondaryPlayer uint8                   // Temperature units secondary player
	NumSafetyCarPeriods             uint8                   // Number of safety car periods
	NumVirtualSafetyCarPeriods      uint8                   // Number of virtual safety car periods
	NumRedFlagPeriods               uint8                   // Number of red flag periods
	EqualCarPerformance             uint8                   // Equal car performance
	RecoveryMode                    uint8                   // Recovery mode
	FlashbackLimit                  uint8                   // Flashback limit
	SurfaceType                     uint8                   // Surface type
	LowFuelMode                     uint8                   // Low fuel mode
	RaceStarts                      uint8                   // Race starts
	TyreTemperature                 uint8                   // Tyre temperature
	PitLaneTyreSim                  uint8                   // Pit lane tyre sim
	CarDamage                       uint8                   // Car damage
	CarDamageRate                   uint8                   // Car damage rate
	Collisions                      uint8                   // Collisions
	CollisionsOffForFirstLapOnly    uint8                   // Collisions off for first lap only
	MpUnsafePitRelease              uint8                   // MP unsafe pit release
	MpOffForGriefing                uint8                   // MP off for griefing
	CornerCuttingStringency         uint8                   // Corner cutting stringency
	ParcFermeRules                  uint8                   // Parc ferme rules
	PitStopExperience               uint8                   // Pit stop experience
	SafetyCar                       uint8                   // Safety car
	SafetyCarExperience             uint8                   // Safety car experience
	FormationLap                    uint8                   // Formation lap
	FormationLapExperience          uint8                   // Formation lap experience
	RedFlags                        uint8                   // Red flags
	AffectsLicenceLevelSolo         uint8                   // Affects licence level solo
	AffectsLicenceLevelMP           uint8                   // Affects licence level MP
	NumSessionsInWeekend            uint8                   // Number of sessions in weekend
	WeekendStructure                [12]uint8               // Weekend structure
	Sector2LapDistanceStart         float32                 // Sector 2 lap distance start
	Sector3LapDistanceStart         float32                 // Sector 3 lap distance start
}
