const SUBSCRIBE_TO_DISABLE_INTERACTION_VIEW = [];
function DisableInteractionView() {
    SUBSCRIBE_TO_DISABLE_INTERACTION_VIEW.forEach(function (func) {
        func();
    });
};