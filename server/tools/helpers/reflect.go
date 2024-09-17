package helpers

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
