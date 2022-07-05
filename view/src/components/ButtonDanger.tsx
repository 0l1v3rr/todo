import { FC } from "react";

interface ButtonProps {
    text: string,
    onClick(): void,
    isActive: boolean
}

const ButtonPrimary:FC<ButtonProps> = (props) => {
    return (
        <button type="button" onClick={props.onClick} disabled={!props.isActive} className="bg-red-600 
            text-white py-1 px-3 rounded-md transition-all duration-300 border border-solid border-red-500 
            flex items-center justify-center hover:bg-red-500 w-fit disabled:bg-red-400 
            disabled:cursor-default disabled:border-red-400">
            {props.text}
        </button>
    );
};

export default ButtonPrimary;