import { Link } from "react-router-dom";
import ButtonSuccess from "../components/ButtonSuccess"

import { MdArrowBackIos } from "react-icons/md"

const ListNavbar = () => {
    const onTaskCreate = () => {

    };
    
    return (
        <div className="flex items-center justify-between w-full gap-14 shadow-md bg-slate-100 
            border border-solid border-slate-200 rounded-md py-2 px-4">
            <Link to="/">
                <div className="flex gap-1 items-center cursor-pointer transition-all 
                    duration-300 hover:text-blue-600">
                    <MdArrowBackIos />
                </div>
            </Link>

            <div className="text-sm">
                <ButtonSuccess
                    isActive={true}
                    onClick={onTaskCreate}
                    text="Create New Task"
                />
            </div>
        </div>
    );
}

export default ListNavbar;