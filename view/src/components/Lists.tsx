import axios from "axios";
import { FC, useEffect, useState } from "react";
import Loading from "../pages/Loading";
import { List } from "../types";
import ListItem from "./ListItem";

interface ListsProps {
    userId: number | undefined,
}

interface ListListProps {
    lists: any[]
}

const Lists:FC<ListsProps> = (props) => {
    const [isLoaded, setIsLoaded] = useState(false);
    const [lists, setLists] = useState([]);

    useEffect(() => {
        (async () => {
            await axios.get(`${process.env.REACT_APP_BACKEND_DOMAIN}/api/v1/lists/user/${props.userId}`, {
                withCredentials: true
            })
            .then((res) => setLists(res.data))
            .catch(() => setLists([]));

            setIsLoaded(true);
        })();
    }, []);

    return (
        <div className="flex flex-col items-center justify-between w-full shadow-md bg-slate-100 
            border border-solid border-slate-200 rounded-md py-2 px-4">
            <div className="font-bold flex items-center justify-center w-full mb-2">Your Lists</div>

            <div className="flex flex-col gap-2 w-full items-center justify-center">
                {isLoaded ? <ListList lists={lists} /> : <Loading />}
            </div>
        </div>
    );
}

const ListList:FC<ListListProps> = (props) => {
    return (
        <>
            {props.lists.length != 0 ? 
                props.lists.map((list: List) => <ListItem key={list.id} list={list} />) :
                <>
                    You don't have a list yet. <br />
                    Consider creating one. 😉
                </>}
        </>
    );
};

export default Lists;