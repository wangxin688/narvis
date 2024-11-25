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

package network

import "github.com/samber/lo"

type RadioType string

const (
	Radio24GHz RadioType = "2.4GHz"
	Radio5GHz  RadioType = "5GHz"
	Radio6GHz  RadioType = "6GHz"
)

// ChannelToRadioType returns the radio type of a wireless channel number
// The wireless channel numbers are according to 802.11 standard
// The function returns Radio24GHz for channel numbers from 1 to 13
// The function returns Radio5GHz for channel numbers from 36 to 165
// The function returns Radio6GHz for other channel numbers
func ChannelToRadioType(channel uint64) RadioType {
	channel5G := Get5GChannel()
	if lo.Contains(channel5G, channel) {
		return Radio5GHz
	}
	channel24G := Get24gChannel()
	if lo.Contains(channel24G, channel) {
		return Radio24GHz
	}
	// other channel numbers are considered as 6GHz
	return Radio6GHz
}

func Get24gChannel() []uint64 {
	return []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
}

func Get5GChannel() []uint64 {
	// TODO: this list is not complete, add more channels
	return []uint64{36, 40, 44, 48, 52, 56, 60, 64, 100, 104,
		108, 112, 116, 132, 136, 140, 149, 153, 157, 161, 165}
}
