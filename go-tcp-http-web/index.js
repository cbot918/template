const log = console.log
window.ws = {}
let connBtn = {}

window.onload = ()=>{
  const app = document.querySelector("#app")
  rTitle(app, "Qchat")
  rBlank(app)
  connBtn = rConnBtn(app, (t)=>{t.addEventListener("click",initWs)})
  rBlank(app)
  inputText = rInput(app, (t)=>{t.addEventListener("keydown",sendMessage)})
}

function sendMessage(e){
  if(e.key=="Enter"){
    log(ws)
    ws.send(e.target.value)
    e.target.value = ""
  }
}
function rTitle(root, title){
  root.innerHTML = title
}
function rConnBtn(root, fn){
  const connBtn = document.createElement('input')
  connBtn.setAttribute("type","button")
  connBtn.setAttribute("value","連線")
  root.appendChild(connBtn)
  fn(connBtn)
  return connBtn
}
function rBlank(root){
  const blank = document.createElement('br')
  root.appendChild(blank)
}
function rInput(root, fn){
  const inputText = document.createElement("input")
  inputText.setAttribute("type","text")
  inputText.setAttribute("placeholder","message")
  root.appendChild(inputText)
  fn(inputText)
  return inputText
}
function initWs(e){
  log("trigger")
  // e.preventDefault()
  const socket = new WebSocket("ws://localhost:8787")
  socket.onopen = function () {
    log("open socket")
    ws = socket
    log(ws)
    socket.send("hihihi")
  }
  socket.onmessage = function(m){ 
    log(m.data)
  }
  socket.onerror = function(err){
    log(err)
  }
}