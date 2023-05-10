import { createContext, useReducer, useEffect } from "react";
import { Routes, Route, useNavigate } from "react-router-dom";

import { reducer, initialState } from "./reducer/useReducer";

import { Login } from "./components/Login";
import { Home } from "./components/Home";
import "./App.css";

export const UserContext = createContext();

function App() {
  const [state, dispatch] = useReducer(reducer, initialState);
  let navigate = useNavigate();

  useEffect(() => {
    const user = JSON.parse(localStorage.getItem("user"));

    if (!user) {
      navigate("/login");
    }

    dispatch({ type: "USER", payload: user });
  }, []);

  return (
    <UserContext.Provider value={{ state, dispatch }}>
      <Routes>
        <Route path="/" element={<Home />}></Route>
        <Route path="/login" element={<Login />}>
          {console.log(state)}
        </Route>
      </Routes>
    </UserContext.Provider>
  );
}

export default App;
