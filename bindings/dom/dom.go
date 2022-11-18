// +build js, wasm

package dom

import "syscall/js"

func Document() *Doc {
	return &Doc{js.Global().Get("document")}
}
