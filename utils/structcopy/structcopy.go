package structcopy

import (
	"fmt"
	"reflect"
)

// StructCopy is a function that copies the fields of a struct to another struct.
// using reflection and based on json tag.
func StructCopy(src interface{}, dst interface{}) error {
	sourceValue := reflect.ValueOf(src)
	destValue := reflect.ValueOf(dst)

	if sourceValue.Kind() != reflect.Struct || destValue.Kind() != reflect.Struct {
		return fmt.Errorf("both source and destination must be structs")
	}

	for i := 0; i < sourceValue.NumField(); i++ {
		sourceField := sourceValue.Type().Field(i)
		sourceTag := sourceField.Tag.Get("json")
		destField, ok := destValue.Type().FieldByName(sourceField.Name)
		if !ok {
			continue
		}
		destTag := destField.Tag.Get("json")
		if sourceTag != "" && sourceTag == destTag {
			destValue.FieldByName(sourceField.Name).Set(sourceValue.Field(i))
		}
	}

	return nil
}
