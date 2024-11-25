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

package processor

import "reflect"

func HasStructField(obj interface{}, fieldName string) bool {
	if obj == nil {
		return false
	}

	value := reflect.ValueOf(obj)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return false
	}

	return value.FieldByName(fieldName).IsValid()
}

func StructGetFieldValue(obj interface{}, fieldName string) (interface{}, bool) {
	value := reflect.ValueOf(obj)

	if value.Kind() != reflect.Struct {
		return nil, false
	}

	fieldValue := value.FieldByName(fieldName)

	if !fieldValue.IsValid() || !fieldValue.CanInterface() || fieldValue.IsNil() {
		return nil, false
	}

	return fieldValue.Interface(), true
}

func HasStructTypeField(t reflect.Type, fieldName string) bool {
	_, found := t.FieldByName(fieldName)
	return found
}
