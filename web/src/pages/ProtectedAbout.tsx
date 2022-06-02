import { Component } from "solid-js";

import { ProtectedPage } from "../components/ProtectedPage";
import { useAuth } from "../hooks/useAuth";

export const ProtectedAbout: Component = () => {
  const { email } = useAuth();
  return (
    <ProtectedPage>
      <section class="bg-green-100 text-gray-700 p-8">
        <h1 class="text-2xl font-bold">Protected About</h1>

        <p class="mt-4">This is a protected page.</p>
        <p>You're currently logged in as {email()}.</p>
      </section>
    </ProtectedPage>
  );
};

export default ProtectedAbout;
