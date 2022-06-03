import { useRoutes } from "solid-app-router";

import { NavBar } from "./components/NavBar";
import { routes } from "./routes";

import type { Component } from "solid-js";

export const App: Component = () => {
  const Route = useRoutes(routes);

  return (
    <>
      <NavBar />
      <main class="mt-6em">
        <Route />
      </main>
    </>
  );
};

export default App;
