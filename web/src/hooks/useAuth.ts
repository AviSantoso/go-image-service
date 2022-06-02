import { User } from "firebase/auth";
import { createEffect, createSignal } from "solid-js";

import { auth as fbAuth } from "../firebase";

const [loading, setLoading] = createSignal(true);
const [error, setError] = createSignal("");
const [email, setEmail] = createSignal("");
const [token, setToken] = createSignal("");
const [loggedIn, setLoggedIn] = createSignal(false);

createEffect(() => {
  try {
    const cleanup = fbAuth.onAuthStateChanged((user: User | null) => {
      const email = user?.email ?? "";
      setEmail(email);
    });
    setLoading(false);
    return () => cleanup();
  } catch (error) {
    setError(error.message ?? error.toString());
  }
});

createEffect(() => {
  try {
    const cleanup = fbAuth.onIdTokenChanged(async (user: User | null) => {
      const token = (await user?.getIdToken()) ?? "";
      setToken(token);
    });
    setLoading(false);
    return () => cleanup();
  } catch (error) {
    setError(error.message ?? error.toString());
  }
});

createEffect(() => {
  setLoggedIn(() => Boolean(email() && token()));
});

export function useAuth() {
  return { loading, error, loggedIn, email, token };
}
