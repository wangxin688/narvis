package helpers

import (
	"fmt"
	"time"
)

// HumanReadableDuration converts seconds to string
func HumanReadableDuration(seconds int64) string {
	const (
		secondsInMinute = 60
		secondsInHour   = 60 * 60
		secondsInDay    = 24 * 60 * 60
	)

	days := seconds / secondsInDay
	seconds %= secondsInDay

	hours := seconds / secondsInHour
	seconds %= secondsInHour

	minutes := seconds / secondsInMinute
	seconds %= secondsInMinute

	if days > 0 {
		return fmt.Sprintf("%d days %d hours", days, hours)
	} else if hours > 0 {
		return fmt.Sprintf("%d hours %d minutes", hours, minutes)
	} else if minutes > 0 {
		return fmt.Sprintf("%d minutes %d seconds", minutes, seconds)
	}
	return fmt.Sprintf("%d seconds", seconds)
}

// ShortDuration converts seconds to string with shorten string: `25h1m1s“
func ShortDuration(seconds int) string {
	seconds_ := int64(seconds)
	duration := time.Duration(seconds_) * time.Second
	return duration.String()
}

func TimeTicksToDuration(ticks uint64) string {
	seconds := int64(ticks) / 100
	return HumanReadableDuration(seconds)
}
