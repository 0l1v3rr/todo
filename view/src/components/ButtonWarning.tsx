import { FC } from "react";

interface ButtonProps {
    text: string,
    onClick(): void,
    isActive: boolean
}

const ButtonPrimary:FC<ButtonProps> = (props) => {
    return (
        <button type="button" onClick={props.onClick} disabled={!props.isActive} className="bg-yellow-500 
            text-white py-1 px-3 rounded-md transition-all duration-300 border border-solid border-yellow-400 
            flex items-center justify-center hover:bg-yellow-400 w-fit disabled:bg-yellow-300 
            disabled:cursor-default disabled:border-yellow-300">
            {props.text}
        </button>
    );
};

export default ButtonPrimary;