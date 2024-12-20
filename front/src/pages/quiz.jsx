import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";

const Quiz = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const [quizData, setQuizData] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [answers, setAnswers] = useState({});
    const [customAnswers, setCustomAnswers] = useState({});

    useEffect(() => {
        const fetchQuizData = async () => {
            try {
                const response = await fetch(`http://localhost:8080/api/prefix/quiz/${id.slice(1)}`);
                if (!response.ok) {
                    throw new Error("Опрос не найден");
                }

                const data = await response.json();
                const quiz = data.quiz;

                if (!quiz.questions || quiz.questions.length === 0) {
                    throw new Error("У опроса нет вопросов");
                }

                const invalidQuestions = quiz.questions.filter(
                    (question) => !question.answers || question.answers.length === 0
                );
                if (invalidQuestions.length > 0) {
                    throw new Error("У одного или нескольких вопросов отсутствуют варианты ответа");
                }

                setQuizData(quiz);
            } catch (err) {
                setError(err.message);
            } finally {
                setLoading(false);
            }
        };

        fetchQuizData();
    }, [id]);

    const handleAnswerChange = (questionId, answer) => {
        setAnswers((prev) => ({
            ...prev,
            [questionId]: answer,
        }));
    };

    const handleCustomInputChange = (questionId, value) => {
        setCustomAnswers((prev) => ({
            ...prev,
            [questionId]: value,
        }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        const result = Object.keys(answers).map((questionId) => ({
            id: parseInt(questionId, 10),
            reply: customAnswers[questionId] || answers[questionId],
        }));

        try {
            const response = await fetch(`http://localhost:8080/api/prefix/quiz/${id.slice(1)}/results`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(result),
            });

            if (response.ok) {
                alert("Опрос завершён. Спасибо за участие!");
            } else {
                throw new Error("Ошибка отправки результатов");
            }
        } catch (err) {
            console.error("Error submitting results:", err);
            alert("Произошла ошибка. Попробуйте ещё раз.");
        }
    };

    if (loading) {
        return <div className="quiz-loading">Загрузка...</div>;
    }

    if (error) {
        return (
            <div className="quiz-error">
                <p>Ошибка: {error}</p>
                <button className="quiz-back-button" onClick={() => navigate("/panel")}>
                    Вернуться к списку опросов
                </button>
            </div>
        );
    }

    return (
        <div className="quiz-container">
            <h1 className="quiz-title">{quizData.title}</h1>
            <form onSubmit={handleSubmit} className="quiz-form">
                {quizData.questions.map((question) => (
                    <QuestionComponent
                        key={question.id}
                        question={question}
                        answers={answers}
                        customAnswers={customAnswers}
                        onAnswerChange={handleAnswerChange}
                        onCustomInputChange={handleCustomInputChange}
                    />
                ))}
                <button type="submit" className="quiz-submit">
                    Отправить
                </button>
            </form>
        </div>
    );
};

const QuestionComponent = ({ question, answers, customAnswers, onAnswerChange, onCustomInputChange }) => {
    const isOtherSelected = (questionId) => answers[questionId] === "Другой";

    const handleOptionChange = (event) => {
        const { name, value } = event.target;
        onAnswerChange(name, value);
    };

    const handleCustomInputChange = (event) => {
        const { name, value } = event.target;
        onCustomInputChange(name, value);
    };

    return (
        <div className="quiz-question">
            <h3>{question.question}</h3>
            {question.answers.map((answer, idx) => (
                <label key={idx} className="quiz-answer">
                    <input
                        type="radio"
                        name={question.id.toString()}
                        value={answer}
                        checked={answers[question.id] === answer}
                        onChange={handleOptionChange}
                    />
                    {answer}
                </label>
            ))}
            <label className="quiz-answer">
                <input
                    type="radio"
                    name={question.id.toString()}
                    value="Другой"
                    checked={answers[question.id] === "Другой"}
                    onChange={handleOptionChange}
                />
                Другое
            </label>
            <input
                type="text"
                placeholder="Введите ваш ответ"
                disabled={!isOtherSelected(question.id)}
                value={customAnswers[question.id] || ""}
                onChange={handleCustomInputChange}
                name={question.id.toString()}
                className="quiz-input"
            />
        </div>
    );
};

export default Quiz;
