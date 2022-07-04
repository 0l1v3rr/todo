import { FC } from "react";
import Header from "../components/Header";
import Lists from "../components/Lists";

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
        <div className="w-fit flex flex-col gap-2">
            <Header username={props.user?.name} />
            <Lists userId={props.user?.id} />
        </div>
    );
}

export default Home;