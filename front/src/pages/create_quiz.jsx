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

  const [answerAmount, setAnswerAmount] = useState([2])
  const [answers, setAnswers] = useState([[]])

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

  const addQuestion = () => {
    setQuestionAmount(questionAmount + 1);
    setAnswerAmount([...answerAmount, 1]); // Добавляем один ответ для нового вопроса
    setAnswers([...answers, []]); // Добавляем пустой массив ответов для нового вопроса
  };

  // Обработчик удаления вопроса
  const removeQuestion = () => {
    if (questionAmount > 1) {
      setQuestionAmount(questionAmount - 1);
      setAnswerAmount(answerAmount.slice(0, -1)); // Удаляем последний элемент из массива answerAmount
      setAnswers(answers.slice(0, -1)); // Удаляем последний массив ответов
    }
  };


  const addAnswer = (questionIndex) => {
    const newAnswerAmount = [...answerAmount];
    newAnswerAmount[questionIndex] += 1; // Увеличиваем количество ответов для конкретного вопроса
    setAnswerAmount(newAnswerAmount);

    const newAnswers = [...answers];
    newAnswers[questionIndex].push(""); // Добавляем пустой ответ для конкретного вопроса
    setAnswers(newAnswers);
  };

  // Обработчик удаления ответа
  const removeAnswer = (questionIndex) => {
    if (answerAmount[questionIndex] > 1) {
      const newAnswerAmount = [...answerAmount];
      newAnswerAmount[questionIndex] -= 1; // Уменьшаем количество ответов для конкретного вопроса
      setAnswerAmount(newAnswerAmount);

      const newAnswers = [...answers];
      newAnswers[questionIndex].pop(); // Удаляем последний ответ для конкретного вопроса
      setAnswers(newAnswers);
    }
  };
  
  return (
    <div className="create-quiz-container">
      <HeaderUser />

      <h1>Создать опрос</h1>
      <form className="form-quiz" onSubmit={handleSubmit}>
        {/* tilte */}
        <input type="text" placeholder="Название опроса" required />
        {/* questions */}
        {Array.from({ length: questionAmount }).map((_, index) => (
          <div className="d-flex wrap space-between div-question" key={index}>
            <input className="input-question" required type="text" placeholder={`Вопрос ${index + 1}`} />
            {Array.from({ length: answerAmount[index] }).map((_, answerIndex) => (
              <input
                required
                key={answerIndex}
                type="text"
                placeholder={`Ответ ${answerIndex + 1}`}
              />
            ))}
            <div className="d-flex space-between">
              <button className="button-auth auth" type="button" onClick={() => addAnswer(index)}>
                Добавить ответ
              </button>
              <button className="button-auth auth" onClick={() => removeAnswer(index)}>
                Удалить ответ
              </button>
            </div>
          </div>
        ))}

        <div className="my-10 d-flex space-between">
          <button className="button-auth auth" type="button" onClick={addQuestion}>
            Добавить вопрос
          </button>
          <button className="button-auth auth" onClick={removeQuestion}>
            Удалить вопрос
          </button>
        </div>


        {/* time to live */}
        <h3>Дата и время окончания опроса:</h3>
        <div className="d-flex space-between align-center">
          <label>Дата:</label>
          <input
            type="date"
            value={date}
            onChange={(e) => setDate(e.target.value)}
          />
          <label>Время:</label>
          <input
            type="time"
            value={time}
            onChange={(e) => setTime(e.target.value)}
          />
        </div>
        <button className="button-auth auth mx-auto" type="submit">Создать</button>
      </form>
    </div>
  );
};

export default CreateQuiz;