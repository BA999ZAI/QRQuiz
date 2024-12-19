import React, { useEffect, useState, useContext } from "react";
import { Link, useNavigate } from "react-router-dom"
import LogoutButton from "./logout_button"
import { AuthContext } from "../auth/AuthContext";

const HeaderUser = () => {
    const [isProfileOpen, setIsProfileOpen] = useState(false);
    const { isAuthenticated, userId, logout } = useContext(AuthContext);
    const [user, setUser] = useState({})

    const navigate = useNavigate();

    useEffect(() => {
        const fetchUserData = async () => {
            try {
                const response = await fetch(`http://localhost:8080/api/prefix/user/${userId}`);
                if (!response.ok) {
                    throw new Error("Пользователь не найден");
                }
                const data = await response.json();
                setUser(data);
            } catch (err) {
                prompt(err.message);
            }
        };

        fetchUserData();
    }, [userId]);

    const handleLogout = () => {
        logout()
        navigate("/")
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
                <div className="modal-overlay">
                    <div className="modal-content">
                        <h2>Профиль пользователя</h2>
                        <div className="profile-info">
                            <img
                                src="https://via.placeholder.com/100"
                                alt="Profile"
                                className="profile-avatar"
                            />
                            <p>Email: {user.email}</p>
                        </div>
                        <button className="close-button" onClick={setIsProfileOpen(false)}>
                            Закрыть
                        </button>
                    </div>
                </div>
            )}
        </header>
    )
}

export default HeaderUser