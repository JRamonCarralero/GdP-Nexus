import Head from 'next/head';


export default async function Home() {
  const res = await fetch(process.env.NEXT_PUBLIC_API_URL!);
  const data = await res.text();

  return (
    <div>
      <Head>
        <title>Mi App Next.js</title>
      </Head>

      <main>
        <h1>Mensaje del Backend</h1>
        <p>{data}</p>
      </main>
    </div>
  );
}