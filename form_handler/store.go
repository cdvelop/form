package form_handler

type jsInputFunctions interface {
	AttachJsFunctions() string
}

type jsInputListeners interface {
	FieldAddEventListener(field_name string) string
	FieldRemoveEventListener(field_name string) string
}

// que hace el imput cuando se selecciona con nueva informacion
type selectedTargetChanges interface {
	SelectedTargetChanges() string
}

type inputValueChanges interface {
	InputValueChanges() string
}

type store struct {
	//html form nombre tabla y el formulario compilado
	built_html_form map[string]string

	//js nombre input (Name) y js attach
	js_functions map[string]jsInputFunctions

	js_listeners map[string]jsInputListeners

	js_target_changes map[string]selectedTargetChanges

	js_value_changes map[string]inputValueChanges
}

var form_store = store{
	built_html_form:   map[string]string{},
	js_functions:      map[string]jsInputFunctions{},
	js_listeners:      map[string]jsInputListeners{},
	js_target_changes: map[string]selectedTargetChanges{},
	js_value_changes:  map[string]inputValueChanges{},
}
