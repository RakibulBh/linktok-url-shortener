"use client";

import { useResultStore, useUrlStore } from "@/store/urlStore";
import { Link, Sparkles } from "lucide-react";
import { motion } from "framer-motion";
import React from "react";
import { shortenUrl } from "@/services/requests";
import toast from "react-hot-toast";

const InputForm = () => {
  const { setUrl, url } = useUrlStore();
  const { setResult } = useResultStore();

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const response = await shortenUrl(url);

    // Validate each link should start with http:// or https://
    if (!url.startsWith("http://") && !url.startsWith("https://")) {
      toast.error("Link should start with http:// or https://");
      return;
    }

    if (response.error) {
      toast.error(response.message);
      return;
    }

    setResult(response.data);
  };

  return (
    <form
      onSubmit={onSubmit}
      className="relative flex items-center bg-white rounded-lg shadow-lg hover:shadow-xl transition-shadow"
    >
      <Link className="ml-4 text-purple-400" size={20} />
      <input
        value={url}
        onChange={(e) => setUrl(e.target.value)}
        placeholder="Paste your link here"
        className="w-full py-4 px-6 bg-transparent text-gray-900 placeholder:text-gray-400 focus:outline-none"
      />
      <motion.button
        type="submit"
        whileHover={{ scale: 1.05 }}
        whileTap={{ scale: 0.95 }}
        className="m-1 px-6 py-2 bg-gradient-to-r from-purple-600 to-pink-600 text-white rounded-lg font-medium flex items-center gap-2"
      >
        <Sparkles size={18} />
        Shorten
      </motion.button>
    </form>
  );
};

export default InputForm;
