import React, { useEffect, useState, useContext } from "react";
import { Link } from "react-router-dom";
import { AuthContext } from "../auth/AuthContext";
import { useNavigate } from "react-router-dom";
import "../index.css";

const QuizList = () => {
    const [quizzes, setQuizzes] = useState([])
    const { isAuthenticated, userId, logout } = useContext(AuthContext);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [imageData, setImageData] = useState({});
    const navigate = useNavigate();

    const [isUpQr, setIsUpQr] = useState(false);


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

    const upperQrCode = (quiz_id) => {
        const qrCode = document.getElementById(`qr_${quiz_id}`);
        if (qrCode.style.position == "relative" && isUpQr) {
            return
        }

        if (!isUpQr) {
            qrCode.style.position = "absolute";
            qrCode.style.top = "5%";
            qrCode.style.left = "25%";
            qrCode.style.zIndex = "1000";
            qrCode.style.width = "auto";
            qrCode.style.height = "80vh";
            qrCode.style.objectFit = "cover";
            qrCode.style.borderRadius = "10px";
            qrCode.style.boxShadow = "0px 0px 10px rgba(0, 0, 0, 0.5)";
            qrCode.style.transform = "scale(1.2)";
            qrCode.style.opacity = "1";
            setIsUpQr(true);
            return
        }

        qrCode.style.position = "relative";
        qrCode.style.top = "0";
        qrCode.style.left = "0";
        qrCode.style.zIndex = "0";
        qrCode.style.width = "60px";
        qrCode.style.height = "60px";
        qrCode.style.objectFit = "cover";
        qrCode.style.borderRadius = "10px";
        qrCode.style.boxShadow = "0px 0px 10px rgba(0, 0, 0, 0.5)";
        qrCode.style.transform = "scale(1)";
        qrCode.style.opacity = "1";
        setIsUpQr(false);
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
                    <img id={`qr_${quiz.id}`} onClick={() => upperQrCode(quiz.id)} className="qr cursor-pointer" src={`data:image/png;base64,${imageData[quiz.id]}`} alt="qr-code" />
                    <button className="cursor-pointer" onClick={deleteQuiz(quiz.id)}>
                        Удалить
                    </button>
                </div>
            )) : <div>Нет опросов</div>}
        </div>
    );
};

export default QuizList;