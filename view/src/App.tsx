import Home from "./pages/Home";

import { useEffect, useState } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from "./pages/Login";
import Register from "./pages/Register";

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
    setIsLoaded(true);
  });

  return (
    <div className="w-screen min-h-screen flex items-center justify-center bg-slate-50 text-slate-800">
        <Router>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/login" element={<Login />} />
                <Route path="/register" element={<Register />} />
            </Routes>
        </Router>
    </div>
  );
}

export default App;
