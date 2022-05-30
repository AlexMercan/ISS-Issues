import { ISSUES_API_URL } from "./Service"
import { User } from "../model/User"

const API_URL = ISSUES_API_URL + "/api"

export const GetUser = async (): Promise<[number, User]> => {
    const response = await fetch(API_URL + "/user", {
        headers: {
            "accept": "application/json"
        },
    })
    return [response.status, await (response.json())];
}
