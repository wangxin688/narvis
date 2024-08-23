package helpers

import (
	"github.com/dustin/go-humanize"
)

// Converts bytes to human readable format
func HumanReadableSize(bytes uint64) string {
	return humanize.Bytes(bytes)
}
