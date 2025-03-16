"use client";
import React, { useState } from "react";
import { motion } from "framer-motion";
import { Copy, Link2, Sparkles } from "lucide-react";
import { useResultStore } from "@/store/urlStore";

const Result = () => {
  const [copied, setCopied] = useState(false);
  const { result, setResult } = useResultStore();

  const handleCopy = () => {
    navigator.clipboard.writeText(result);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <section
      id="result"
      className="relative min-h-screen flex flex-col items-center bg-gradient-to-b from-[#F3E8FF] via-[#FCE7F3] to-[#F9FAFB]"
    >
      <motion.div
        className="max-w-4xl px-4 mt-16 md:mt-32 gap-8 flex flex-col items-center"
        initial={{ opacity: 0, y: 40 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.8 }}
      >
        <div className="text-center space-y-4">
          <h2 className="text-3xl md:text-5xl font-bold bg-gradient-to-r from-purple-600 to-pink-600 bg-clip-text text-transparent">
            Your Digital Business Card
          </h2>
          <p className="text-purple-900/80 font-medium">
            Share your link anywhere - it works like magic
          </p>
        </div>

        <motion.div
          className="w-full max-w-xl relative group"
          whileHover={{ scale: 1.02 }}
        >
          <div className="absolute -inset-1 bg-gradient-to-r from-purple-100 to-pink-100 rounded-2xl blur opacity-50" />
          <div className="relative flex items-center bg-white rounded-2xl shadow-lg p-1">
            <Link2 className="ml-4 text-purple-400" size={20} />
            <input
              readOnly
              value={result}
              className="w-full py-4 px-6 bg-transparent text-gray-900 focus:outline-none"
            />
            <motion.button
              onClick={handleCopy}
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
              className="m-1 px-6 py-2 bg-gradient-to-r from-purple-600 to-pink-600 text-white rounded-xl font-medium flex items-center gap-2"
            >
              <Copy size={18} />
              {copied ? "Copied!" : "Copy"}
            </motion.button>
          </div>
        </motion.div>

        <motion.div
          className="flex flex-wrap gap-4 justify-center"
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ delay: 0.4 }}
        >
          <motion.button
            onClick={() => {
              window.scrollTo({
                top: document.getElementById("hero")?.offsetTop,
                behavior: "smooth",
              });
              setTimeout(() => {
                setResult("");
              }, 1000);
            }}
            whileHover={{ y: -2 }}
            className="px-6 py-2.5 rounded-full bg-white text-purple-600 hover:bg-purple-50 transition-colors font-medium flex items-center gap-2 shadow-sm hover:shadow-md"
          >
            <Sparkles size={16} />
            Create Another
          </motion.button>
        </motion.div>
      </motion.div>
    </section>
  );
};

export default Result;
