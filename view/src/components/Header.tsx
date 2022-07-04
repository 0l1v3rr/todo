import axios from "axios";
import { FC } from "react";
import ButtonPrimary from "./ButtonPrimary";

interface HeaderProps {
    username: string | undefined,
}

const Header:FC<HeaderProps> = (props) => {
    const handleLogoutClick = () => {
        axios.post(`${process.env.REACT_APP_BACKEND_DOMAIN}/api/v1/logout`, null, { withCredentials: true });
        window.location.reload();
    };
    
    return (
        <div className="flex items-center justify-between w-fit gap-14 shadow-md bg-slate-100 
            border border-solid border-slate-200 rounded-md py-2 px-4">
            <div className="whitespace-nowrap">Welcome, <strong>{props.username}</strong>! 👋</div>
            <ButtonPrimary text="Log Out" isActive={true} onClick={handleLogoutClick} />
        </div>
    );
};

export default Header;