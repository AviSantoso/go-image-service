import { Component } from "solid-js";

import { BaseContent } from "../components/BaseContent";
import { ProtectedPage } from "../components/Protected";
import { Title } from "../components/Title";
import { useAuth } from "../hooks/useAuth";

export const Dashboard: Component = () => {
  const { email } = useAuth();
  return (
    <ProtectedPage>
      <BaseContent>
        <Title>Dashboard</Title>
        <p>You're currently logged in as {email()}.</p>
      </BaseContent>
    </ProtectedPage>
  );
};

export default Dashboard;
