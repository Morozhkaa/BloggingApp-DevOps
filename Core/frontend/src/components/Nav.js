import React from 'react';
import {NavLink} from "react-router-dom"

const Nav = () => {
    return (
        <nav className="NavBar">
            <div>
                <div className='navbar-smile'></div>
                <NavLink to="/logout" className="navbar-logout">Logout</NavLink>
                <NavLink to="/register" className="navbar-auth">Register</NavLink>
                <NavLink to="/login" className="navbar-auth">Login</NavLink>

                <NavLink to="/chat" className="navbar-brand">Chat</NavLink>
                <NavLink to="/posts" className="navbar-brand">Posts</NavLink>
            </div>
        </nav>
    );
};

export default Nav;