package model

type PackageIds uint8

const (
	Motion              PackageIds = iota // 0
	Session                               // 1
	Lap                                   // 2
	Event                                 // 3
	Participants                          // 4 // List of participants in the session, mostly relevant for multiplayer
	CarSetups                             // 5 // Packet detailing car setups for cars in the race
	CarTelemetry                          // 6 // Telemetry data for all cars
	CarStatus                             // 7 // Status data for all cars
	FinalClassification                   // 8 // Final classification confirmation at the end of a race
	LobbyInfo                             // 9 // Information about players in a multiplayer lobby
	CarDamage                             // 10 // Damage status for all cars
	SessionHistory                        // 11 // Lap and tyre data for session
	TyreSets                              // 12 // Extended tyre set data
	MotionEx                              // 13 // Extended motion data for player car
	TimeTrial                             // 14 // Time Trial specific data
	LapPositions                          // 15 // Lap positions data
)
