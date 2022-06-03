import { Component, JSX } from "solid-js";

interface BasePageProps {
  children: JSX.Element;
}

export const BasePage: Component<BasePageProps> = (props) => {
  return (
    <section class="container mx-auto max-w-1024px p-4 py-8 text-gray-700">
      {props.children}
    </section>
  );
};
