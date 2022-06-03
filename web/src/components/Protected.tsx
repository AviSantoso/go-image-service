import { Component, JSX, Show } from "solid-js";

import { useAuth } from "../hooks/useAuth";
import { BaseContent } from "./BaseContent";
import { BasePage } from "./BasePage";
import { Title } from "./Title";

interface ProtectedPageProps {
  children: JSX.Element;
}

export const ProtectedPage: Component<ProtectedPageProps> = (props) => {
  const { loading, email } = useAuth();

  return (
    <BasePage>
      <Show when={!loading()} fallback={<BaseContent>Loading...</BaseContent>}>
        <Show
          when={email()}
          fallback={
            <BaseContent>
              <Title>Protected Page</Title>
              <p>Please login before accessing this page.</p>
            </BaseContent>
          }
        >
          {props.children}
        </Show>
      </Show>
    </BasePage>
  );
};
