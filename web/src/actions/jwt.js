export function createLocalstorageItem(token) {
    sessionStorage.setItem("jwtToken", JSON.stringify(token));
}

export function readValue() {
    var jwtToken = sessionStorage.getItem("jwtToken")
    if (!jwtToken) {
        return null
    }
    return JSON.parse(jwtToken).access_token
}

export function tokenIsValid() {
    let token = sessionStorage.getItem("jwtToken")
    if (token == null || token == "") {
        return false
    }

    token = JSON.parse(token)

    if (token.access_token == null || token.access_token == "" || Date.now() >= token.expires_in * 1000 ) {
        return false
    }
    return true;
}

export function setNone() {
    sessionStorage.setItem("jwtToken", "");
}