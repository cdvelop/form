	// activa o desactiva listener según foco
	function focusInputsActivate(inout) {
		// console.log("focus input activo: ", inout);
		for (let i = 0; i < fieldset.length; i++) {
			if (!inout) {
				fieldset[i].removeEventListener('focusin', focusInputInsideFieldset, true);
			} else {
				fieldset[i].addEventListener('focusin', focusInputInsideFieldset, true);
			};
		};
	};


