import { FC, useState } from "react";
import { Task } from "../types";

import ButtonDanger from "../components/ButtonDanger"
import ButtonWarning from "../components/ButtonWarning"
import axios from "axios";
import PopupContainer from "./PopupContainer";
import BlurOverlay from "./BlurOverlay";
import ButtonSuccess from "./ButtonSuccess";

interface TaskItemProps {
    task: Task,
}

const TaskItem:FC<TaskItemProps> = (props) => {
    const [task, setTask] = useState(props.task);
    const [isPopupActive, setIsPopupActive] = useState(false);
    
    const handleDeleteClick = () => {
        axios.delete(`${process.env.REACT_APP_BACKEND_DOMAIN}/api/v1/tasks/${task.id}`, { withCredentials: true });
        window.location.reload();
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
            <PopupContainer isActive={isPopupActive}>
                <>
                    <div className="font-bold text-center border-b border-solid border-slate-300 py-1">
                        {task.title}
                    </div>
                    <div className="text-slate-700 py-1 px-2">
                        Are you sure you want to delete this task?
                    </div>
                    <div className="border-t border-solid border-slate-300 p-2 flex items-center 
                        text-sm justify-between">
                        <ButtonSuccess
                            isActive={true}
                            onClick={() => setIsPopupActive(false)}
                            text="Cancel"
                        />

                        <ButtonDanger 
                            isActive={true}
                            onClick={handleDeleteClick}
                            text="I'm sure"
                        />
                    </div>
                </>
            </PopupContainer>
            <BlurOverlay isActive={isPopupActive} />
            
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
                    onClick={() => setIsPopupActive(true)}
                />
            </div>
        </div>
    );
}

export default TaskItem;