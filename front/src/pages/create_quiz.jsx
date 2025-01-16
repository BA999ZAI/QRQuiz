import React, { useState, useContext } from "react";
import { useNavigate } from "react-router-dom";
import HeaderUser from "../components/header_user";
import { formatISO } from "date-fns";
import { AuthContext } from "../auth/AuthContext";
import Header from "../components/header";
import Footer from "../components/footer";

const CreateQuiz = () => {
  const navigate = useNavigate();
  const [title, setTitle] = useState("");
  const [questionAmount, setQuestionAmount] = useState(1);
  const [questions, setQuestions] = useState([]);
  const [answerAmount, setAnswerAmount] = useState([2]);
  const [answers, setAnswers] = useState([[]]);
  const [date, setDate] = useState("");
  const [time, setTime] = useState("");
  const { userId } = useContext(AuthContext);

  const handleSubmit = (e) => {
    e.preventDefault();

    // Объединяем дату и время в один объект Date
    const dateTime = new Date(`${date}T${time}`);

    // Преобразуем в формат ISO с временной зоной
    const isoDateTime = formatISO(dateTime, {
      format: "extended",
      representation: "complete",
    });

    // Собираем данные из инпутов
    const questionsData = Array.from({ length: questionAmount }).map((_, index) => ({
      id: index + 1,
      question: questions[index] || "",
      answers: answers[index] || [],
    }));


    // Пример отправки на бэкенд (замените на ваш API)
    fetch("http://localhost:8080/api/prefix/quiz", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        title: title,
        questions: questionsData,
        time_to_live: isoDateTime,
        user_id: userId,
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        alert("Ответ от сервера:", data);

        navigate("/panel"); // Возвращаемся на панель пользователя
      })
      .catch((error) => {
        console.error("Ошибка при отправке данных:", error);
      });
  };

  const addQuestion = () => {
    setQuestionAmount(questionAmount + 1);
    setAnswerAmount([...answerAmount, 2]); // Добавляем два ответа для нового вопроса
    setAnswers([...answers, []]); // Добавляем пустой массив ответов для нового вопроса
  };

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

  const handleQuestionChange = (index, value) => {
    const newQuestions = [...questions];
    newQuestions[index] = value;
    setQuestions(newQuestions);
  };

  const handleAnswerChange = (questionIndex, answerIndex, value) => {
    const newAnswers = [...answers];
    newAnswers[questionIndex][answerIndex] = value;
    setAnswers(newAnswers);
  };

  return (
    <div className="create-quiz-container-page">
      <Header page={"panel"} />

      <form onSubmit={handleSubmit}>
        <h2>СОЗДАНИЕ ОПРОСА</h2>


        {/* Название опроса */}
        <input
          className="input-question-name"
          type="text"
          placeholder="Название"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
        />

        {/* Вопросы */}
        {Array.from({ length: questionAmount }).map((_, index) => (
          <div className="div-questions" key={index}>
            <input
              className="input-question-question"
              required
              type="text"
              placeholder={`Вопрос ${index + 1}`}
              value={questions[index] || ""}
              onChange={(e) => handleQuestionChange(index, e.target.value)}
            />
            {Array.from({ length: answerAmount[index] }).map((_, answerIndex) => (
              <div className="div-question-answer">
                <p></p>
                <input
                  className="input-question-answer"
                  required
                  key={answerIndex}
                  type="text"
                  placeholder={`Ответ ${answerIndex + 1}`}
                  value={answers[index][answerIndex] || ""}
                  onChange={(e) => handleAnswerChange(index, answerIndex, e.target.value)}
                />
              </div>
            ))}
            <span className="div-container-add-otvet">
              <button className="button-add-otvet" type="button" onClick={() => addAnswer(index)}>
                Добавить ответ
              </button>
              <button className="button-remove-otvet" onClick={() => removeAnswer(index)}>
                Удалить ответ
              </button>
            </span>
          </div>
        ))}

        <div className="my-10 d-flex space-between margin margin-top-20">
          <button className="button-remove-otvet margin-left-92" type="button" onClick={addQuestion}>
            Добавить вопрос
          </button>
          <button className="button-add-otvet margin-right-92" onClick={removeQuestion}>
            Удалить вопрос
          </button>
        </div>

        {/* Дата и время окончания опроса */}
        <h3 className="h3-date-time">Выберите дату и время окончания опроса</h3>
        <div className="div-date-time">
          <div>
            <label>Дата</label>
            <input
              type="date"
              value={date}
              onChange={(e) => setDate(e.target.value)}
            />
          </div>

          <div>
            <label>Время</label>
            <input
              type="time"
              value={time}
              onChange={(e) => setTime(e.target.value)}
            />
          </div>
        </div>
        <button className="button-create-quiz" type="submit">Создать</button>
      </form>

      <Footer />
    </div>
  );
};

export default CreateQuiz;