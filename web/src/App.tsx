import { Component } from "solid-js";

import { DisplayCount } from "./components/DisplayCount";
import { Login } from "./components/Login";
import { AuthProvider } from "./providers/AuthContext";

export const App: Component = () => {
  return (
    <div>
      <AuthProvider>
        <Login />
        <DisplayCount />
      </AuthProvider>
    </div>
  );
};
