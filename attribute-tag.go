package form

import (
	"fmt"
	"log"
	"strings"
)

type attribute struct {
	InputId string //id único input

	TagType               string //radio,check,text,number
	CssClass              string // clase css ej: "age"
	PropertieName         string //name="nombre"
	DataSetFieldName      string //"age-date","run","ip"
	ValueIn               string //valor por defecto
	TitleInfo             string //info
	Pattern               string //caracteres para validación
	SkipCompletionAllowed bool   //campo requerido validado
	Minimum               string //valor mínimo
	Maximum               string //valor máximo
	Autocomplete          string //on por defecto
	//textarea
	Rows        string //filas ej 4,5,6
	Cols        string //columnas ej 50,80
	Other       string //otro atributo
	PlaceHolder string
	Step        string
}

// TagType:radio,text,number etc TagInput <input type="input"
func (a attribute) Build() string {

	if a.InputId == "" { //si el id no esta vació ej:(idusuario)
		log.Fatalf("Error ID no declarado en input %v", a)
	}

	internalID := ` id="` + strings.ToLower(a.InputId) + `"`

	if a.PlaceHolder != "" {
		a.PlaceHolder = ` placeholder="` + a.PlaceHolder + `"`
	}

	var class string
	if a.CssClass != "" {
		class = ` class="` + a.CssClass + `"`
	}

	var name string
	if a.PropertieName != "" {
		name = ` name="` + a.PropertieName + `"`
	}

	var dataSetFieldName string
	if a.DataSetFieldName != "" {
		dataSetFieldName = ` data-name="` + a.DataSetFieldName + `"`
	}

	var value string
	if a.ValueIn != "" {
		value = ` value="` + a.ValueIn + `"`
	}

	var info string
	if a.TitleInfo != "" {
		info = ` title="` + a.TitleInfo + `"`
	}

	var valide string
	if a.Pattern != "" { //si cadena de validación no esta vacía
		valide = ` pattern="` + a.Pattern + `"`
	}

	var minmax string
	if a.Maximum != "" || a.Minimum != "" {
		minmax = fmt.Sprintf(` min="%v" max="%v"`, a.Minimum, a.Maximum)
	}

	var autocomplete string
	if a.Autocomplete != "" {
		autocomplete = fmt.Sprintf(` autocomplete="%v"`, a.Autocomplete)
	}

	var required string
	if !a.SkipCompletionAllowed { //si cadena de validación no esta vacía
		required = ` required`
	}

	return fmt.Sprintf(`<input type="%v"%v%v%v%v%v%v%v%v%v%v%v%v%v>`, a.TagType, internalID, class, name, dataSetFieldName, value, info, valide, minmax, autocomplete, required, a.PlaceHolder, a.Other, a.Step)
}
