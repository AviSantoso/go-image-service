import { useLocation } from "solid-app-router";
import { Component, Show } from "solid-js";

import logo from "../assets/logo.svg";
import { useAuth } from "../hooks/useAuth";
import { NavBarLink } from "./NavBarLink";

export const NavBar: Component = () => {
  const { loggedIn } = useAuth();
  const location = useLocation();

  return (
    <nav class="bg-hex-f64970 px-8 fixed left-0 right-0 top-0 h-5em">
      <div class="container mx-auto max-w-1024px flex flex-wrap justify-between items-center p-4">
        <div class="flex flex-row w-3/10 space-x-8 items-center justify-start text-sm md:text-md font-bold w-auto">
          <NavBarLink
            title="Home"
            href="/"
            active={location.pathname === "/"}
          />
        </div>
        <img src={logo} class="h-3em" alt="Go Image Service Logo" />
        <div class="flex flex-row  w-3/10 space-x-8 items-center justify-end text-sm md:text-md font-bold w-auto">
          <Show
            when={loggedIn()}
            fallback={
              <NavBarLink
                title="Login"
                href="/login"
                active={location.pathname === "/login"}
              />
            }
          >
            <NavBarLink
              title="Dashboard"
              href="/dashboard"
              active={location.pathname === "/dashboard"}
            />
            <NavBarLink title="Logout" href="/logout" />
          </Show>
        </div>
      </div>
    </nav>
  );
};
