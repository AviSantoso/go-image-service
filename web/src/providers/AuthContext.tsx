import { User } from "firebase/auth";
import { createContext, createEffect, useContext } from "solid-js";
import { createStore } from "solid-js/store";
import { auth } from "../firebase";

export type Auth = {
  email: string;
  token: string;
};

const defaultAuth: Auth = { email: "", token: "" };

export const AuthContext = createContext<Auth>(defaultAuth);

export const AuthProvider = ({ children }: any) => {
  const [state, setState] = createStore({ email: "", token: "" });

  createEffect(() => {
    const cleanup = auth.onAuthStateChanged((user: User | null) => {
      const email = user?.email ?? "";
      setState("email", email);
      console.log("email", email);
    });
    return () => cleanup();
  });

  createEffect(() => {
    const cleanup = auth.onIdTokenChanged(async (user: User | null) => {
      const token = (await user?.getIdToken()) ?? "";
      console.log("token", token);
      setState("token", token);
    });
    return () => cleanup();
  });

  return <AuthContext.Provider value={state}>{children}</AuthContext.Provider>;
};

export const useAuth = () => useContext(AuthContext);
