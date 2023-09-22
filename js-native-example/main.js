
window.onload = () =>{
  // 把 html element 先一次宣告好
  let body = document.querySelector("body")
  let inputText = document.createElement("input")
  let view = document.createElement("div")

  // setup inputText
  body.appendChild(inputText) // 將 inputText 掛到 body 下面
  inputText.setAttribute("type","text") // 為了輸入文字, 設定為 text 類型
  inputText.setAttribute("placeholder","input text") // 給個預設文字友善使用者
  inputText.addEventListener("keydown",(e)=>{   // 註冊事件監聽器, keydown 為鍵盤壓下的時候
    if( e.key === "Enter"){                     // 捕捉 Enter 按鍵
      let temp = document.createElement("div")  // 每一個輸入的文字 用個 div 去渲染
      temp.innerHTML = e.target.value           // 把 div 加上文字, e.target.value 就是使用者輸入
      view.appendChild( temp )                  // 將 temp 加入 view
      e.target.value = ""                       // 將 輸入框的文字清掉, 讓使用者輸入下一個
    }
  })
  inputText.setAttribute("style","margin-bottom:10px")  // 這邊做個 css 設定

  // setup view
  body.appendChild(view)  // 將 view 加入 body下面

}