import React, { useEffect, useState, useContext } from "react";
import { Link, useNavigate } from "react-router-dom";
import { AuthContext } from "../auth/AuthContext";
import "../index.css";

const QuizList = () => {
    const [quizzes, setQuizzes] = useState([]);
    const { isAuthenticated, userId } = useContext(AuthContext);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [imageData, setImageData] = useState({});
    const navigate = useNavigate();

    const [isUpQr, setIsUpQr] = useState(false);

    useEffect(() => {
        const fetchQuizData = async () => {
            try {
                const response = await fetch(`http://localhost:8080/api/prefix/quiz/user/${userId}`);
                if (!response.ok) {
                    alert("Опрос не найден");
                    return;
                }
                const data = await response.json();
                const quizzesData = data.quizzes;
                setQuizzes(quizzesData);

                quizzesData.map(async (quiz) => {
                    const imageResponse = await fetch(`http://localhost:8080/api/prefix/quiz/${quiz.id}`);
                    if (!imageResponse.ok) {
                        alert("Ошибка при загрузке изображения");
                    }
                    const responseImage = await imageResponse.json();
                    const imageUrl = responseImage.qr;
                    setImageData((prevImageData) => ({
                        ...prevImageData,
                        [quiz.id]: imageUrl,
                    }));
                });
            } catch (err) {
                setError(err.message);
            } finally {
                setLoading(false);
            }
        };

        fetchQuizData();
    }, [userId]);

    if (loading) return <div>Загрузка...</div>;
    if (error) return <div>Ошибка: {error}</div>;

    const deleteQuiz = (quizId) => async () => {
        try {
            const response = await fetch(`http://localhost:8080/api/prefix/quiz/${quizId}`, { method: "DELETE" });
            if (!response.ok) alert("Ошибка при удалении опроса");
            setQuizzes((prevQuizzes) => prevQuizzes.filter((quiz) => quiz.id !== quizId));
        } catch (err) {
            alert("Произошла ошибка при удалении опроса");
        }
    };

    const upperQrCode = (quiz_id) => {
        const qrCode = document.getElementById(`qr_${quiz_id}`);
        if (!isUpQr) {
            qrCode.style = "position: absolute; top: 5%; left: 25%; z-index: 1000; width: auto; height: 80vh; transform: scale(1.2); box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.5);";
            setIsUpQr(true);
        } else {
            qrCode.style = "position: relative; width: 60px; height: 60px; z-index: 0; transform: scale(1);";
            setIsUpQr(false);
        }
    };

    return (
        <div className="quiz-list">
            {quizzes.length > 0 ? (
                quizzes.map((quiz) => (
                    <div key={quiz.id} className="quiz-item">
                        <div className="quiz-item-content">
                            <h3><Link to={`/quiz/:${quiz.id}`}>{quiz.title}</Link></h3>
                            <p>{quiz.status ? "Время вышло" : "Опрос идёт"}</p>
                        </div>
                        <img
                            id={`qr_${quiz.id}`}
                            onClick={() => upperQrCode(quiz.id)}
                            className="qr cursor-pointer"
                            src={`data:image/png;base64,${imageData[quiz.id]}`}
                            alt="qr-code"
                        />
                        {quiz.status && (
                            <button className="btn btn-results" onClick={() => navigate(`/results/${quiz.id}`)}>
                                Посмотреть результаты
                            </button>
                        )}
                        <button className="btn btn-delete" onClick={deleteQuiz(quiz.id)}>
                            Удалить
                        </button>
                    </div>
                ))
            ) : (
                <div>Нет опросов</div>
            )}
        </div>
    );
};

export default QuizList;