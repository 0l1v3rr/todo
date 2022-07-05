import axios from "axios";
import { FC, useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import Header from "../components/Header";
import ListNavbar from "../components/ListNavbar";
import TaskItem from "../components/TaskItem";
import { List, Task } from "../types";

interface ListPageProps {
    user: User | null,
}

interface User {
    id: number,
    email: string,
    name: string,
    isEnabled: boolean
}

const ListPage:FC<ListPageProps> = (props) => {
    const [currentList, setCurrentList] = useState<List | null>(null);
    const [tasks, setTasks] = useState<any[]>([]);
    const [isLoaded, setIsLoaded] = useState(false);
    const { listUrl } = useParams();

    let navigate = useNavigate();

    useEffect(() => {
        (async () => {
            await axios.get(
                `${process.env.REACT_APP_BACKEND_DOMAIN}/api/v1/lists/${listUrl}`, 
                { withCredentials: true }
            ).then((res) => setCurrentList(res.data)).catch(() => navigate("/"));

            setIsLoaded(true);
        })();
    }, []);

    useEffect(() => {
        if(currentList?.id != undefined) {
            axios.get(
                `${process.env.REACT_APP_BACKEND_DOMAIN}/api/v1/tasks/list/${currentList?.id}`,
                { withCredentials: true }
            ).then((res) => setTasks(res.data)).catch(() => setTasks([]));
        }
    }, [currentList]);
    
    return (
        <div className="w-fit flex flex-col gap-2">
            <Header username={props.user?.name} />

            <ListNavbar />

            <div className="flex flex-col items-center justify-between shadow-md bg-slate-100 
                border border-solid border-slate-200 rounded-md py-2 px-4 w-full">
                <div className="font-bold flex items-center justify-center w-full mb-2">
                    {isLoaded ? `${currentList?.name}: tasks` : 'Loading...'}
                </div>

                <div className="flex flex-col gap-2 w-full items-center justify-center">
                    {tasks.length != 0 ? 
                        tasks.map((task: Task) => <TaskItem key={task.id} task={task} />) :
                        <>
                            You don't have a task yet. <br />
                            Consider creating one. 😉
                        </>}
                </div>
            </div>
        </div>
    );
};

export default ListPage;