import { useState } from "react";

function App() {
  const [counter, setCounter] = useState(0);

  const increaseCounter = () => {
    setCounter(counter + 1);
  };

  return (
    <div>
      <p>Your counter is {counter}</p>
      <button onClick={increaseCounter}>Increase</button>
    </div>
  );
}

export default App;
