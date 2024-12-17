import React from "react";
import { useNavigate } from "react-router-dom";
import HeaderUser from "../components/header_user";

const CreateQuiz = () => {
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();
    // Логика создания опроса
    alert("Опрос создан!");
    navigate("/panel"); // Возвращаемся на панель пользователя
  };

  return (
    <div className="create-quiz-container">
      <HeaderUser />

      <h1>Создать опрос</h1>
      <form onSubmit={handleSubmit}>
        <input type="text" placeholder="Название опроса" required />
        <textarea placeholder="Описание опроса" required />
        <button type="submit">Создать</button>
      </form>
    </div>
  );
};

export default CreateQuiz;