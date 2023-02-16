import { useState } from "react";
import YouShallPass from "./components/YouShallPass";

function App() {
  const [pass, setPass] = useState(false);

  const togglePass = () => {
    setPass(!pass);
  };

  return (
    <div>
      <button onClick={togglePass}>Change</button>
      <br />
      {pass && <YouShallPass />}
    </div>
  );
}

export default App;
