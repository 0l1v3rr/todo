import { Link } from "react-router-dom";
import ButtonSuccess from "../components/ButtonSuccess"
import ButtonWarning from "../components/ButtonWarning"

import { MdArrowBackIos, MdDriveFileRenameOutline, MdOutlineDescription } from "react-icons/md"
import { FC, useState } from "react";
import PopupContainer from "./PopupContainer";
import BlurOverlay from "./BlurOverlay";
import InputField from "./InputField";

type Settings = {
    isError: boolean, 
    errorMessage: string
}

interface NavbarProps {
    createTask: (name: string, description: string) => Promise<void>
}

const ListNavbar:FC<NavbarProps> = ({ createTask }) => {
    const [isPopupActive, setIsPopupActive] = useState(false);
    const [isBtnActive, setIsBtnActive] = useState(false);

    const [name, setName] = useState('');
    const [nameSettings, setNameSettings] = useState<Settings>({
        isError: false, 
        errorMessage: "" 
    });

    const [desc, setDesc] = useState("");
    const [descSettings, setDescSettings] = useState<Settings>({
        isError: false, 
        errorMessage: "" 
    });

    const onTaskCreate = () => {
        createTask(name, desc);
        setIsPopupActive(false);
    };

    const checkForActiveButton = (): void => {
        if(descSettings.errorMessage === "" && nameSettings.errorMessage === "") {
            if(name.trim() == "" || desc.trim() == "") {
                setIsBtnActive(false);
                return;
            }

            setIsBtnActive(true);
            return;
        }

        setIsBtnActive(false);
    }

    // name validating
    const validateName = (value: string): Settings => {
        // trimming the value
        value = value.trim();

        // if the value is empty
        if(value == "") {
            setNameSettings({
                isError: false,
                errorMessage: ""
            });
            checkForActiveButton();
            return nameSettings;
        }

        // if the value is less than 3 characters long
        if(value.length < 3) {
            setNameSettings({
                isError: true,
                errorMessage: "This name is too short."
            });
            checkForActiveButton();
            return nameSettings;
        }

        // if the value is more than 32 characters long
        if(value.length < 3) {
            setNameSettings({
                isError: true,
                errorMessage: "This name is too long."
            });
            checkForActiveButton();
            return nameSettings;
        }
        
        // if the value is valid
        setNameSettings({
            isError: false,
            errorMessage: ""
        });
        checkForActiveButton();
        return nameSettings;
    };

    // description validating
    const validateDescription = (value: string): Settings => {
        // trimming the value
        value = value.trim();

        // if the value is empty
        if(value == "") {
            setDescSettings({
                isError: false,
                errorMessage: ""
            });
            checkForActiveButton();
            return descSettings;
        }

        // if the value is more than 256 characters long
        if(value.length > 256) {
            setDescSettings({
                isError: true,
                errorMessage: "The description is too long."
            });
            checkForActiveButton();
            return descSettings;
        }
        
        // if the value is valid
        setDescSettings({
            isError: false,
            errorMessage: ""
        });
        checkForActiveButton();
        return descSettings;
    };
    
    return (
        <div className="flex items-center justify-between w-full gap-14 shadow-md bg-slate-100 
            border border-solid border-slate-200 rounded-md py-2 px-4">
            <PopupContainer isActive={isPopupActive}>
                <>
                    <div className="font-bold text-center border-b border-solid border-slate-300 py-1">
                        Create task
                    </div>
                    <div className="flex items-center justify-center px-5 py-3 flex-col gap-6">
                        <InputField 
                            setValue={setName} 
                            value={name}
                            label="Task name:" 
                            type="text" 
                            placeholder="Feed the dog"
                            icon={MdDriveFileRenameOutline}
                            settings={nameSettings}
                            validate={validateName}
                        />

                        <InputField 
                            setValue={setDesc} 
                            value={desc}
                            label="Task description:" 
                            type="text" 
                            placeholder="Feed the dog with fresh dog food. :)"
                            icon={MdOutlineDescription}
                            settings={descSettings}
                            validate={validateDescription}
                        />
                    </div>
                    <div className="border-t border-solid border-slate-300 p-2 flex items-center 
                        text-sm justify-between">
                        <ButtonWarning
                            isActive={true}
                            onClick={() => {
                                setIsPopupActive(false);
                                setName("");
                                setDesc("");
                            }}
                            text="Cancel"
                        />

                        <ButtonSuccess 
                            isActive={isBtnActive}
                            onClick={onTaskCreate}
                            text="Create"
                        />
                    </div>
                </>
            </PopupContainer>
            <BlurOverlay isActive={isPopupActive} />
            
            <Link to="/">
                <div className="flex gap-1 items-center cursor-pointer transition-all 
                    duration-300 hover:text-blue-600">
                    <MdArrowBackIos />
                </div>
            </Link>

            <div className="text-sm">
                <ButtonSuccess
                    isActive={true}
                    onClick={() => setIsPopupActive(true)}
                    text="Create New Task"
                />
            </div>
        </div>
    );
}

export default ListNavbar;