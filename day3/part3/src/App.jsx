import { useState } from "react";

function App() {
  const [list, setList] = useState([]);
  const [todo, setTodo] = useState("");

  const addTodo = () => {
    list.push(todo);
    setTodo("");
  };

  const onChangeHandler = (e) => {
    setTodo(e.target.value);
  };

  return (
    <div>
      <label for={"todoItem"}>New Todo Item</label>
      <input id={"todoItem"} value={todo} onChange={onChangeHandler} />
      <button onClick={addTodo}>Add Todo</button>
      <br />
      <ul>
        {list.map((todo) => {
          return <li>{todo}</li>;
        })}
      </ul>
    </div>
  );
}

export default App;
