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

// package file provide standardized ways for file processing

package file

import (
	"errors"
	"os"
)

// PathExists 判断所给文件夹是否存在
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("duplicate file name for the given path: " + path)
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateDir 创建文件夹
func CreateDir(dirName string) error {
	exist, err := PathExists(dirName)
	if err != nil {
		return err
	}
	if !exist {
		// 创建文件夹
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteDir 删除文件夹
func DeleteDir(dirName string) error {
	err := os.RemoveAll(dirName)
	if err != nil {
		return err
	}
	return nil
}

// FileExists 判断所给路径文件否存在
func FileExists(path string) bool {
	fi, err := os.Lstat(path) // os.Stat获取文件信息
	if err != nil {
		return !fi.IsDir()
	}
	return os.IsExist(err)
}

// DeleteFile 删除文件
func DeleteFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	return nil
}

// FileSize 获取文件大小
func FileSize(path string) (int64, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}
