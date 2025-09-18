export default function Home() {
  return (
    <div className="main-bg">
      <div className="w-full flex flex-row justify-end items-center">
        <nav className="flex flex-row justify-around items-center min-w-[200px]">
          <a href="/login" className="text-black">Login</a>
          <a href="/register" className="text-black">Register</a>
        </nav>
      </div>
    </div>
  );
}
