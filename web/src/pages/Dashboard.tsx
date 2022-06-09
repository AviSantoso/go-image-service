import { Component } from "solid-js";

import { ProtectedPage } from "../components/Protected";
import { useAuth } from "../hooks/useAuth";

interface DashboardLinkProps {
  icon: string;
  text: string;
}

export const DashboardLink: Component<DashboardLinkProps> = (props) => {
  return (
    <div class="flex flex-1 flex-row p-2 rounded hover:cursor-pointer hover:bg-blue-100">
      <div class="flex h-full min-w-2em">
        <i class={`fa-solid fa-${props.icon}`} />
      </div>
      <div class="flex-auto h-full">{props.text}</div>
    </div>
  );
};

export const Dashboard: Component = () => {
  const { email } = useAuth();
  return (
    <ProtectedPage fullWidth>
      <div class="grid grid-cols-5 gap-4 w-1/1 h-1/1">
        <div class="flex flex-col gap-10">
          <div class="flex flex-col gap-2">
            <DashboardLink icon="upload" text="Upload" />
            <DashboardLink icon="link" text="From Link" />
          </div>
          <div class="flex flex-col gap-2">
            <DashboardLink icon="hard-drive" text="Drive" />
            <DashboardLink icon="person" text="Shared" />
            <DashboardLink icon="clock" text="Recent" />
            <DashboardLink icon="star" text="Starred" />
            <DashboardLink icon="trash" text="Trash" />
            <DashboardLink icon="archive" text="Archived" />
          </div>
          <div class="flex flex-col gap-2">
            <DashboardLink icon="floppy-disk" text="Drive Size" />
            <DashboardLink icon="database" text="Archive Size" />
          </div>
        </div>
        <div>Content bar</div>
      </div>
    </ProtectedPage>
  );
};

export default Dashboard;
