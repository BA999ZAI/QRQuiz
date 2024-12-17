import React from "react";
import QuizList from "../components/quiz_list";
import AddQuizButton from "../components/add_quiz_button";
import HeaderUser from "../components/header_user";
import { useNavigate } from "react-router-dom";

const UserPanel = () => {
    const navigate = useNavigate();

    return (
        <div className="dashboard-container">
            <HeaderUser />

            <main className="dashboard-main">
                <h1>Ваши опросы</h1>
                <QuizList />
                <AddQuizButton onClick={() => navigate("/create_quiz")} />
            </main>
        </div>
    );
};

export default UserPanel;