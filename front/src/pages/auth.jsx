import Header from "../components/header"
import AuthForm from "../components/auth_form"
import Footer from "../components/footer"

const Auth = () => {
    return (
        <div className="auth-page">
            <Header />

            <AuthForm />
            <Footer />
        </div>
    )
}

export default Auth