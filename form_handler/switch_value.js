function switchValueChange(e) {
    console.log("form switchValueChange: ", e);
   
    let input = e.target;

    console.log("INPUT CHANGE by DATA SET NAME: [", input.dataset.name, "] INPUT: ", input);
    switch (input.dataset.name) {
        // {{.SwitchOptions}}

        default:
            InputValidationWithValidity(input);
            break;
    }
};