import React, {useState} from 'react';
import { Navigate } from 'react-router-dom';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { Cookies } from 'react-cookie';

//const baseAuthUrl = "https://myapp.com/api/auth-service/v1/"
//const baseAuthUrl = "http://localhost:3030/api/auth-service/v1/"
const baseAuthUrl = process.env.REACT_APP_authURL

const Login = () => {

    console.log("Auth url:")
    console.log(process.env.REACT_APP_authURL)

    const [login, setLogin] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);

    const submit = async (e) => {
        e.preventDefault();
        let response = await fetch(baseAuthUrl + 'login', {
            method: 'POST',
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify({
                login,
                password
            })
        });
        console.log(response)
        if (response.ok) {
            let body = await response.json();
            console.log(body)
            const cookies = new Cookies();
            cookies.set('access_token', body.access_token, { sameSite: 'lax', secure: false, httpOnly: false})
            cookies.set('refresh_token', body.refresh_token, { sameSite: 'lax', secure: false, httpOnly: false});
            console.log(cookies.get('access_token'))
            console.log(cookies.get('refresh_token'))
            setRedirect(true);
        } else if (response.status === 404) {
            toast('🦄 Вы еще не зарегистрированы!', {
                position: "top-right",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                });
        } else if (response.status === 403) {
            toast('🦄 Неправильный пароль', {
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
        return <Navigate to="/posts" />
    }

    return (
        <form onSubmit={submit}>
            <h1>Please sign in</h1>
            <input type='auth' placeholder="Login" className='AuthInput' required 
                    onChange={e => setLogin(e.target.value)}
            /> 
            <input type="password" placeholder="Password" className='AuthInput' required
                    onChange={e => setPassword(e.target.value)}
            />
            <button className='AuthSubmit' type="submit">Sign in</button>
            <ToastContainer />
        </form>
    );
};

export default Login;