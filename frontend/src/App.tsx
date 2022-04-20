import React from 'react';
import logo from './logo.svg';
import './App.css';

import { RouteObject, useRoutes } from 'react-router'

import { Link, Route } from 'react-router-dom';
import SignIn from './components/SignIn';
import SignUp from './components/SignUp';
import { Projects } from './components/Projects'

function App() {
    const loginRoute: RouteObject = {
        path: "/login",
        element: <SignIn />
    }
    const registerRoute: RouteObject = {
        path: "/register",
        element: <SignUp />
    }
    const mainRoute: RouteObject = {
        path: "/",
        element: <Projects />
    }
    const routing = useRoutes([loginRoute, registerRoute, mainRoute])
    return (
        <>
            {routing}
        </>
    );
}
export default App;
