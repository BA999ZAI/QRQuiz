import { useContext } from "react"
import Footer from "../components/footer"
import Header from "../components/header"
import { useNavigate } from "react-router-dom"
import { AuthContext } from "../auth/AuthContext"


const Base = () => {
    const navigate = useNavigate();
    const { isAuthenticated, userId, logout } = useContext(AuthContext);

    const handleCreateQuiz = () => {
        if (isAuthenticated) {
            navigate("/create_quiz")
        } else {
            navigate("/authorization")
        }
    }

    return (
        <div className="d-flex flex-column vh-100">
            <Header page={"base"} />

            <div className="info-first-div">
                <h2>ОПРОСЫ, КОТОРЫЕ ОБЪЕДИНЯЮТ</h2>

                <p>
                    Платформа создана для быстрого и удобного проведения опросов,
                </p><p>
                    где каждый может собрать мнения, получить полезные данные или просто
                </p><p>
                    узнать больше о своей аудитории.
                </p>

                <button onClick={handleCreateQuiz} className="create-quiz-button">
                    Создать опрос
                </button>
            </div>

            <div className="add-new-quiz-block-base-page">
                <p onClick={handleCreateQuiz} className="add-new-quiz-block-white">+</p>
                <p onClick={handleCreateQuiz} className="add-new-quiz-block-orange">+</p>
                <p onClick={handleCreateQuiz} className="add-new-quiz-block-orange">+</p>
                <p onClick={handleCreateQuiz} className="add-new-quiz-block-white">+</p>
                <p onClick={handleCreateQuiz} className="add-new-quiz-block-orange">+</p>
                <p onClick={handleCreateQuiz} className="add-new-quiz-block-white">+</p>
            </div>

            <div className="your-opinion">
                <h2>ВАШЕ МНЕНИЕ ВАЖНО ДЛЯ НАС!</h2>

                <p>Поделитесь своими впечатлениями, идеями или замечаниями.</p>
                <p>Расскажите, что вам понравилось или что можно улучшить в нашем сайте.</p>
                <p>Ваши отзывы помогут нам стать лучше!</p>

                <form>
                    <input type="text" placeholder="Введите ваше имя" />

                    <div>
                        <label>Распишите ваши пожелания или замечания</label>
                        <textarea type="text" defaultValue="Мне все очень понравилось, но я хочу добавить" />
                    </div>

                    <button>
                        Отправить
                    </button>
                </form>
            </div>

            <Footer />
        </div>
    )
}

export default Base