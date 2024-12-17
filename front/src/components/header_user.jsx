import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom"
import LogoutButton from "./logout_button"
import UserProfileModal from "./user_profile_modal"


const HeaderUser = () => {
    const [isProfileOpen, setIsProfileOpen] = useState(false);

    const navigate = useNavigate();

    const handleLogout = () => {
        // Логика выхода пользователя
        localStorage.removeItem("token"); // Удаляем токен
        navigate("/"); // Редирект на главную страницу
    };

    return (
        <header className="header">
            <div className="profile-section" onClick={() => setIsProfileOpen(true)}>
                <img
                    src="https://via.placeholder.com/40"
                    alt="Profile"
                    className="profile-avatar"
                />
                <span>Профиль</span>
            </div>

            <div className="div-auth">
                <Link to="/panel" className="button-auth auth mr-10">Панель</Link>

                <LogoutButton onClick={handleLogout} />
            </div>

            {isProfileOpen && (
                <UserProfileModal onClose={() => setIsProfileOpen(false)} />
            )}
        </header>
    )
}

export default HeaderUser