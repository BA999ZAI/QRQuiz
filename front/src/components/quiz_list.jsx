import React from "react";

const quizzes = [
    { id: 1, title: "Опрос о React", description: "Как вы оцениваете React?" },
    { id: 2, title: "Опрос о JavaScript", description: "Ваш опыт с JS?" },
];

const QuizList = () => {
    return (
        <div className="quiz-list">
            {quizzes.map((quiz) => (
                <div key={quiz.id} className="quiz-item">
                    <h3>{quiz.title}</h3>
                    <p>{quiz.description}</p>
                </div>
            ))}
        </div>
    );
};

export default QuizList;