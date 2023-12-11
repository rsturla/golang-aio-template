"use client";
import useSWR from "swr";

async function fetcher(url: string) {
  const resp = await fetch(url);
  return resp.text();
}

export default function Home() {
  const { data, error } = useSWR("/api", fetcher, { refreshInterval: 1000 });

  return (
    <main>
      <div>
        <h1>Hello World!</h1>
        <h2>Memory allocation stats from GoLang API</h2>
        {error && (
          <div>
            Failed to load: <strong>{error}</strong>
          </div>
        )}
        {!error && !data && <p>Loading...</p>}
        {!error && data && <pre>{data}</pre>}
      </div>
    </main>
  );
}
