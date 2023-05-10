import { useState, useContext } from "react";
import { useNavigate } from "react-router-dom";
import { UserContext } from "../App";

export function Login() {
  const { state, dispatch } = useContext(UserContext);
  let navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  function Login(e) {
    e.preventDefault();

    const URL = "http://localhost:3500/auth";
    fetch(URL, {
      method: "POST",
      headers: { "Content-type": "application/json" },
      body: JSON.stringify({
        id: 1,
        email: email,
        password: password,
      }),
    })
      .then((res) => res.json())
      .then((data) => {
        if (data.error) {
          console.log(`錯誤發生在apiCall.js.fetch(): ${data.error}`);
          return;
        }

        localStorage.setItem("jwt", data.token);
        localStorage.setItem("user", JSON.stringify(data.user));

        dispatch({ type: "USER", payload: data.user });

        console.log("使用者登入成功");
        navigate("/");
      })
      .catch((err) => console.log(err));
  }

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
      }}
    >
      <input
        type="text"
        placeholder="email"
        onChange={(e) => setEmail(e.target.value)}
      />
      <input
        type="text"
        placeholder="password"
        onChange={(e) => setPassword(e.target.value)}
      />
      <button onClick={(e) => Login(e)}>Login</button>
    </div>
  );
}
