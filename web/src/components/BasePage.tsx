import { Component, JSX } from "solid-js";

interface BasePageProps {
  children: JSX.Element;
  fullWidth?: boolean;
}

export const BasePage: Component<BasePageProps> = (props) => {
  const classes = props.fullWidth ? "" : "mx-auto max-w-1024px";
  return (
    <section class={`container p-4 py-8 text-gray-700 ${classes}`}>
      {props.children}
    </section>
  );
};
