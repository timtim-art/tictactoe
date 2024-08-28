import Warrior from "./Warrior.jsx";


function Battlefield() {

    const warrior_list = [
        {
            'id': 1,
            'type': 'scissor',
            'x': 250,
            'y': 300
        },
        {
            'id': 2,
            'type': 'stone',
            'x': 200,
            'y': 120
        },
        {
            'id': 3,
            'type': 'paper',
            'x': 350,
            'y': 10
        }
    ]

    return (
        <div className="w-[500px] h-[700px] border-4 border-black relative">
            <h1 className="text-center">This is the battlefield</h1>
            {warrior_list.map(warrior => <Warrior key={warrior.id} type={warrior.type} x={warrior.x} y={warrior.y}/>)}

        </div>
    )
}

export default Battlefield