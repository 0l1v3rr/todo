import { FC } from "react";

interface ButtonProps {
    text: string,
    onClick(): void,
    isActive: boolean
}

const ButtonPrimary:FC<ButtonProps> = (props) => {
    return (
        <button type="button" onClick={props.onClick} disabled={!props.isActive} className="bg-green-600 
            text-white py-1 px-3 rounded-md transition-all duration-300 border border-solid border-green-500 
            flex items-center justify-center hover:bg-green-500 w-fit disabled:bg-green-400 
            disabled:cursor-default disabled:border-green-400">
            {props.text}
        </button>
    );
};

export default ButtonPrimary;