import { lazy } from "solid-js";

import AboutData from "./pages/About.data";
import Home from "./pages/Home";

import type { RouteDefinition } from "solid-app-router";

export const routes: RouteDefinition[] = [
  {
    path: "/",
    component: Home,
  },
  {
    path: "/about",
    component: lazy(() => import("./pages/About")),
    data: AboutData,
  },
  {
    path: "/protected",
    component: lazy(() => import("./pages/ProtectedAbout")),
  },
  {
    path: "**",
    component: lazy(() => import("./errors/404")),
  },
];
