import { FC } from "react";
import { Navigate } from "react-router-dom";

interface User {
    id: number,
    email: string,
    name: string,
    isEnabled: boolean
}

interface RegisterProps {
    user: User | null
}

const Register:FC<RegisterProps> = (props) => {
    if(props.user != null) {
        return <Navigate to="/" />
    }
    
    return (
        <div>reg</div>
    );
};

export default Register;