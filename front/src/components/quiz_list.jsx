import React, { useEffect, useState, useContext } from "react";
import { Link } from "react-router-dom";
import { AuthContext } from "../auth/AuthContext";
import { useNavigate } from "react-router-dom";

const QuizList = () => {
    const [quizzes, setQuizzes] = useState([])
    const { isAuthenticated, userId, logout } = useContext(AuthContext);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchQuizData = async () => {
            try {
                console.log(userId)
                const response = await fetch(`http://localhost:8080/api/prefix/quiz/user/${userId}`);
                if (!response.ok) {
                    alert("Опрос не найден");
                }
                const data = await response.json();
                const quizzesData = data.quizzes;
                console.log(quizzesData)
                setQuizzes(quizzesData);
            } catch (err) {
                setError(err.message);
            } finally {
                setLoading(false);
            }
        };

        fetchQuizData();
    }, [userId]);

    if (loading) {
        return <div>Загрузка...</div>;
    }

    if (error) {
        return <div>Ошибка: {error}</div>;
    }

    const deleteQuiz = (quizId) => async () => {
        try {
            const response = await fetch(`http://localhost:8080/api/prefix/quiz/${quizId}`, {
                method: "DELETE",
            });
            if (!response.ok) {
                alert("Ошибка при удалении опроса");
            }
            setQuizzes((prevQuizzes) => prevQuizzes.filter((quiz) => quiz.id !== quizId));
        } catch (err) {
            alert("Произошла ошибка при удалении опроса");
        }
    }

    return (
        <div className="quiz-list">
            {quizzes ? quizzes.map((quiz) => (
                <div key={quiz.id} className="quiz-item">
                    <div className="quiz-item-content">
                        <h3>{quiz.title}</h3>
                        <p>{quiz.status ? "Время вышло" : "Опрос идёт"}</p>
                    </div>

                    <Link className="cursor-pointer" to={`/quiz/:${quiz.id}`}>Ссылка: {quiz.link_to_quiz}</Link>
                    <button className="cursor-pointer" onClick={deleteQuiz(quiz.id)}>
                        Удалить
                    </button>
                </div>
            )) : <div>Нет опросов</div>}
        </div>
    );
};

export default QuizList;