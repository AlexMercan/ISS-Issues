import { IssueTag } from "../model/Issue";

export const ISSUES_API_URL = "http://localhost"

export const Unique = (arr: IssueTag[]) => {
    let uniques = [];

    let itemsFound: Record<string, boolean> = {};
    for (let val of arr) {
        if (itemsFound[val.name]) {
            continue;
        }

        uniques.push(val);
        itemsFound[val.name] = true;
    }

    return uniques;
}
