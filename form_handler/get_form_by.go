package form_handler

func GetHtmlFormBy(table_name string) string {
	html_form, exist := form_store.built_html_form[table_name]
	if !exist {
		return "Error Formulario " + table_name + "  no ha sido registrado"
	}

	return html_form
}
