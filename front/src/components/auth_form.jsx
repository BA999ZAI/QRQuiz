import React, { useState, useContext } from "react";
import { AuthContext } from "../auth/AuthContext";
import { useNavigate } from "react-router-dom";

const AuthForm = () => {
  const [isLogin, setIsLogin] = useState(false);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { isAuthenticated, login, logout, userId } = useContext(AuthContext);
  const navigate = useNavigate()

  const Login = async () => {
    const response = await fetch(`http://localhost:8080/api/prefix/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: email,
        password: password,
      })
    });

    if (!response.ok) {
      alert("Пользователь не найден");
      return
    }

    const data = await response.json();
    const user = data.user;
    console.log("Успешная авторизация:", user);


    if (response.status == 200) {
      login(user.id, user.email)
      navigate("/panel")
      return
    }

    alert("Произошла ошибка, возможно ввели неверные данные")
  }

  const Register = async () => {
    const response = await fetch(`http://localhost:8080/api/prefix/register`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: email,
        password: password,
      })
    });

    if (!response.ok) {
      alert("Произошла ошибка при создание пользователя");
    }

    const data = await response.json();
    const user = data.user;
    console.log("Успешная регистрация:", data);

    if (response.status == 201) {
      login(user.id, user.email)
      navigate("/panel")
      return
    }

    alert("Произошла ошибка при создание пользователя")
  }

  const auth = (e) => {
    e.preventDefault();
    isLogin ? Login() : Register();
  }

  return (

    // <div className="auth-container">
    //   <div className="auth-box">
    //     <h2>{isLogin ? "Авторизация" : "Регистрация"}</h2>
    //     <form onSubmit={auth}>
    //       <div className="form-group">
    //         <label htmlFor="email">Email</label>
    //         <input
    //           type="email"
    //           id="email"
    //           placeholder="Введите email"
    //           value={email}
    //           onChange={(e) => setEmail(e.target.value)}
    //           required
    //         />
    //       </div>
    //       <div className="form-group">
    //         <label htmlFor="password">Пароль</label>
    //         <input
    //           type="password"
    //           id="password"
    //           placeholder="Введите пароль"
    //           value={password}
    //           onChange={(e) => setPassword(e.target.value)}
    //           required
    //         />
    //       </div>
    //       <button type="submit" className="auth-button">
    //         {isLogin ? "Войти" : "Зарегистрироваться"}
    //       </button>
    //     </form>
    //     <div className="toggle-form">
    //       <p>
    //         {isLogin ? "Нет аккаунта?" : "Уже есть аккаунт?"}
    //         <span onClick={() => setIsLogin(!isLogin)}>
    //           {isLogin ? " Зарегистрируйтесь" : " Войдите"}
    //         </span>
    //       </p>
    //     </div>
    //   </div>
    // </div>
    <div className="auth-container">
      <h2>{isLogin ? "ВХОД В ЛИЧНЫЙ КАБИНЕТ" : "РЕГИСТРАЦИЯ"}</h2>

      <form>
        <input type="email" placeholder="Введите email" value={email} onChange={(e) => setEmail(e.target.value)} />

        <input type="password" placeholder={isLogin ? "Пароль" : "Придумайте пароль"} value={password} onChange={(e) => setPassword(e.target.value)} />
        {isLogin ? null : <input type="password" placeholder="Подтвердите пароль" value={password} onChange={(e) => setPassword(e.target.value)} />}
        <button onClick={auth}>{isLogin ? "Войти" : "Зарегистрироваться"}</button>
      </form>


      <p>
        {isLogin ? "У Вас нет аккаунта?" : "У Вас есть аккаунт?"}
        <span onClick={() => setIsLogin(!isLogin)}>
          {isLogin ? " Зарегистрируйтесь" : " Войдите"}
        </span>
      </p>
    </div>
  );
};

export default AuthForm;