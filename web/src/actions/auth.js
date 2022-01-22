import { createLocalstorageItem, setNone } from './jwt';
import {auth} from '../client/client';
import router from '../router.js'

export function signInAction(username, password, errField) {
    auth
        .post('/user/entry', {
            'username': username,
            'password': password,
        })
        .then((response) => {
            let token = response.data
            createLocalstorageItem(token)
            router.push('/profile')
        })
        .catch((error) => {
            if (error.response.status === 400) {
                errField.innerHTML = "Invalid username or password"
            } else if (error.response.status === 401) {
                errField.innerHTML = "Invalid username or password"
            } else if (error.response.status === 500) {
                errField.innerHTML = "Server error"
            }
        });
}

export function signUpAction(username, password, email, errField) {
    auth
        .post('/user/new', {
            'username': username,
            'password': password,
            'email': email
        })
        .then((response) => {
            let token = response.data["access_token"]
            createLocalstorageItem(token)
            router.push('/profile')
        })
        .catch((error) => {
            if (error.response.status === 400){
                errField.innerHTML = "Invalid credentials"
            }
            else if (error.response.status === 500) {
                errField.innerHTML = "Server error"
            }
        });
}

export function signOutAction() {
    setNone()
    auth
        .get('/user/exit')
        .then(() => {
            router.push("/")
        })
}

export async function refreshTokens() {
    const response = await auth
        .put('/user/refresh')
        .catch(() => {
            router.push('auth')
        })
    if (response){
        let token = response.data
        createLocalstorageItem(token)
    }
}