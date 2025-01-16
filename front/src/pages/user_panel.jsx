import React from "react";
import QuizList from "../components/quiz_list";
import { useNavigate } from "react-router-dom";
import Header from "../components/header";
import Footer from "../components/footer";

const UserPanel = () => {
    const navigate = useNavigate();

    return (
        <div className="dashboard-container">
            <Header page={"panel"} />

            <QuizList />

            <Footer />
        </div>
    );
};

export default UserPanel;