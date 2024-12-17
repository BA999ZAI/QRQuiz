import React, { useState } from "react";

const AuthForm = () => {
  const [isLogin, setIsLogin] = useState(true); // Флаг для переключения между формами
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    if (isLogin) {
      console.log("Авторизация:", { email, password });
    } else {
      console.log("Регистрация:", { email, password });
    }
  };

  return (
    <div className="auth-container">
      <div className="auth-box">
        <h2>{isLogin ? "Авторизация" : "Регистрация"}</h2>
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="email">Email</label>
            <input
              type="email"
              id="email"
              placeholder="Введите email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="password">Пароль</label>
            <input
              type="password"
              id="password"
              placeholder="Введите пароль"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
          <button type="submit" className="auth-button">
            {isLogin ? "Войти" : "Зарегистрироваться"}
          </button>
        </form>
        <div className="toggle-form">
          <p>
            {isLogin ? "Нет аккаунта?" : "Уже есть аккаунт?"}
            <span onClick={() => setIsLogin(!isLogin)}>
              {isLogin ? " Зарегистрируйтесь" : " Войдите"}
            </span>
          </p>
        </div>
      </div>
    </div>
  );
};

export default AuthForm;