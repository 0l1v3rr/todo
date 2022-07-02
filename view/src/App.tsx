import Home from "./pages/Home";

import { useEffect, useState } from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Login from "./pages/Login";
import Register from "./pages/Register";
import axios from "axios";
import Loading from "./pages/Loading";

interface User {
    id: number,
    email: string,
    name: string,
    isEnabled: boolean
}

const App = () => {
  const [isLoaded, setIsLoaded] = useState(false);
  const [loggedInUser, setLoggedInUser] = useState<User | null>(null);

  useEffect(() => {
    (async () => {
        await axios.get(`http://localhost:8080/api/v1/user`, { withCredentials: true })
            .then(res => setLoggedInUser(res.data))
            .catch(() => {});

        setIsLoaded(true);
    })();
  }, []);

  return (
    <div className="w-screen min-h-screen flex items-center justify-center bg-slate-50 text-slate-800">
        <Router>
            <Routes>
                <Route path="/" element={
                    isLoaded ? (loggedInUser == null ? <Login /> : <Home />) : <Loading />
                } />

                <Route path="/register" element={
                    isLoaded ? <Register user={loggedInUser} /> : <Loading />
                } />
            </Routes>
        </Router>
    </div>
  );
}

export default App;
