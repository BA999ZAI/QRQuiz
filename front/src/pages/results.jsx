import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

const SurveyResults = () => {
    const { id: quizId } = useParams(); // Получаем quizId из URL параметров
    const [results, setResults] = useState({});
    const [questions, setQuestions] = useState([]);

    useEffect(() => {
        fetch(`/api/quiz/${quizId}`)
            .then((res) => res.json())
            .then((data) => {
                setQuestions(data.quiz.questions);
                setResults(data.results);
            })
            .catch((error) => console.error("Error fetching results:", error));
    }, [quizId]);

    const calculatePercent = (answers) => {
        const total = Object.values(answers).reduce((sum, count) => sum + count, 0);
        return Object.entries(answers).map(([answer, count]) => ({
            answer,
            percent: ((count / total) * 100).toFixed(2),
        }));
    };

    return (
        <div className="survey-results">
            {questions.map((q) => (
                <div key={q.id} className="question-block">
                    <h3>{q.question}</h3>
                    <div className="circles-container">
                        {results[q.id] &&
                            calculatePercent(results[q.id]).map(({ answer, percent }) => (
                                <div
                                    key={answer}
                                    className="circle"
                                    style={{ width: `${percent * 2}px`, height: `${percent * 2}px` }}
                                >
                  <span className="circle-text">
                    {answer} - {percent}%
                  </span>
                                </div>
                            ))}
                    </div>
                </div>
            ))}
        </div>
    );
};

export default SurveyResults;
