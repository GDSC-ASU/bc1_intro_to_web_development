import { useState } from "react";
import "./App.css";
import Student from "./classes/Student";
import Students from "./componets/Students";

function App() {
  const [arr, setStudents] = useState([
    new Student("1", "Ali", 90),
    new Student("2", "Jaber", 70),
    new Student("3", "Omar", 80.3),
  ]);

  return (
    <div>
      <Students students={arr} setStudents={setStudents} />
    </div>
  );
}

export default App;
