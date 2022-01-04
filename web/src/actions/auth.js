import { createLocalstorageItem, setNone } from './jwt';
import client from '../client/client';

export function signInAction(username, password, errField) {
    client
        .post('/user/entry', {
            'username': username,
            'password': password,
        }, { withCredentials: true })
        .then((response) => {
            let token = response.data["access_token"]
            createLocalstorageItem(token)
            window.location.href = "/profile"
        })
        .catch((error) => {
            if (error.response.status === 401) {
                errField.innerHTML = "Username or password is not valid"
            }
        });
}

export function signUpAction(username, password, email, errField) {
    client
        .post('/user/new', {
            'username': username,
            'password': password,
            'email': email
        }, { withCredentials: true })
        .then((response) => {
            let token = response.data["access_token"]
            createLocalstorageItem(token)
            window.location.href = "/profile"
        })
        .catch((error) => {
            if (error.response.status === 403) {
                errField.innerHTML = error.response.data
            }
        });
}

export function signOutAction() {
    setNone()
    client
        .get('/user/exit', { withCredentials: true })
        .then(() => {
            window.location.href = "/"
        })
}

export function refreshTokens() {
    client
        .put('/user/refresh', {}, {
            withCredentials: true,
        })
        .then((response) => {
            let token = response.data["access_token"]
            createLocalstorageItem(token)
            document.location.reload();
        })
        .catch(() => {
            window.location.href = "/auth"
        });


}