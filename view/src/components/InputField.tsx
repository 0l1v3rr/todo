import { Dispatch, FC, SetStateAction, useState } from "react";
import { IconType } from "react-icons/lib";

type Settings = {
    isError: boolean, 
    errorMessage: string
}

interface InputProps {
    type: "text" | "email" | "number" | "url" | "password",
    placeholder: string,
    label: string,
    setValue: Dispatch<SetStateAction<string>>,
    value: string,
    icon: IconType,
    settings: Settings,
    validate(value: string): Settings
}

const InputField:FC<InputProps> = (props) => {
    const [currentSettings, setCurrentSettings] = useState<Settings>(props.settings);

    return (
        <div className="flex flex-col gap-1 w-full text-base">
            <label className="text-md">{props.label}</label>

            <div className="flex gap-0">
                <div className="h-9 bg-slate-50 rounded-tl-md rounded-bl-md border border-solid 
                    border-slate-200 text-base border-r-0 p-1 flex items-center justify-center text-slate-600">
                    {<props.icon />}
                </div>
                <input 
                    onChange={(e) => {
                        props.setValue(e.target.value);
                        setCurrentSettings(props.validate(e.target.value));
                    }}
                    type={props.type} 
                    placeholder={props.placeholder}
                    value={props.value}
                    className={`px-2 py-1 outline-none h-9 bg-slate-50 rounded-tr-md rounded-br-md w-full 
                        border border-solid  text-base placeholder:text-slate-600 transition-all duration-300
                        ${!currentSettings.isError ? `focus:border-blue-500 active:border-blue-500 
                            border-slate-200` : `focus:border-red-500 active:border-red-500 
                            border-red-500`}`}
                />
                {props.type == "password" && <div></div>}
            </div>

            {currentSettings.isError && <div className="text-red-600 text-sm">{currentSettings.errorMessage}</div>}
        </div>
    )
}

export default InputField;