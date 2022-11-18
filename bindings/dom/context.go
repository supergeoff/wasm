// +build js, wasm

package dom

import "syscall/js"

type Context struct {
	js.Value
}
