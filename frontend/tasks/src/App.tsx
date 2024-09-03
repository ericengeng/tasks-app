import React from 'react';
import logo from './logo.svg';
import './App.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import MyTasks from './MyTasks';
import Login from './Login';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/mytasks" element={<MyTasks />} />
      </Routes>
    </Router>
  );
}

export default App;
