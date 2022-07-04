import InputField from "./InputField";
import { BsCardChecklist } from 'react-icons/bs';
import { FC, useState } from "react";
import ButtonSecondary from "./ButtonSecondary";

type Settings = {
    isError: boolean, 
    errorMessage: string
}

interface CreateListProps {
    createList(name: string, setBtnActive: React.Dispatch<React.SetStateAction<boolean>>): void
}

const CreateList:FC<CreateListProps> = (props) => {
    const [name, setName] = useState("");

    const [isBtnActive, setIsBtnActive] = useState(false);

    const validateInput = (value: string) : Settings => {
        // trimming the value
        value = value.trim();

        // if the value is empty
        if(value == "") {
            setIsBtnActive(false);
            return {
                isError: false,
                errorMessage: ""
            }
        }

        // if the value is less than 3 characters long
        if(value.length < 3) {
            setIsBtnActive(false);
            return {
                isError: false,
                errorMessage: ""
            }
        }

        // if the value is valid
        setIsBtnActive(true);
        return {
            isError: false,
            errorMessage: ""
        }
    }

    const handleBtnClick = () => {
        setName("");
        props.createList(name, setIsBtnActive);
    }
    
    return (
        <div className="w-full pb-3 pt-1 border-t border-b border-solid border-slate-300">
            <div>
                Create list
            </div>
            <div className="flex items-stretch justify-center gap-2">
                <InputField 
                    icon={BsCardChecklist}
                    label=""
                    placeholder="List name"
                    setValue={setName}
                    value={name}
                    type="text"
                    validate={validateInput}
                    settings={{
                        isError: false,
                        errorMessage: ""
                    }}
                />
                <ButtonSecondary 
                    text="Create"
                    isActive={isBtnActive}
                    onClick={handleBtnClick}
                />
            </div>
        </div>
    );
};

export default CreateList;