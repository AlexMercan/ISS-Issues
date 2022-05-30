export enum IssueStatus {
    Open = "Open",
    Closed = "Closed",
}

export type IssueTag = {
    id: number,
    name: string
}

export type Edge = {
    assignedTags: IssueTag[]
}

export type Issue = {
    id: number,
    name: string,
    description: string,
    status: IssueStatus,
    owner_id: number,
    edges: Edge
}
