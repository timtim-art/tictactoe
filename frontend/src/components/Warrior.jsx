function Warrior({type, y, x}) {


    return (
        <div style={{position: 'absolute', left: `${x}px`, top: `${y}px`}}>
            <p>{type}</p>
        </div>
    )
}

export default Warrior