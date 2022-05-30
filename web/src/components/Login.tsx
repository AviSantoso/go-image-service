import { signInWithGoogle } from "../firebase";
import { useAuth } from "../providers/AuthContext";

export const Login = () => {
  const { email } = useAuth();

  return (
    <div style={{ display: "flex", "flex-direction": "column" }}>
      {email.length
        ? `Currently logged in as ${email}.`
        : `Please login before continuing.`}
      {!email.length && (
        <button style={{ "max-width": "5em" }} onClick={signInWithGoogle}>
          Login
        </button>
      )}
    </div>
  );
};
