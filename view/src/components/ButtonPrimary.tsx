import { FC } from "react";

interface ButtonProps {
    text: string,
    onClick(): void,
    isActive: boolean
}

const ButtonPrimary:FC<ButtonProps> = (props) => {
    return (
        <button type="button" onClick={props.onClick} disabled={!props.isActive} className="bg-blue-600 
            text-white py-1 px-3 rounded-md transition-all duration-300 border border-solid border-blue-500 
            flex items-center justify-center hover:bg-blue-500 w-full disabled:bg-blue-400 
            disabled:cursor-default disabled:border-blue-400">
            {props.text}
        </button>
    );
};

export default ButtonPrimary;