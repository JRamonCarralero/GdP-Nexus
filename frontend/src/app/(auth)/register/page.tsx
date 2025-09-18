export default function Register() {
  return (
    <div>
      <h1>Register</h1>
      <form>
        <input type="email" placeholder="Email" />
        <input type="password" placeholder="Password" />
        <input type="text" placeholder="First Name" />
        <input type="text" placeholder="Last Name" />
        <input type="text" placeholder="Nickname" />
        <button type="submit">Register</button>
      </form>
    </div>
  );
}