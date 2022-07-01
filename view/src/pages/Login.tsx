import { useState } from "react";
import { MdOutlineEmail } from "react-icons/md";

import ButtonPrimary from "../components/ButtonPrimary";
import InputField from "../components/InputField";

type Settings = {
    isError: boolean, 
    errorMessage: string
}

const Login = () => {
    const [isBtnActive, setIsBtnActive] = useState(false);

    // email specific things
    const [email, setEmail] = useState('');
    const [emailSettings, setEmailSettings] = useState<Settings>({
        isError: false, 
        errorMessage: "" 
    });
    const validateEmail = (value: string): Settings => {
        // trimming the value
        value = value.trim();

        // if the value is empty
        if(value == "") {
            setIsBtnActive(false);
            return {
                isError: false,
                errorMessage: ""
            };
        }

        // if the value is less than 4 characters long
        if(value.length < 4) {
            setIsBtnActive(false);
            return {
                isError: true,
                errorMessage: "This email is too short."
            };
        }
        
        // if the value is valid
        setIsBtnActive(true);
        return {
            isError: false,
            errorMessage: ""
        };
    };

    // handle form submit function
    const handleFormSubmit = (): void => {

    };

    return (
        <div className="shadow-md bg-slate-100">
            <div className="text-center border-b border-b-slate-200 border-solid font-bold px-10 
                py-2 rounded-tl-md text-lg">
                Log In
            </div>

            <div className="flex items-center justify-center px-5 py-3 flex-col gap-2">
                <InputField 
                    setValue={setEmail} 
                    value={email}
                    label="Your email address:" 
                    type="email" 
                    placeholder="john@doe.com"
                    icon={MdOutlineEmail}
                    settings={emailSettings}
                    validate={validateEmail}
                />

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