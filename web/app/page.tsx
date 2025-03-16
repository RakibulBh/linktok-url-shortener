"use client";

import { useEffect } from "react";
import Navbar from "@/components/navbar";
import Result from "@/components/sections/Result";
import Hero from "@/components/sections/Hero";
import { useResultStore, useUrlStore } from "@/store/urlStore";

export default function Home() {
  const { result, setResult } = useResultStore();
  const { url } = useUrlStore();

  useEffect(() => {
    if (url.trim() === "") {
      setResult("");
      return;
    }
    if (result && url !== "") {
      window.scrollTo({
        top: document.getElementById("result")?.offsetTop,
        behavior: "smooth",
      });
    }
  }, [result, url]);

  return (
    <main className="overflow-x-hidden">
      <Navbar />
      <Hero />
      {result && url !== "" && <Result />}
    </main>
  );
}
