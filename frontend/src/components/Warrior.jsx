function Warrior({type, y, x}) {


    return (
        <div style={{position: 'absolute', left: `${x}px`, top: `${y}px`}}>
            {
                type === 'scissor' ? (
                    <p>✂️</p>
                ) : type === 'stone' ? (
                    <p>🪨</p>
                ) : (
                    <p>📄</p>
                )
            }
        </div>
    )
}

export default Warrior