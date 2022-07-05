import { FC, useState } from "react";
import { Task } from "../types";

import ButtonDanger from "../components/ButtonDanger"
import ButtonWarning from "../components/ButtonWarning"
import axios from "axios";

interface TaskItemProps {
    task: Task
}

const TaskItem:FC<TaskItemProps> = (props) => {
    const [task, setTask] = useState(props.task);
    
    const handleDeleteClick = () => {

    }

    const handleChangeStatus = () => {
        setTask(current => {
            return {...current, isDone: !current.isDone}
        });

        axios.patch(
            `${process.env.REACT_APP_BACKEND_DOMAIN}/api/v1/tasks/${task.id}`, 
            null, 
            { withCredentials: true }
        );
    }
    
    return (
        <div className="w-full border border-solid border-slate-200 rounded-md shadow-sm py-1 px-2">
            <div>
                <strong>{task.title}</strong> - {task.isDone ? 'done' : 'in progress'}
            </div>
            <div className="text-slate-600 border-b border-solid border-slate-300 pb-1 mb-1">
                {task.description}
            </div>

            <div className="flex gap-2 mt-2 items-center justify-between text-sm">
                <ButtonWarning 
                    isActive={true}
                    text={task.isDone ? "Mark as IN PROGRESS" : "Mark as DONE"}
                    onClick={handleChangeStatus}
                />

                <ButtonDanger 
                    isActive={true}
                    text="Delete"
                    onClick={handleDeleteClick}
                />
            </div>
        </div>
    );
}

export default TaskItem;