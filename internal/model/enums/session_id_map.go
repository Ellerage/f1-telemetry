package model

var SessionIdMap = map[uint8]string{
	0:  "Unknown",
	1:  "Practice 1",
	2:  "Practice 2",
	3:  "Practice 3",
	4:  "Short Practice",
	5:  "Qualifying 1",
	6:  "Qualifying 2",
	7:  "Qualifying 3",
	8:  "Short Qualifying",
	9:  "One-Shot Qualifying",
	10: "Sprint Shootout 1",
	11: "Sprint Shootout 2",
	12: "Sprint Shootout 3",
	13: "Short Sprint Shootout",
	14: "One-Shot Sprint Shootout",
	15: "Race",
	16: "Race 2",
	17: "Race 3",
	18: "Time Trial",
}

const (
	SessionTypeUnknown uint8 = iota
	SessionTypePractice1
	SessionTypePractice2
	SessionTypePractice3
	SessionTypeShortPractice
	SessionTypeQualifying1
	SessionTypeQualifying2
	SessionTypeQualifying3
	SessionTypeShortQualifying
	SessionTypeOneShotQualifying
	SessionTypeSprintShootout1
	SessionTypeSprintShootout2
	SessionTypeSprintShootout3
	SessionTypeShortSprintShootout
	SessionTypeOneShotSprintShootout
	SessionTypeRace
	SessionTypeRace2
	SessionTypeRace3
	SessionTypeTimeTrial
)
