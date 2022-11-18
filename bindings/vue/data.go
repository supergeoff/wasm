// +build js, wasm

package vue

import (
	"reflect"
	"syscall/js"
)

func NewData(data interface{}) js.Value {
	value := reflect.ValueOf(data)
	switch value.Kind() {
	case reflect.Struct:
		obj := js.Global().Get("Object").New()
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			tag := value.Type().Field(i).Tag.Get("js")
			obj.Set(tag, NewData(field.Interface()))
		}
		return obj
	case reflect.Slice:
		arr := js.Global().Get("Array").New()
		for i := 0; i < value.Len(); i++ {
			field := value.Index(i)
			arr.SetIndex(i, NewData(field.Interface()))
		}
		return arr
	}
	return js.ValueOf(data)
}
