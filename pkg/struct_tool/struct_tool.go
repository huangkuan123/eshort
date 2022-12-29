package struct_tool

import "reflect"

func Copy(dst any, src any) {
	voValue := reflect.ValueOf(dst).Elem()
	voType := reflect.TypeOf(dst).Elem()
	srcType := reflect.TypeOf(src).Elem()
	srcValue := reflect.ValueOf(src).Elem()
	voFiledCount := voValue.NumField()
	for i := 0; i < voFiledCount; i++ {
		field := voType.Field(i)
		srcField, has := srcType.FieldByName(field.Name)
		if has && field.Type == srcField.Type {
			tSrcValue := srcValue.FieldByName(field.Name)
			voValue.Field(i).Set(tSrcValue)
		}
	}
}
