import { User } from "firebase/auth";
import { createEffect, createSignal } from "solid-js";
import { auth as fbAuth } from "../firebase";

const [email, setEmail] = createSignal("");
const [token, setToken] = createSignal("");

createEffect(() => {
  const cleanup = fbAuth.onAuthStateChanged((user: User | null) => {
    const email = user?.email ?? "";
    console.log("email", email);
    setEmail(email);
  });
  return () => cleanup();
});

createEffect(() => {
  const cleanup = fbAuth.onIdTokenChanged(async (user: User | null) => {
    const token = (await user?.getIdToken()) ?? "";
    console.log("token", token);
    setToken(token);
  });
  return () => cleanup();
});

export const auth = [email, token];
