// import logo from './logo.svg';
import './App.css';
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
// import Home from './components/pages/Home';
import SignUp from './components/pages/Login'; 

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<SignUp/>}></Route>
      </Routes>
    </Router>

  )
}

export default App;
