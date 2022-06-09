import { Component } from "solid-js";

import { ProtectedPage } from "../components/Protected";
import { useAuth } from "../hooks/useAuth";

interface DashboardLinkProps {
  icon: string;
  text: string;
}

export const DashboardLink: Component<DashboardLinkProps> = (props) => {
  return (
    <div class="flex flex-1 flex-row p-2 rounded bg-white hover:cursor-pointer hover:bg-blue-400 hover:text-slate-50">
      <div class="flex min-w-2em items-center">
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
      <div class="flex flex-col gap-4">
        <div class="gap-4 px-4 w-1/1">
          <input
            class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="username"
            type="text"
            placeholder="Search"
          />
        </div>
        <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6 gap-4 px-4 w-1/1 h-1/1">
          <div class="flex flex-col gap-6 bg-slate-100 p-2">
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
              <DashboardLink icon="floppy-disk" text="Drive 56mb/1024mb" />
              <DashboardLink icon="database" text="Archive" />
            </div>
          </div>
          <div class="flex flex-col gap-2 col-start-2 col-end-3 sm:col-end-4 md:col-end-5 lg:col-end-7">
            <input
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="username"
              type="text"
              placeholder="Filter"
            />
            <table class="table-auto md:table-fixed w-1/1 text-left">
              <thead>
                <tr>
                  <th>Name</th>
                  <th>Owner</th>
                  <th>Modified</th>
                  <th>Extension</th>
                  <th>File Size</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>The Sliding Mr. Bones (Next Stop, Pottersville)</td>
                  <td>Malcolm Lockyer</td>
                  <td>1961</td>
                  <td>1961</td>
                  <td>512kb</td>
                </tr>
                <tr>
                  <td>The Sliding Mr. Bones (Next Stop, Pottersville)</td>
                  <td>Malcolm Lockyer</td>
                  <td>1961</td>
                  <td>1961</td>
                  <td>512kb</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
        <div class="flex justify-between gap-4 px-4 w-1/1 text-center text-slate-300">
          <span>Copyright Go Drive 2022</span>
          <span>Logged in as {email()}</span>
          <span>Built by Aviciena Santoso</span>
        </div>
      </div>
    </ProtectedPage>
  );
};

export default Dashboard;
