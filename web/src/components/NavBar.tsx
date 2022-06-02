import { useLocation } from "solid-app-router";
import { Component, Show } from "solid-js";

import logo from "../assets/logo.svg";
import { useAuth } from "../hooks/useAuth";
import { NavBarLink } from "./NavBarLink";

export const NavBar: Component = () => {
  const { loggedIn } = useAuth();
  const location = useLocation();

  return (
    <nav class="bg-hex-f64970 px-8">
      <div class="container flex flex-wrap justify-between items-center mx-auto w-full py-3">
        <div class="flex flex-row space-x-8 items-center text-sm md:text-md font-bold w-auto">
          <NavBarLink
            title="Home"
            href="/"
            active={location.pathname === "/"}
          />
          <NavBarLink
            title="About"
            href="/about"
            active={location.pathname === "/about"}
          />
        </div>
        <img src={logo} class="h-8 md:h-16" alt="Go Image Service Logo" />
        <div class="flex flex-row space-x-8 items-center text-sm md:text-md font-bold w-auto">
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
