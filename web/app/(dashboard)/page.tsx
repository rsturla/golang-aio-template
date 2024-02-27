"use client";
import useSWR from "swr";

interface Data {
  count: number;
}

async function fetcher(url: string): Promise<string> {
  const resp = await fetch(url);
  if (!resp.ok) {
    throw new Error("Failed to fetch data");
  }
  return resp.text();
}

export default function Home() {
  const { data, error } = useSWR("/api/count", fetcher, { refreshInterval: 1000 });

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  if (!data) {
    return <div>Loading...</div>;
  }

  const parsedData: Data = JSON.parse(data);

  return (
    <div>
      <h2>Counter service from GoLang API</h2>
      <h1>Count: {parsedData.count}</h1>
    </div>
  );
}
