package device360_utils

import "strconv"

func ScoreToHealth(score string) string {
	intScore, _ := strconv.Atoi(score)
	switch {
	case intScore > 7:
		return "good"
	case intScore > 4 && intScore < 8:
		return "needsAttention"
	case intScore < 4 && intScore > 0:
		return "poor"
	case intScore == 0:
		return "unhealthy"
	default:
		return "unknown"
	}
}
