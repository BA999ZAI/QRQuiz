import Header from "../components/header"
import AuthForm from "../components/auth_form"

const Auth = () => {
    return (
        <div className="d-flex flex-column h-full">
            <Header />

            <AuthForm />
        </div>
    )
}

export default Auth