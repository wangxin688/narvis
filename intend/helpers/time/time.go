// Copyright 2024 wangxin.jeffry@gmail.com

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// package nettyx_time defines standardized ways for time processing.

package nettyx_time

import (
	"errors"
	"fmt"
	"math"
	"time"
)

var errNaNOrInf = errors.New("value is Null or Infinity")

func Float2Time(v float64) (*time.Time, error) {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return nil, errNaNOrInf
	}
	timestamp := v * 1e9
	if timestamp > math.MaxInt64 || timestamp < math.MinInt64 {
		return nil, fmt.Errorf("%v cannot be represented as a nanoseconds timestamp since it overflows int64", v)
	}
	sec := int64(v)
	nanoSec := int64(v - timestamp)
	result := time.Unix(sec, nanoSec)
	return &result, nil
}

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
		return fmt.Sprintf("%dd %dh", days, hours)
	} else if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	} else if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	}
	return fmt.Sprintf("%ds", seconds)
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
