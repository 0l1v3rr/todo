import { FC, useState } from "react";
import { VscKey, VscMail } from "react-icons/vsc";
import { BiUser } from "react-icons/bi";
import { Link, Navigate, useNavigate } from "react-router-dom";
import ButtonPrimary from "../components/ButtonPrimary";
import InputField from "../components/InputField";
import axios, { AxiosError } from "axios";

interface User {
    id: number,
    email: string,
    name: string,
    isEnabled: boolean
}

interface RegisterProps {
    user: User | null
}

type Settings = {
    isError: boolean, 
    errorMessage: string
}

const Register:FC<RegisterProps> = (props) => {
    let navigate = useNavigate(); 
    
    const [isBtnActive, setIsBtnActive] = useState(false);
    const [requestError, setRequestError] = useState("");
    
    // name specific things
    const [name, setName] = useState('');
    const [nameSettings, setNameSettings] = useState<Settings>({
        isError: false, 
        errorMessage: "" 
    });

    // email specific things
    const [email, setEmail] = useState('');
    const [emailSettings, setEmailSettings] = useState<Settings>({
        isError: false, 
        errorMessage: "" 
    });

    // password specific things
    const [password, setPassword] = useState("");
    const [passwordSettings, setPasswordSettings] = useState<Settings>({
        isError: false, 
        errorMessage: "" 
    });
    
    if(props.user != null) {
        return <Navigate to="/" />
    }

    const checkForActiveButton = (): void => {
        if(passwordSettings.errorMessage === "" && emailSettings.errorMessage === "" && nameSettings.errorMessage === "") {
            if(password.trim() == "" || email.trim() == "" || name.trim() == "") {
                setIsBtnActive(false);
                return;
            }

            setIsBtnActive(true);
            return;
        }

        setIsBtnActive(false);
    }

    // name validating
    const validateName = (value: string) : Settings => {
        // trimming the value
        value = value.trim();

        // if the value is empty
        if(value == "") {
            setNameSettings({
                isError: false,
                errorMessage: ""
            });
            checkForActiveButton();
            return nameSettings;
        }

        // if the value is less than 6 characters long
        if(value.length < 6) {
            setNameSettings({
                isError: true,
                errorMessage: "This name is too short."
            });
            checkForActiveButton();
            return nameSettings;
        }

        // if the value is more than 6 characters long
        if(value.length > 64) {
            setNameSettings({
                isError: true,
                errorMessage: "This name is too long."
            });
            checkForActiveButton();
            return nameSettings;
        }

        // if the value is valid
        setNameSettings({
            isError: false,
            errorMessage: ""
        });
        checkForActiveButton();
        return nameSettings;
    }

    // email validating
    const validateEmail = (value: string): Settings => {
        // trimming the value
        value = value.trim();

        // if the value is empty
        if(value == "") {
            setEmailSettings({
                isError: false,
                errorMessage: ""
            });
            checkForActiveButton();
            return emailSettings;
        }

        // matching the regex
        if(!value.match(/^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$/)) {
            setEmailSettings({
                isError: true,
                errorMessage: "This email is not valid."
            });
            checkForActiveButton();
            return emailSettings;
        }
        
        // if the value is valid
        setEmailSettings({
            isError: false,
            errorMessage: ""
        });
        checkForActiveButton();
        return emailSettings;
    };

    // password validating
    const validatePassword = (value: string): Settings => {
        // trimming the value
        value = value.trim();

        // if the value is empty
        if(value == "") {
            setPasswordSettings({
                isError: false,
                errorMessage: ""
            });
            checkForActiveButton();
            return passwordSettings;
        }

        // if the value is less than 8 characters long
        if(value.length < 8) {
            setPasswordSettings({
                isError: true,
                errorMessage: "The password is too short."
            });
            checkForActiveButton();
            return passwordSettings;
        }
        
        // if the value is valid
        setPasswordSettings({
            isError: false,
            errorMessage: ""
        });
        checkForActiveButton();
        return passwordSettings;
    };

    const handleFormSubmit = () => {
        setIsBtnActive(false);
        setRequestError("");
        
        axios.post(`${process.env.REACT_APP_BACKEND_DOMAIN}/api/v1/register`, {
            name: name,
            email: email,
            password: password
        }, { 
            withCredentials: true 
        })
        .then(() => navigate("/"))
        .catch(err => {
            const error = err as AxiosError;
            console.log(error.response?.data);
            setRequestError((error.response?.data as any).error);
        });
    }

    return (
        <div className="shadow-md bg-slate-100 border border-solid border-slate-200 rounded-md">
            <div className="text-center border-b border-b-slate-200 border-solid font-bold px-10 
                py-2 rounded-tl-md text-2xl">
                Registration
            </div>

            {requestError !== "" && 
                <div className="mx-5 my-3 px-3 py-1 text-center bg-red-200 text-red-900 rounded-md border border-solid border-red-300">
                    {requestError}
                </div>}

            <div className="flex items-center justify-center px-5 py-3 flex-col gap-6">
                <InputField 
                    setValue={setName} 
                    value={name}
                    label="Your name:" 
                    type="text" 
                    placeholder="John Doe"
                    icon={BiUser}
                    settings={nameSettings}
                    validate={validateName}
                />
                
                <InputField 
                    setValue={setEmail} 
                    value={email}
                    label="Your email address:" 
                    type="email" 
                    placeholder="john@doe.com"
                    icon={VscMail}
                    settings={emailSettings}
                    validate={validateEmail}
                />

                <InputField 
                    setValue={setPassword} 
                    value={password}
                    label="Your password:" 
                    type="password" 
                    placeholder="secret123"
                    icon={VscKey}
                    settings={passwordSettings}
                    validate={validatePassword}
                />

                <div className="text-sm text-slate-700">
                    Have an account? Log in&nbsp;
                    <span className="text-blue-700 cursor-pointer transition-all duration-300 hover:text-blue-500">
                        <Link to="/">here</Link>
                    </span>!
                </div>

                <ButtonPrimary 
                    isActive={isBtnActive} 
                    onClick={handleFormSubmit} 
                    text="Register" 
                />
            </div>
        </div>
    );
};

export default Register;