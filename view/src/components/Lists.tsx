import axios from "axios";
import { FC, useEffect, useState } from "react";
import Loading from "../pages/Loading";
import { List } from "../types";
import CreateList from "./CreateList";
import ListItem from "./ListItem";

interface ListsProps {
    userId: number | undefined,
}

interface ListListProps {
    lists: any[]
}

const Lists:FC<ListsProps> = (props) => {
    const [isLoaded, setIsLoaded] = useState(false);
    const [lists, setLists] = useState<any[]>([]);

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

    const createList = async (name: string, setBtnActive: React.Dispatch<React.SetStateAction<boolean>>) => {
        setBtnActive(false);

        await axios.post(`${process.env.REACT_APP_BACKEND_DOMAIN}/api/v1/lists`, {
            name: name
        }, { withCredentials: true })
            .then((res) => setLists([res.data, ...lists]))
            .catch();

        setBtnActive(true);
    }

    return (
        <div className="flex flex-col items-center justify-between w-full shadow-md bg-slate-100 
            border border-solid border-slate-200 rounded-md py-2 px-4">
            <div className="font-bold flex items-center justify-center w-full mb-3">Your Lists</div>

            <CreateList createList={createList} />

            <div className="flex flex-col gap-2 w-full items-center justify-center mt-3">
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