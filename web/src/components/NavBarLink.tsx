import { Link } from "solid-app-router";
import { Component, Show } from "solid-js";

interface NavBarLinkProps {
  title: string;
  href: string;
  active?: boolean;
}

export const NavBarLink: Component<NavBarLinkProps> = (props) => {
  return (
    <Show
      when={props.active}
      fallback={
        <Link href={props.href} style={{ color: "#fff8f2" }}>
          {props.title}
        </Link>
      }
    >
      <Link href={props.href} style={{ color: "#f7ddc7" }}>
        {props.title}
      </Link>
    </Show>
  );
};
