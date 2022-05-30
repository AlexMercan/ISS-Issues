import React from 'react';
import './App.css';

import { Navigate, RouteObject, useRoutes } from 'react-router'

import Login from './components/Login';
import Register from './components/Register';
import { Issues } from './components/Issues'
import { IssuePage } from "./components/IssuePage";
import { UserProvider } from './components/UserProvider';
import { AddIssue } from './components/AddIssue';

function App() {
    const loginRoute: RouteObject = {
        path: "/login",
        element: <Login />
    }
    const registerRoute: RouteObject = {
        path: "/register",
        element: <Register />
    }
    const mainRoute: RouteObject = {
        path: "/issues",
        element: <Issues />
    }
    const issueRoute: RouteObject = {
        path: "/issues/:issueId",
        element: <IssuePage />
    }
    const addIssueRoute: RouteObject = {
        path: "/issues/add",
        element: <AddIssue />
    }
    const redirectToMain: RouteObject = {
        path: "*",
        element: <Navigate replace to="/issues" />
    }
    let routing = useRoutes([loginRoute, registerRoute, mainRoute, issueRoute, redirectToMain, addIssueRoute])
    return (
        <UserProvider children={routing}></UserProvider>
    );
}

export default App;
