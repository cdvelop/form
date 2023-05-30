function focusInputInsideFieldset(f) {
    // console.log("FOCO: ", f.target)
    if (f.srcElement.tagName === 'FIELDSET') {
        let parents = f.srcElement.children;
        if (parents.length > 1) {
            let element = parents[1];
            // console.log("ELEMENT: ", element, "hijos: ", parents);
            switch (element.tagName) {
                case "DIV":
                    let inputs = element.querySelectorAll("input");
                    // console.log(inputs, " tamaño: ", inputs.length)
                    if (inputs != undefined && inputs.length > 0) {
                        inputs[0].focus();
                    }
                    break;

                default:
                    element.focus();
                    break;
            }
        };
    };
};