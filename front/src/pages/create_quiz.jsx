import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import HeaderUser from "../components/header_user";
import { formatISO } from "date-fns";
import { AuthContext } from "../auth/AuthContext";
import { useContext } from "react";

const CreateQuiz = () => {
  const navigate = useNavigate();
  const [title, setTitle] = useState([])
  const [questionAmount, setQuestionAmount] = useState(1)
  const [questions, setQuestions] = useState([])
  const [timeToLive, setTimeToLive] = useState()
  const [date, setDate] = useState("");
  const [time, setTime] = useState("");
  const { isAuthenticated, login, logout, userId } = useContext(AuthContext);


  const handleSubmit = (e) => {
    e.preventDefault();
    // Объединяем дату и время в один объект Date
    const dateTime = new Date(`${date}T${time}`);

    // Преобразуем в формат ISO с временной зоной
    const isoDateTime = formatISO(dateTime, {
      format: "extended", // Используем расширенный формат
      representation: "complete", // Включаем временную зону
    });

    console.log("Отправляем на бэкенд:", isoDateTime);

    // Пример отправки на бэкенд (замените на ваш API)
    fetch("http://localhost:8080/api/prefix/quiz", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ 
        title: title,
        questions: questions,
        time_to_live: isoDateTime,
        user_id: userId,
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log("Ответ от сервера:", data);
      })
      .catch((error) => {
        console.error("Ошибка при отправке данных:", error);
      });
    navigate("/panel"); // Возвращаемся на панель пользователя
  };

  return (
    <div className="create-quiz-container">
      <HeaderUser />

      <h1>Создать опрос</h1>
      <form onSubmit={handleSubmit}>
        {/* tilte */}
        <input type="text" placeholder="Название опроса" required />
        {/* questions */}
        {Array.from({ length: questionAmount }).map((_, index) => (
          <div key={index}>
            <input required type="text" placeholder={`Вопрос ${index + 1}`} />
          </div>
        ))}

        <div className="my-10 d-flex space-between">
          <button className="button-auth auth" type="button" onClick={() => setQuestionAmount(questionAmount + 1)}>
            Добавить вопрос
          </button>
          <button className="button-auth auth" onClick={() => setQuestionAmount(questionAmount - 1)}>
            Удалить вопрос
          </button>
        </div>


        {/* time to live */}
        <div>
          <label>Дата:</label>
          <input
            type="date"
            value={date}
            onChange={(e) => setDate(e.target.value)}
            required
          />
        </div>
        <div>
          <label>Время:</label>
          <input
            type="time"
            value={time}
            onChange={(e) => setTime(e.target.value)}
            required
          />
        </div>
        <button className="button-auth auth mx-auto" type="submit">Создать</button>
      </form>
    </div>
  );
};

export default CreateQuiz;