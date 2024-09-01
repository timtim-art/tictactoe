const express = require("express")
const cors = require("cors")
const { createServer } = require("node:http")
const { Server } = require("socket.io")

let warrior_list = [
  {
    id: 1,
    type: "scissor",
    x: 250,
    y: 300,
  },
  {
    id: 2,
    type: "stone",
    x: 200,
    y: 120,
  },
  {
    id: 3,
    type: "paper",
    x: 350,
    y: 10,
  },
]

const app = express()
app.use(
  cors({
    origin: "http://localhost:5173", // Update with your front-end origin
  })
)
const server = createServer(app)
const io = new Server(server, {
  cors: {
    origin: "http://localhost:5173", // Update with your front-end origin
  },
})

io.on("connection", (socket) => {
  socket.emit("connection", warrior_list)

  setInterval(() => {
    socket.emit("connection", warrior_list)
  }, 5000)
})

server.listen(3000, () => {
  console.log("server running at http://localhost:3000")
})
