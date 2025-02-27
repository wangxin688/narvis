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

package bgtask

import (
	"fmt"
	"runtime/debug"

	nlog "github.com/wangxin688/narvis/intend/logger"
)

func BackgroundTask(taskFunc func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				nlog.Logger.Error(fmt.Sprintf("Recovered in startBackgroundTask: %v\n", r))
				nlog.Logger.Error(string(debug.Stack()))
			}
		}()
		taskFunc()
	}()
}
