import React, { createContext, useState, useEffect } from "react";

// Создаем контекст
export const AuthContext = createContext();

// Провайдер контекста
export const AuthProvider = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(null);
  const [userId, setUserId] = useState(null); // Добавляем состояние для userId
  const [email, setEmail] = useState(null); // Добавляем состояние для userId


  // Проверяем состояние авторизации при загрузке приложения
  useEffect(() => {
    const storedUserId = localStorage.getItem("userId"); // Проверяем userId в localStorage
    const storedEmail = localStorage.getItem("email"); // Проверяем email в localStorage

    if (storedUserId) {
      setIsAuthenticated(true);
      setUserId(storedUserId);
      setEmail(storedEmail);
    } else {
      setIsAuthenticated(false);
      setUserId(null);
      setEmail(null);
    }
  }, []);

  // Функция для авторизации
  const login = (userId, email) => {
    localStorage.setItem("userId", userId); // Сохраняем userId в localStorage
    setIsAuthenticated(true);
    setUserId(userId); // Устанавливаем userId в состояние
    setEmail(email); // Устанавливаем email в состояние
  };

  // Функция для выхода
  const logout = () => {
    localStorage.removeItem("userId"); // Удаляем userId из localStorage
    setIsAuthenticated(false);
    setUserId(null); // Сбрасываем userId в состояние
    setEmail(null); // Сбрасываем email в состояние
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, login, logout, userId, email }}>
      {children}
    </AuthContext.Provider>
  );
};