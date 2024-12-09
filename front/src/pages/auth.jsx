import Header from "../components/header"


const Auth = () => {
    return (
        <div className="d-flex flex-column">
            <Header />

            <div className="d-flex mx-60 mt-200">
                <form className="auth-form d-flex flex-column mx-auto border-1">
                    <p>Авторизация</p>

                    <input type="text" placeholder="email" className="" />
                    <input type="password" placeholder="password" className="" />

                    <button>
                        Отправить
                    </button>
                </form>
            </div>

        </div>
    )
}

export default Auth