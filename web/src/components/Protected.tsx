import { Component, JSX, Show } from "solid-js";

import { useAuth } from "../hooks/useAuth";
import { BasePage } from "./BasePage";

interface ProtectedPageProps {
  children: JSX.Element;
}

export const ProtectedPage: Component<ProtectedPageProps> = (props) => {
  const { loading, email } = useAuth();

  return (
    <BasePage>
      <Show when={!loading()} fallback={<div>Loading...</div>}>
        <Show
          when={email()}
          fallback={
            <>
              <h1 class="text-2xl font-bold">Unauthorized</h1>
              <p class="mt-4">Please login before accessing this website.</p>
            </>
          }
        >
          {props.children}
        </Show>
      </Show>
    </BasePage>
  );
};
