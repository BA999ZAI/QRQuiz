import React, { useContext, useEffect } from "react";
import { BrowserRouter, Route, Routes, Navigate } from "react-router-dom";
import Auth from "./pages/auth";
import Base from "./pages/base";
import CreateQuiz from "./pages/create_quiz";
import Quiz from "./pages/quiz";
import UserPanel from "./pages/user_panel";
import { AuthContext } from "./auth/AuthContext";

const ProtectedRoute = ({ element }) => {
  const { isAuthenticated } = useContext(AuthContext);

  // Если пользователь не авторизован, перенаправляем на страницу авторизации
  if (!isAuthenticated) {
    return <Navigate to="/authorization" />;
  }

  // Если авторизован, показываем элемент
  return element;
};

function App() {
  const { isAuthenticated } = useContext(AuthContext);

  // Проверяем состояние авторизации при загрузке приложения
  useEffect(() => {
    console.log("Состояние isAuthenticated:", isAuthenticated);
  }, [isAuthenticated]);

  return (
    <BrowserRouter>
      <Routes>
        {/* Главная страница */}
        <Route path="/" element={<Base />} />

        {/* Страница авторизации */}
        <Route
          path="/authorization"
          element={isAuthenticated ? <Navigate to="/panel" /> : <Auth />}
        />

        {/* Защищенные маршруты */}
        <Route
          path="/create_quiz"
          element={<ProtectedRoute element={<CreateQuiz />} />}
        />
        <Route path="/quiz/:id" element={<Quiz />} />
        <Route
          path="/panel"
          element={<ProtectedRoute element={<UserPanel />} />}
        />

        {/* Если маршрут не найден, перенаправляем на главную страницу */}
        <Route path="*" element={<Navigate to="/" />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;