import { Link } from "react-router-dom"


const Header = () => {
    return (
        <header className="header">
            <Link className="logo" to="/">
                QRQuiz
            </Link>

            <div className="div-auth">
                <Link className="button-auth auth mr-10" to="/authorization">
                    Войти
                </Link>
                <Link className="button-auth reg" to="/registration">
                    Регистрация
                </Link>
            </div>
        </header>
    )
}

export default Header