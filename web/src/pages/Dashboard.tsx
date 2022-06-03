import { Component } from "solid-js";

import { ProtectedPage } from "../components/Protected";
import { useAuth } from "../hooks/useAuth";

export const Dashboard: Component = () => {
  const { email } = useAuth();
  return (
    <ProtectedPage>
      <div class="flex flex-col gap-y-2">
        <h1 class="text-2xl font-bold">Dashboard</h1>
        <p>You're currently logged in as {email()}.</p>
      </div>
    </ProtectedPage>
  );
};

export default Dashboard;
