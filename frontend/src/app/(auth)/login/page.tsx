import { LoginForm } from "@/app/components/auth/login-form";
import { login } from "./actions";

export default function Login() {
  return (
    <div>
      <h1>Login</h1>
      <LoginForm/>
    </div>
  );
}