import { FC } from "react";

interface ButtonProps {
    text: string,
    onClick(): void,
    isActive: boolean
}

const ButtonSecondary:FC<ButtonProps> = (props) => {
    return (
        <button type="button" onClick={props.onClick} disabled={!props.isActive} className="text-blue-600 
            py-1 px-3 rounded-md transition-all duration-300 border border-solid border-blue-600 
            flex items-center justify-center hover:bg-blue-600 hover:text-white disabled:hover:text-blue-300 
            disabled:hover:bg-transparent w-fit disabled:text-blue-300 disabled:cursor-default 
            disabled:border-blue-300">
            {props.text}
        </button>
    );
};

export default ButtonSecondary;