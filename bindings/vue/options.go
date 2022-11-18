// +build js, wasm

package vue

import (
	"syscall/js"
)

func NewOptions(el string, data interface{}) js.Value {
	this := js.Global().Get("Object").New()
	this.Set("el", el)
	this.Set("data", data)
	return this
}
