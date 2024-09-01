import Warrior from "./Warrior.jsx";
import io from 'socket.io-client'
import {useEffect, useState} from "react";


function Battlefield() {

    const [warrior_list, setWarriorList] = useState([])


    // let warrior_list = [
    //     {
    //         id: 1,
    //         type: 'scissor',
    //         x: 250,
    //         y: 300
    //     },
    //     {
    //         id: 2,
    //         type: 'stone',
    //         x: 200,
    //         y: 120
    //     },
    //     {
    //         id: 3,
    //         type: 'paper',
    //         x: 350,
    //         y: 10
    //     }
    // ];

    //const socket = io('localhost:8080/api/v1/stream/join?gameId=1')


    let conn
    function connectSocket() {
        
        if (!window.WebSocket) {
            console.log("Your browser does not suppot Websockets")
            return
        }
        
        conn = new WebSocket(`ws://localhost:8080/api/v1/stream/join?gameId=1`)
        conn.onclose = () => {
            console.log("Websocket connection closed")
        }
        conn.onmessage = (event) => {
            let message = event.data
            console.log(message)
            setWarriorList(JSON.parse(message))
        }

        /* socket.on('connection', (warrior_list) => {
            console.log(warrior_list)
            setWarriorList(warrior_list)
        }) */
    }

    useEffect(() => {
        connectSocket()
    }, []);

    return (
        <div className="w-[500px] h-[700px] border-4 border-black relative">
            <h1 className="text-center">This is the battlefield</h1>
            {warrior_list.map(warrior => <Warrior key={warrior.id} type={warrior.type} x={warrior.x} y={warrior.y}/>)}
        </div>
    )
}

export default Battlefield