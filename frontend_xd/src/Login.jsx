import React, { useState } from "react";
import axios from "axios";
import "./Login.css";

const Login = () => {
  const [user, setUser] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState(null);

  const handleLogin = async (e) => {
    e.preventDefault();

    try {
      const response = await axios.post("http://localhost:8000/login", {
        email: user.trim(),
        password: password.trim(),
      });

      console.log("Login exitoso:", response.data);
      setError(null);

      // Guardar token en localStorage
      localStorage.setItem("token", response.data.token);

      // (Opcional) Redirigir a otra página después del login
      // window.location.href = "/home"; // o usá navigate si tenés react-router
    } catch (error) {
      console.error("Error al hacer login:", error);
      if (error.response && error.response.status === 401) {
        setError("Email o contraseña incorrectos");
      } else {
        setError("Error de conexión con el servidor");
      }
    }
  };

  return (
    <div className="login-container">
      <form className="login-form" onSubmit={handleLogin}>
        <h2>Iniciar Sesión</h2>

        <input
          type="text"
          placeholder="Email"
          value={user}
          onChange={(e) => setUser(e.target.value)}
          required
        />

        <input
          type="password"
          placeholder="Contraseña"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />

        <button type="submit">Iniciar Sesión</button>

        {error && <p style={{ color: "red" }}>{error}</p>}
      </form>
    </div>
  );
};

export default Login;
