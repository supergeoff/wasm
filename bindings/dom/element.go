// +build js, wasm

package dom

import "syscall/js"

type Element struct {
	js.Value
}

func (e *Element) ClassList() *ClassList {
	return &ClassList{e.Get("classList")}
}

func (e *Element) AppendChild(elem *Element) {
	e.Call("appendChild", elem)
}

func (e *Element) GetContext(args ...interface{}) *Context {
	return &Context{e.Call("getContext", args)}
}

func (e *Element) AddEventListener(event string, callback func(js.Value, []js.Value) interface{}) {
	jsCallback := js.FuncOf(callback)
	e.Call("addEventListener", event, jsCallback)
}

func (e *Element) SetInnerHTML(text string) {
	e.Set("innerHTML", text)
}
