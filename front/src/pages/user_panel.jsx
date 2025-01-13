import React from "react";
import QuizList from "../components/quiz_list";
import AddQuizButton from "../components/add_quiz_button";
import HeaderUser from "../components/header_user";
import { useNavigate } from "react-router-dom";
import Header from "../components/header";
import Footer from "../components/footer";

const UserPanel = () => {
    const navigate = useNavigate();

    return (
        <div className="dashboard-container">
            <Header page={"panel"} />


            {/* <main className="dashboard-main">
                <h1>Ваши опросы</h1>
                <AddQuizButton onClick={() => navigate("/create_quiz")} />
                </main> */}
            <QuizList />
            
            <Footer />
        </div>
    );
};

export default UserPanel;