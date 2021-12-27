export function createLocalstorageItem(txtJwtTokenValue) {
    sessionStorage.setItem("jwtToken", txtJwtTokenValue);
}

export function readValue() {
    var jwtToken = sessionStorage.getItem("jwtToken");
    return jwtToken;
}

export function hasValue() {
    let token = sessionStorage.getItem("jwtToken")
    if (token == null || token == "") {
        return false
    }
    return true;
}

export function setNone() {
    sessionStorage.setItem("jwtToken", "");
}