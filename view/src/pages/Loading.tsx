import { FaSpinner } from "react-icons/fa";

const Loading = () => {
    return (
        <div className="flex flex-col items-center justify-center">
            <span className="animate-spin"><FaSpinner /></span>
            Loading...
        </div>
    );
}

export default Loading;