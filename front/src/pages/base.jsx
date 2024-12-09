import Header from "../components/header"

const Base = () => {
    return (
        <div className="d-flex flex-column vh-100">
            <Header />
            
            <div className="d-flex mx-60 mt-60 flex-column">
                <p className="mx-auto">Нужно войти чтобы создать опрос.</p>
            </div>
        </div>
    )
}

export default Base