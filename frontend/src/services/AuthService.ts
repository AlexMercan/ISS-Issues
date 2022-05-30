import { User } from "../model/User"
import { ISSUES_API_URL } from "./Service"

const API_URL = ISSUES_API_URL + "/auth"

export const Register = async (username: string, password: string, role: string) => {
    const user = {
        "username": username,
        "password": password,
        "role": role
    }
    const resp = await fetch(API_URL + '/register',
        {
            method: 'POST',
            credentials: 'include',
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

export const Login = async (username: string, password: string): Promise<[number, User | null]> => {
    let user = {
        "username": username,
        "password": password
    }
    const resp = await fetch(API_URL + '/login',
        {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(user)
        })

    return [resp.status, await resp.json()]
}
