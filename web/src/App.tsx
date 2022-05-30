import { Component } from "solid-js";

import { DisplayCount } from "./components/DisplayCount";
import { Login } from "./components/Login";

export function App() {
  return (
    <div>
      <Login />
      <DisplayCount />
    </div>
  );
}
