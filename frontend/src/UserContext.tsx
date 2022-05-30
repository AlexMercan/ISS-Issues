import { createContext, useContext } from "react";
import { User } from "./model/User";

interface IUserContext {
    currentUser?: User;
    setUser: (newUser: User) => void;
}

export const defaultUserContext: IUserContext = {
    setUser: ({ }) => console.warn("Set user not implemented on user context")
}

export const UserContext: React.Context<IUserContext> = createContext(defaultUserContext)
export const useUser = () => useContext(UserContext)
