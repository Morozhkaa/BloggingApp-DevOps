import React, {useState} from 'react';
import { Navigate } from 'react-router-dom';
import { ToastContainer, toast } from 'react-toastify';

//const baseAuthUrl = "https://myapp.com/api/auth-service/v1/"
//const baseAuthUrl = "http://localhost:3030/api/auth-service/v1/"
const baseAuthUrl = process.env.REACT_APP_authURL

const Register = () => {
    const [login, setLogin] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e) => {
        e.preventDefault();

        let response = await fetch(baseAuthUrl + 'signup', {
            method: 'POST',
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify({
                login,
                email,
                password
            })
        });
        if (response.ok) {
            setRedirect(true);
        } else if (response.status === 400) {
            toast('🦄 Пользователь с данным email уже существует!', {
                position: "top-right",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                });
        }
    }
    if (redirect) {
        return <Navigate to="/login" />
    }

    return (
        <form onSubmit={submit}>
            <h1>Please register</h1>
            <input type='auth' placeholder="Login" className='AuthInput' required 
                    onChange={e => setLogin(e.target.value)}
            />
            <input type='email' placeholder="Email address" className='AuthInput' required
                    onChange={e => setEmail(e.target.value)}
            />
            <input type="password" placeholder="Password" className='AuthInput' required
                    onChange={e => setPassword(e.target.value)}
            />
            <button className='AuthSubmit' type="submit">Submit</button>
            <ToastContainer />
        </form> 
    );
};

export default Register;