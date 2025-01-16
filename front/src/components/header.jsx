import { Link, useNavigate } from "react-router-dom"
import { useContext, useEffect } from "react"
import { AuthContext } from "../auth/AuthContext"

const Header = ({ page }) => {
    const { isAuthenticated, userId, logout } = useContext(AuthContext);

    const navigate = useNavigate();

    const handleLogout = () => {
        logout()
        navigate("/")
    };

    const handleAuth = () => {
        navigate("/authorization")
    }

    useEffect(() => {
        if (page === "panel") {
            const doc = document.getElementById("panel")
            doc.style.color = "white"
            doc.style.backgroundColor = "#FB9461"
        } else {
            const doc = document.getElementById("base")
            doc.style.color = "white"
            doc.style.backgroundColor = "#FB9461"
        }
    }, [1])

    return (
        <header className="header">
            <a className="logo">QRQuiz</a>

            {/* Switch Главная||Мои опросы||Профиль */}
            <div className="switch-page-div">
                <Link to="/" id="base" className="switch-page-link">Главная</Link>
                <Link to="/panel" id="panel" className="switch-page-link">Мои опросы</Link>
            </div>


            {/* Если авторизован Выход, иначе Зарегистрироваться */}
            {isAuthenticated ? (
                <button onClick={handleLogout} className="auth-button">
                    Выход
                </button>
            ) : (
                <button onClick={handleAuth} className="auth-button">
                    Зарегистрироваться
                </button>
            )}
        </header>
    )
}

export default Header