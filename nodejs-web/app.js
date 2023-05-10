const express = require('express')
const app = express()
const port = 3000

app.use(express.static(__dirname + '/ui/dist'));

app.get('/', (req, res) => {
  console.log("request in")
  res.sendFile(__dirname + '/ui/dist/index.html')
})

app.listen(port, () => {
  console.log(`listening on port ${port}`)
})