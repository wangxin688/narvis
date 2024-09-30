package helpers

import (
	"fmt"
	"runtime/debug"

	"github.com/wangxin688/narvis/client/utils/logger"
)

// startBackgroundTask 是启动后台任务的函数
// 它接受一个函数作为参数，这个函数将在新的goroutine中执行
func BackgroundTask(taskFunc func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Logger.Error(fmt.Sprintf("Recovered in startBackgroundTask: %v\n", r))
				logger.Logger.Error(string(debug.Stack()))
			}
		}()
		taskFunc()
	}()
}
