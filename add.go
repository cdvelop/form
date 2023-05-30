package form

import "github.com/cdvelop/model"

var Form = model.Component{
	Name:        "form",
	CssGlobal:   nil,
	CssPrivate:  nil,
	JsGlobal:    form{},
	JsPrivate:   nil,
	JsListeners: nil,
}

type form struct{}
