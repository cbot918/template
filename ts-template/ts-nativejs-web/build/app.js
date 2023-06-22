import Hello from "./components/Hello.js";
window.onload = () => {
    initApp("app", Hello());
};
const initApp = (root, message) => {
    let app = document.getElementById(root);
    app.innerHTML = message;
};
