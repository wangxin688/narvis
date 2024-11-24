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

package nettyx_time_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	nettyx_time "github.com/wangxin688/narvis/intend/helpers/time"
)

func TestHumanizeDuration(t *testing.T) {
	tc := []struct {
		name     string
		input    int64
		expected string
	}{
		// Integers
		{name: "zero", input: 0, expected: "0s"},
		{name: "one second", input: 1, expected: "1s"},
		{name: "one minute", input: 60, expected: "1m 0s"},
		{name: "one hour", input: 3600, expected: "1h 0m"},
		{name: "one day", input: 86400, expected: "1d 0h"},
		{name: "one day and one hour", input: 86400 + 3600, expected: "1d 1h"},
		{name: "negative duration", input: -(86400*2 + 3600*3 + 60*4 + 5), expected: "-5s"},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			result := nettyx_time.HumanReadableDuration(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
