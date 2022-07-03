import { FC } from "react";
import Header from "../components/Header";

interface User {
    id: number,
    email: string,
    name: string,
    isEnabled: boolean
}

interface HomeProps {
    user: User | null,
}

const Home:FC<HomeProps> = (props) => {
    return (
        <div className="w-fit">
            <Header username={props.user?.name} />
            Home content
        </div>
    );
}

export default Home;