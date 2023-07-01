const log = console.log
const ws = {}
let connBtn = {}

window.onload = ()=>{
  const app = document.querySelector("#app")
  renderTitle(app, "Qchat")
  insertBlank(app)
  connBtn = renderConnBtn(app)
  log("in")
  initWs()
  
}

// connBtn.addEventListener("click",initWs)


function renderTitle(root, title){
  root.innerHTML = title
}

function renderConnBtn(root){
  log("in")
  const connBtn = document.createElement('input')
  connBtn.setAttribute("type","button")
  connBtn.setAttribute("value","連線")
  root.appendChild(connBtn)
}

function insertBlank(root){
  const blank = document.createElement('br')
  root.appendChild(blank)
}

function initWs(){
  log("hihi")
  const socket = new WebSocket("ws://localhost:8787")
  socket.onopen = function () {
    log("open socket")
    ws = socket
    log(ws)
  }
  socket.onerror = function(err){
    log(err)
  }
}