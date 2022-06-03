import { ComponentWithChildren } from "../types/ComponentWithChildren";

export const Title: ComponentWithChildren<{}> = (props) => {
  return <h1 class="py-2 text-4xl font-400">{props.children}</h1>;
};
