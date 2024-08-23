package audit

import (
	"encoding/json"
	"fmt"
	"reflect"

	"gorm.io/datatypes"
)

func getItemType(v reflect.Type) reflect.Type {
	switch v.Kind() {
	case reflect.Slice, reflect.Array, reflect.Pointer:
		return getItemType(v.Elem())
	default:
		return v
	}
}

func getPkKeyValue(v reflect.Value) string {
	for v.Kind() == reflect.Pointer {
		v = reflect.Indirect(v)
	}
	var primaryKeyValues string
	switch v.Kind() {
	case reflect.Struct:
		primaryKeyValues = v.FieldByName("ID").String()
	case reflect.Map:
		primaryKeyValues = v.MapIndex(reflect.ValueOf("id")).String()
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
