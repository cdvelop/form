package form_handler

import (
	_ "embed"
	"strings"
)

//go:embed jsmodule/listener_add.js
var listener_add string

func (f form) AttachADDListener() string {
	sb := strings.Builder{}
	// 1- cargar add listeners básicos
	sb.WriteString(listener_add)

	// 2- cargar listener add input registrados
	for _, field := range f.RenderFields() {
		if js, ok := form_store.js_listeners[field.Input.Name()]; ok {
			sb.WriteString(js.FieldAddEventListener(field.Name))
		}
	}

	return sb.String()
}

//go:embed jsmodule/listener_rem.js
var listener_rem string

func (f *form) AttachREMOVEListener() string {
	sb := strings.Builder{}
	// 1- cargar  remove listeners basicos
	sb.WriteString(listener_rem)

	// 2- cargar listener remove  input registrados
	for _, field := range f.RenderFields() {
		if js, ok := form_store.js_listeners[field.Input.Name()]; ok {
			sb.WriteString(js.FieldRemoveEventListener(field.Name))
		}
	}

	return sb.String()
}
