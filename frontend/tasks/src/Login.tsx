import React from "react";
import { useNavigate } from "react-router-dom";
import './css/Login.css'

const Login: React.FC = () => {
  const navigate = useNavigate();
  const handleLogin = async () => {
    navigate('/mytasks');
  };
  const handleSignUp = async () => {
    navigate('/mytasks');
  };
  return (
    <div className="login-container">

      Enter Username:
      <input type="text"
        placeholder="Username" />

      Enter Password:
      <input type="password"
        placeholder="Password" />


      <button onClick={handleLogin}>Sign In</button>
      <button onClick={handleSignUp}>Sign Up</button>
    </div>
  );
};

export default Login;