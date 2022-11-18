// +build js, wasm

package dom

import "syscall/js"

type Doc struct {
	js.Value
}

func (d *Doc) QuerySelector(selectors string) *Element {
	return &Element{d.Call("querySelector", selectors)}
}

func (d *Doc) CreateElement(args ...interface{}) *Element {
	return &Element{d.Call("createElement", args)}
}

func (d *Doc) GetElementById(id string) *Element {
	return &Element{d.Call("getElementById", id)}
}

func (d *Doc) GetElementsByName(name string) []*Element {
	array := []*Element{}
	elements := d.Call("getElementsByName", name)
	for i := 0; i < elements.Length(); i++ {
		array = append(array, &Element{elements.Index(i)})
	}
	return array
}
