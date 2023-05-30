
	function focusInputsEnable() {
		focusInputsActivate(true)	
	};
	
	SUBSCRIBE_ACTION_BUTTON_ADD.push(focusInputsEnable);



	
	function FocusInputFormByINDEX(number) {
		// foco en elemento del formulario
		form.querySelectorAll("FIELDSET")[number].focus();
	};

	function FocusOnFirstInput() {
		form.querySelectorAll("FIELDSET")[0].focus();
	};

