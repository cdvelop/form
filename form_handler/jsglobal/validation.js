function InputValidationWithValidity(e) {
    let input = e.target;
    if (input.validity.valid && input.value != "") {
        changeFieldsetColor(input, true);
    } else {
        changeFieldsetColor(input, false);
    }
};

function InputValidationWithPattern(e) {
    let input = e.target;
    let pattern = e.target.pattern;
    
    // console.log("VALOR:", input.value," PATTERN: ", pattern)
    if (pattern.test(input.value)) {
        changeFieldsetColor(input, true);
    } else {
        changeFieldsetColor(input, false);
    }
};

function changeFieldsetColor(input, option) {
    // console.log("INPUT A CAMBIAR DE COLOR: ", input)
    if (option) {
        InputRightInForm(input);
    } else if (input.hasAttribute('required')) {
        InputWrong(input);
    } else {
        InputWrong(input);
    }
};