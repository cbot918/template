window.onload = () => {
    initApp("app", "hello ");
};
const initApp = (root, message) => {
    let app = document.getElementById(root);
    app.innerHTML = message;
};
