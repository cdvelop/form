package form_handler

import (
	"embed"

	"github.com/cdvelop/model"
)

type form struct {
	*model.Object
	config
}

type config struct {
	Reset         string
	SwitchOptions string
}

//go:embed jsglobal/*.js
var jsglobal embed.FS

// form form function ej: "resetForm();" si esta vació se ejecutara form.reset();
func New(t *model.Object, reset_form_function string) *form {
	f := form{
		Object: t,
		config: config{
			Reset: reset_form_function,
		},
	}

	f.RegisteredHtmlForm()
	return &f

}
