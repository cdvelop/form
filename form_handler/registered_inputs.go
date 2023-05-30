package form_handler

func RegisterInputsJsModule(dataset_name string, F jsInputFunctions, L jsInputListeners, T selectedTargetChanges, V inputValueChanges) {

	if F != nil {
		if _, no_exist := form_store.js_functions[dataset_name]; !no_exist {
			form_store.js_functions[dataset_name] = F

		}
	}

	if L != nil {

		if _, no_exist := form_store.js_listeners[dataset_name]; !no_exist {
			form_store.js_listeners[dataset_name] = L

		}
	}

	if T != nil {

		if _, no_exist := form_store.js_target_changes[dataset_name]; !no_exist {
			form_store.js_target_changes[dataset_name] = T

		}
	}

	if V != nil {

		if _, no_exist := form_store.js_value_changes[dataset_name]; !no_exist {
			form_store.js_value_changes[dataset_name] = V
		}
	}

}
