export function createLocalstorageItem(txtJwtTokenValue) {
    sessionStorage.setItem("jwtToken", txtJwtTokenValue);
}

export function readValue() {
    var jwtToken = sessionStorage.getItem("jwtToken");
    return jwtToken;
}

export function hasValue() {
    return sessionStorage.getItem("jwtToken") == "";
}

export function setNone() {
    sessionStorage.setItem("jwtToken", "");
}