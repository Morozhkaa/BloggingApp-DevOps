import React from 'react';
import { Navigate } from 'react-router-dom';
import 'react-toastify/dist/ReactToastify.css';
import Cookies from 'universal-cookie';

const cookies = new Cookies();

const Logout = () => {
    cookies.remove("access_token")
    cookies.remove("refresh_token")
    cookies.set("role", "GUEST")
    return <Navigate to="/"/>
};

export default Logout;