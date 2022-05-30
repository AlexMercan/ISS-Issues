import React, { useEffect, useState } from "react"
import { GetUser } from "../services/UserService"
import { UserContext } from "../UserContext"
import { User } from "../model/User"

type Props = {
    children: React.ReactNode
}

export const UserProvider: React.FC<Props> = (props: Props) => {
    const [user, setUser] = useState<User>({} as User)

    useEffect(() => {
        (async () => {
            const [status, user] = await GetUser()
            if (status == 200)
                setUser(user)
        })()
    }, [])

    return (
        <UserContext.Provider value={{ currentUser: user, setUser: setUser }}>
            {props.children}
        </UserContext.Provider >
    )
}
