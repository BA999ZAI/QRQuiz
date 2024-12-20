import React, { useEffect, useState, useContext } from "react";
import { Link, useNavigate } from "react-router-dom";
import LogoutButton from "./logout_button";
import { AuthContext } from "../auth/AuthContext";
import { use } from "react";
import { cache } from "react";

const HeaderUser = () => {
    const [isProfileOpen, setIsProfileOpen] = useState(false);
    const { isAuthenticated, userId, logout } = useContext(AuthContext);
    const [email, setEmail] = useState("");
    const navigate = useNavigate();



    useEffect(() => {
        const getuserbyid = async (id) => {
            try {
                const response = await fetch(`http://localhost:8080/api/prefix/user/${id}`, {
                    method: "GET",
                    headers: {
                        "Content-Type": "application/json",
                    },
                });
                const data = await response.json();
                const resp = data.user
                setEmail(resp.email)
                console.log(resp.email)
            } catch (error) {
                console.error("Error fetching user:", error);
            }
        };
    
        getuserbyid(userId);

    }, [userId]);

    const handleLogout = () => {
        logout();
        navigate("/");
    };

    return (
        <header className="header">
            <div className="profile-section" onClick={() => setIsProfileOpen(true)}>
                <img
                    src="https://png.pngtree.com/png-clipart/20191122/original/pngtree-user-icon-isolated-on-abstract-background-png-image_5192004.jpg"
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
                                src="https://png.pngtree.com/png-clipart/20191122/original/pngtree-user-icon-isolated-on-abstract-background-png-image_5192004.jpg"
                                alt="Profile"
                                className="profile-avatar"
                            />
                            <p>Email: {email}</p>
                        </div>
                        <button className="close-button" onClick={() => setIsProfileOpen(false)}>
                            Закрыть
                        </button>
                    </div>
                </div>
            )}
        </header>
    );
};

export default HeaderUser;