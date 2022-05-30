export enum UserRole {
    Programmer = "Programmer",
    Tester = "Tester",
}

export interface User {
    id: number,
    username: string,
    role: UserRole,
}
