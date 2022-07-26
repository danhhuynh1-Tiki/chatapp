// import logo from './logo.svg';
import './App.css';
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import Login from './pages/Login';
// import Home from './components/pages/Home';
// import SignUp from '; 
import Chat from './pages/Chat';
function App() {
  return (
    <Router>
      <Routes>
        {/* <Route path="/login" element={<SignUp/>}></Route> */}
        <Route path="/login" element={<Login/>}></Route>
        <Route path="/chat" element={<Chat/>}></Route>
      </Routes>
    </Router>

  )
}

export default App;
