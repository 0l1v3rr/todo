import { FC, useState } from "react";
import { VscKey, VscMail } from "react-icons/vsc";
import { Link } from "react-router-dom";

import axios, { AxiosError } from "axios";

import ButtonPrimary from "../components/ButtonPrimary";
import InputField from "../components/InputField";

type Settings = {
    isError: boolean, 
    errorMessage: string
}

const Login:FC = () => {
    const [isBtnActive, setIsBtnActive] = useState(false);
    const [requestError, setRequestError] = useState("");

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

    const checkForActiveButton = (): void => {
        if(passwordSettings.errorMessage === "" && emailSettings.errorMessage === "") {
            if(password.trim() == "" || email.trim() == "") {
                setIsBtnActive(false);
                return;
            }

            setIsBtnActive(true);
            return;
        }

        setIsBtnActive(false);
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

        // if the value is less than 4 characters long
        if(value.length < 4) {
            setEmailSettings({
                isError: true,
                errorMessage: "This email is too short."
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

        // if the value is less than 4 characters long
        if(value.length < 3) {
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

    // handle form submit function
    const handleFormSubmit = (): void => {
        setIsBtnActive(false);
        setRequestError("");
        
        axios.post(`http://localhost:8080/api/v1/login`, {
            email: email,
            password: password
        }, {
            withCredentials: true
        })
        .then(res => {
            if(res.status === 200) {
                window.location.reload();
            }
        })
        .catch(err => {
            const error = err as AxiosError;
            setRequestError((error.response?.data as any).error);
        });
    };

    return (
        <div className="shadow-md bg-slate-100 border border-solid border-slate-200 rounded-md">
            <div className="text-center border-b border-b-slate-200 border-solid font-bold px-10 
                py-2 rounded-tl-md text-2xl">
                Log In
            </div>

            {requestError !== "" && 
                <div className="mx-5 my-3 px-3 py-1 text-center bg-red-200 text-red-900 rounded-md border border-solid border-red-300">
                    {requestError}
                </div>}

            <div className="flex items-center justify-center px-5 py-3 flex-col gap-6">
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
                    Don't have an account? Register&nbsp;
                    <span className="text-blue-700 cursor-pointer transition-all duration-300 hover:text-blue-500">
                        <Link to="/register">here</Link>
                    </span>!
                </div>

                <ButtonPrimary 
                    isActive={isBtnActive} 
                    onClick={handleFormSubmit} 
                    text="Log In" 
                />
            </div>
        </div>
    );
}

export default Login;