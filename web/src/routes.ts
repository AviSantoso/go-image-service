import { lazy } from "solid-js";

import Home from "./pages/Home";

import type { RouteDefinition } from "solid-app-router";

export const routes: RouteDefinition[] = [
  {
    path: "/",
    component: Home,
  },
  {
    path: "/dashboard",
    component: lazy(() => import("./pages/Dashboard")),
  },
  {
    path: "**",
    component: lazy(() => import("./errors/404")),
  },
];

export default routes;
