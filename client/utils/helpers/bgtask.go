package helpers

import (
	"fmt"
)

// startBackgroundTask 是启动后台任务的函数
// 它接受一个函数作为参数，这个函数将在新的goroutine中执行
func BackgroundTask(task func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered in startBackgroundTask: %v\n", r)
			}
		}()
		task()
	}()
}
