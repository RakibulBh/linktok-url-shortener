"use client";

import { useEffect, useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { ArrowRight, Link, Copy, Sparkles, Link2 } from "lucide-react";
import Navbar from "@/components/navbar";
import HeroTag from "@/components/hero-tag";
import abstractArt from "../public/images/abstract-art.png";
import Image from "next/image";
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
