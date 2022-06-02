import { Component, JSX, Show } from "solid-js";

import { useAuth } from "../hooks/useAuth";

interface ProtectedPageProps {
  children: JSX.Element;
}

export const ProtectedPage: Component<ProtectedPageProps> = (props) => {
  const { loading, email } = useAuth();

  return (
    <Show when={!loading()} fallback={<div>Loading authentication...</div>}>
      <Show
        when={email()}
        fallback={
          <section class="bg-pink-100 text-gray-700 p-8">
            <h1 class="text-2xl font-bold">Unauthorized</h1>
            <p class="mt-4">Please login before accessing this website.</p>
          </section>
        }
      >
        {props.children}
      </Show>
    </Show>
  );
};
