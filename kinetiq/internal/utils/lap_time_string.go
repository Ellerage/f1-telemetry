package utils

import "fmt"

func GetLapTimeString(timeMs uint32) string {

	minutes := timeMs / 60_000
	seconds := (timeMs % 60_000) / 1000
	ms := timeMs % 1000

	return fmt.Sprintf("%d:%d.%d", minutes, seconds, ms)
}
