import { useEffect, useState } from "react";

const API_URL = "http://localhost:8080";
const IS_PROD = import.meta.env.MODE === "production";

interface Todo {
  id: number;
  id_done: boolean;
  text: string;
}

async function fetchTodos() {
  const url = IS_PROD ? "/todos" : `${API_URL}/todos`;
  const response = await fetch(url);
  const { todos } = (await response.json()) as { todos: Todo[] };
  return todos;
}

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const imageUrl = IS_PROD ? "/image" : `${API_URL}/image`;

  useEffect(() => {
    fetchTodos().then((todos) => setTodos(todos));
  }, []);

  return (
    <>
      <img src={imageUrl} height={400} width={400} alt="Random image" />
      <div>
        <input type="text" maxLength={140} />
        <button>Create</button>
      </div>
      <ul>
        {todos.map((todo) => (
          <li key={todo.id}>{todo.text}</li>
        ))}
      </ul>
    </>
  );
}

export default App;
