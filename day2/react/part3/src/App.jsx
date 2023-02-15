import { useState } from "react";

function App() {
  const [list, setList] = useState([]);
  const [todo, setTodo] = useState("");

  const addTodo = () => {
    list.push(todo);
    setList(list.flat());
  };

  return (
    <div>
      <input
        className=""
        value={todo}
        onChange={(e) => {
          setTodo(e.target.value);
        }}
      />
      <button onClick={addTodo}>Add Todo</button>
      <br />
      <ul>
        {list.map((todo, index) => {
          return <li key={index}>{todo}</li>;
        })}
      </ul>
    </div>
  );
}

export default App;
