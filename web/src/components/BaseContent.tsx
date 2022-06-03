import { ComponentWithChildren } from "../types/ComponentWithChildren";

export const BaseContent: ComponentWithChildren<{}> = (props) => {
  return <div class="flex flex-col gap-y-2" {...props} />;
};
