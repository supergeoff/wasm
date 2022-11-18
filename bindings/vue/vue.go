// +build js, wasm

package vue

import (
	"syscall/js"
)

func NewVue(options js.Value) js.Value {
	this := js.Global().Get("Vue").New(options)
	return this
}
