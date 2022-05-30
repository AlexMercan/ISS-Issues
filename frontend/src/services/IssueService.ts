import { Issue } from "../model/Issue";
import { ISSUES_API_URL } from "./Service";

const API_URL = ISSUES_API_URL + '/api/issues'

export const UpdateIssue = async (issue: Issue) => {
    let response = await fetch(API_URL + "/" + issue.id, {
        method: "POST",
        headers: {
            "content-type": "application/json"
        },
        body: JSON.stringify(issue)
    })
    return response.status
}

export const SaveIssue = async (issue: Issue) => {
    let response = await fetch(API_URL, {
        method: "POST",
        headers: {
            "content-type": "application/json"
        },
        body: JSON.stringify(issue)
    })
    return response.status
}
