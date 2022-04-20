
const API_URL = "http://localhost/auth/"

export const Register = async (username: string, password: string) => {
    let user = {
        "username": username,
        "password": password
    }
    const resp = await fetch(API_URL + 'register',
        {
            method: 'POST',
            credentials: 'include',
            mode: "cors",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(user)
        })
    if (resp.status === 200) {
        return true
    }
    return false
}

export const Login = async (username: string, password: string) => {
    let user = {
        "username": username,
        "password": password
    }
    const resp = await fetch(API_URL + 'login',
        {
            method: 'POST',
            credentials: 'include',
            mode: "cors",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(user)
        })
    if (resp.status === 200) {
        return true
    }
    return false
}

