import { FC } from "react";
import { Link } from "react-router-dom";
import { List } from "../types";

interface ListItemProps {
    list: List
}

const ListItem:FC<ListItemProps> = (props) => {
    return (
        <div className="w-full border border-solid border-slate-200 rounded-md shadow-sm">
            {/*<div className="w-full h-4 bg-gradient-to-r from-cyan-500 to-blue-500 rounded-tr-md 
                rounded-tl-md" />*/}

            <div className="flex items-center justify-between px-4 py-2">
                <Link to={`/lists/${props.list.url}`}>
                    <span className="transition-all duration-300 hover:underline">
                        {props.list.name}
                    </span>
                </Link>
                <div>Owner: <strong>You</strong></div>
            </div>
        </div>
    );
};

export default ListItem;