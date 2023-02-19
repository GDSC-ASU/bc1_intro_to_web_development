function Stats({ xScore, oScore, turn }) {
    return <div>
        <p>Player X's Score: {xScore}</p>
        <p>Player O's Score: {oScore}</p>
        <p>Current Player: {turn}</p>
    </div>
}

export default Stats;