import React from 'react';
import Nav from './components/Nav';
import Home from './pages/Home'
import Login from './pages/Login'
import Register from './pages/Register'
import Chat from './pages/Chat';
import {BrowserRouter, Route, Routes} from 'react-router-dom';
import Logout from './pages/Logout';

class App extends React.Component {

    render() {
        return (<div>
            <BrowserRouter>
            <div>
              <Nav className='NavBar' />
              <div className="form-signin">
                <Routes>
                  <Route path='/' element={<Login />}/>
                  <Route path='/login' element={<Login />}/>
                  <Route path='/register' element={<Register />}/>
                  <Route path='/logout' element={<Logout />}/>
                </Routes>
                </div>
                <div className='PostPage'>
                <Routes><Route path='/posts' element={<Home />} /></Routes>
                </div>
                <div className='ChatPage'>
                <Routes><Route path='/chat' element={<Chat />} /></Routes>
                </div>
            </div>
            </BrowserRouter>
        </div>)
    }
}

export default App