// +build js, wasm

package dom

import "syscall/js"

type ClassList struct {
	js.Value
}

func (c *ClassList) Add(args ...interface{}) {
	c.Call("add", args)
}
