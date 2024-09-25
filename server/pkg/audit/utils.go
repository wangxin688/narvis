package audit

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/samber/lo"
	"gorm.io/datatypes"
	"gorm.io/gorm/schema"
)

type snapshot struct {
	primaryKeyValues []string
	data             map[string]any
}

func getItemType(v reflect.Type) reflect.Type {
	switch v.Kind() {
	case reflect.Slice, reflect.Array, reflect.Pointer:
		return getItemType(v.Elem())
	default:
		return v
	}
}

func getPkKeyValue(v reflect.Value) []string {
	for v.Kind() == reflect.Pointer {
		v = reflect.Indirect(v)
	}

	primaryKeyValues := make([]string, 0)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			primaryKeyValues = append(primaryKeyValues, getPkKeyValue(v.Index(i))...)
		}
	case reflect.Struct:
		primaryKeyValues = append(primaryKeyValues, v.FieldByName("Id").String())
	case reflect.Map:
		primaryKeyValues = append(primaryKeyValues, v.MapIndex(reflect.ValueOf("Id")).String())
	}
	return primaryKeyValues
}

func diffChange(before, after map[string]any) (diffBefore, diffAfter map[string]any) {
	diffBefore, diffAfter = make(map[string]any), make(map[string]any)
	for key, value := range before {
		if afterValue, ok := after[key]; ok {
			if !reflect.DeepEqual(value, afterValue) {
				diffAfter[key] = afterValue
				diffBefore[key] = value
			}
		} else {
			diffBefore[key] = value
		}
	}

	for key, value := range after {
		if _, ok := before[key]; !ok {
			diffAfter[key] = value
		}
	}
	return diffBefore, diffAfter
}

func getKeyFromMap(key string, m map[string]any) string {

	if v, ok := m[key]; ok {
		return fmt.Sprintf("%v", v)
	}
	return ""
}

func prepareData(data map[string]any) datatypes.JSON {
	dataByte, _ := json.Marshal(&data)
	return dataByte
}

func getSnapshot(data any, fields []*schema.Field) (*snapshot, error) {
	v := reflect.ValueOf(data)
	for v.Kind() == reflect.Pointer {
		v = reflect.Indirect(v)
	}

	if !lo.Contains([]reflect.Kind{reflect.Array, reflect.Slice}, v.Kind()) {
		newV := reflect.New(reflect.SliceOf(v.Type())).Elem()
		newV = reflect.Append(newV, v)
		v = newV
	}
	s := &snapshot{data: make(map[string]any)}
	for i := 0; i < v.Len(); i++ {
		tmp := v.Index(i)
		for tmp.Kind() == reflect.Pointer {
			tmp = reflect.Indirect(tmp)
		}
		if tmp.Kind() != reflect.Struct {
			return nil, fmt.Errorf("unsupported type: %v", tmp.Kind())
		}

		id := tmp.FieldByName("Id").String()
		s.primaryKeyValues = append(s.primaryKeyValues, id)
		tmtData := make(map[string]any)
		for _, field := range fields {
			tag := field.Tag.Get("gorm")
			if strings.Contains(tag, "Ondelete") || strings.Contains(tag, "Onupdate") || strings.Contains(tag, "foreignKey") {
				continue
			}
			tmtData[field.Name] = tmp.FieldByName(field.Name).Interface()
		}
		s.data[id] = tmtData
	}

	return s, nil
}