import { Link, useNavigate } from "react-router-dom"
import LogoutButton from "./logout_button"
import { useContext } from "react"
import { AuthContext } from "../auth/AuthContext"

const Header = () => {
    const { isAuthenticated, userId, logout } = useContext(AuthContext);

    const navigate = useNavigate();

    const handleLogout = () => {
        logout()
        navigate("/")
    };

    return (
        <header className="header">
            <Link className="logo" to="/">
                QRQuiz
            </Link>

            <div className="div-auth">
                <Link className="button-auth auth mr-10" to="/authorization">
                    Авторизация
                </Link>
                <LogoutButton onClick={handleLogout} />
            </div>
        </header>
    )
}

export default Header