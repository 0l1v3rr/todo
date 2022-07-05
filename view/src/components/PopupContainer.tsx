import { FC, ReactElement } from "react";

interface PopupProps {
    children: ReactElement,
    isActive: boolean
}

const PopupContainer:FC<PopupProps> = (props) => {
    const activeClasses = "opacity-100 top-1/2 pointer-events-auto";
    const inactiveClasses = "-top-24 opacity-0 pointer-events-none";

    return (
        <div className={`border bg-slate-50 border-solid border-slate-300 rounded-md shadow-md 
            w-fit absolute left-1/2 z-10 -translate-x-1/2 -translate-y-1/2 select-none transition-all  
            duration-300 ${props.isActive ? activeClasses : inactiveClasses}`}>
            {props.children}
        </div>
    );
};

export default PopupContainer;