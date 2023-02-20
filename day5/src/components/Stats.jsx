import "./Stats.css";

function Stats({ xScore, oScore, turn }) {
  return (
    <div className="container">
      <div>
        <span>Player X</span>
        <br />
        <span>{xScore}</span>
      </div>

      <div>
        <span>Player O</span>
        <br />
        <span>{oScore}</span>
      </div>

      <div>
        <span>Turn</span>
        <br />
        <span>{turn}</span>
      </div>
    </div>
  );
}

export default Stats;
