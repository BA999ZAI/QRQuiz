import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

const Quiz = () => {
    const { id } = useParams(); // Получаем ID из URL
    const [quizData, setQuizData] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    // Функция для загрузки данных опроса по ID
    useEffect(() => {
        const fetchQuizData = async () => {
            try {
                const response = await fetch(`http://localhost:8080/api/prefix/quiz/${id.slice(1)}`); // Замените на ваш API endpoint
                if (!response.ok) {
                    throw new Error("Опрос не найден");
                }
                const data = await response.json();

                const quiz = data.quiz 
                console.log(quiz)
                setQuizData(quiz);
            } catch (err) {
                setError(err.message);
            } finally {
                setLoading(false);
            }
        };

        fetchQuizData();
    }, [id]);

    if (loading) {
        return <div>Загрузка...</div>;
    }

    if (error) {
        return <div>Ошибка: {error}</div>;
    }

    return (
        <div>
            <h1>{quizData.title}</h1>
            <form>
                {/* {quizData.questions.map((question, index) => (
                    <div key={index}>
                        <h3>{question.text}</h3>
                        {question.options.map((option, optionIndex) => (
                            <div key={optionIndex}>
                                <input
                                    type="radio"
                                    id={`option-${optionIndex}`}
                                    name={`question-${index}`}
                                    value={option}
                                />
                                <label htmlFor={`option-${optionIndex}`}>{option}</label>
                            </div>
                        ))}
                    </div>
                ))} */}
                <button type="submit">Отправить</button>
            </form>
        </div>
    );
};

export default Quiz;