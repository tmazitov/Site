export function createLocalstorageItem(token) {
    sessionStorage.setItem("jwtToken", JSON.stringify(token));
}

export function readValue() {
    var jwtToken = JSON.parse(sessionStorage.getItem("jwtToken")).access_token;
    return jwtToken;
}

export function tokenIsValid() {
    let token = JSON.parse(sessionStorage.getItem("jwtToken"))
    if (token.access_token == null || token.access_token == "" || Date.now() >= token.expires_in * 1000 ) {
        return false
    }
    return true;
}

export function setNone() {
    sessionStorage.setItem("jwtToken", "");
}