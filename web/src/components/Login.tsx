import { signInWithGoogle } from "../firebase";
import { auth } from "../providers/AuthContext";

export function Login() {
  const [email] = auth;

  return (
    <div style={{ display: "flex", "flex-direction": "column" }}>
      {email().length
        ? `Currently logged in as ${email()}.`
        : `Please login before continuing.`}
      {!email().length && (
        <button style={{ "max-width": "5em" }} onClick={signInWithGoogle}>
          Login
        </button>
      )}
    </div>
  );
}
