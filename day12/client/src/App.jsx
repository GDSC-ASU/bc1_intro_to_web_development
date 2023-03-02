import "./App.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Game from "./pages/Game";
import Login from "./pages/Login";
import Signup from "./pages/Signup";

function App() {
  const router = createBrowserRouter([
    { path: "/", element: <Game /> },
    { path: "/login", element: <Login /> },
    { path: "/signup", element: <Signup /> },
  ]);

  return (
    <div>
      <RouterProvider router={router} />
    </div>
  );
}

export default App;
