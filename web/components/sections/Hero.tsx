"use client";
import React from "react";
import { motion } from "framer-motion";
import HeroTag from "../hero-tag";
import InputForm from "../input-form";

const Hero = () => {
  return (
    <section
      id="hero"
      className="relative h-[110vh] flex flex-col items-center bg-gradient-to-b from-[#F9FAFB] via-[#FCE7F3] to-[#F3E8FF]"
    >
      <div className="absolute inset-0 overflow-hidden">
        <motion.div
          className="absolute -top-20 -left-20 w-96 h-96 bg-purple-100/30 rounded-full blur-3xl"
          animate={{ rotate: 360 }}
          transition={{ duration: 20, repeat: Infinity, ease: "linear" }}
        />
        <motion.div
          className="absolute top-1/3 right-0 w-64 h-64 bg-pink-100/30 rounded-full blur-3xl"
          animate={{ scale: [1, 1.2, 1], y: [0, -40, 0] }}
          transition={{ duration: 8, repeat: Infinity }}
        />
      </div>

      <motion.div
        className="max-w-4xl px-4 mt-24 md:mt-32 gap-6 flex flex-col items-center"
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.8 }}
      >
        <motion.div
          animate={{ y: [-5, 5, -5] }}
          transition={{ duration: 4, repeat: Infinity }}
        >
          <HeroTag />
        </motion.div>

        <div className="text-center space-y-4">
          <motion.h1
            className="text-4xl md:text-6xl font-bold bg-gradient-to-r from-purple-600 to-pink-600 bg-clip-text text-transparent"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ delay: 0.2 }}
          >
            Shorten Links with Linktok
          </motion.h1>
          <motion.p
            className="text-lg md:text-xl text-purple-900/80 font-medium"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ delay: 0.4 }}
          >
            Transform lengthy URLs into memorable links that spark curiosity
          </motion.p>
        </div>

        <motion.div
          className="w-full max-w-md relative group"
          initial={{ scale: 0.95 }}
          animate={{ scale: 1 }}
          transition={{ delay: 0.6 }}
        >
          <div className="absolute -inset-1 bg-gradient-to-r from-purple-600 to-pink-600 rounded-full blur opacity-30 group-hover:opacity-50 transition-opacity" />
          <InputForm />
        </motion.div>

        <motion.div
          className="mt-8 flex gap-4 opacity-70"
          animate={{ rotate: [0, -5, 5, 0] }}
          transition={{ duration: 2, repeat: Infinity }}
        >
          <div className="w-8 h-8 rounded-full bg-purple-100" />
          <div className="w-8 h-8 rounded-full bg-pink-100" />
          <div className="w-8 h-8 rounded-full bg-purple-100" />
        </motion.div>
      </motion.div>

      <motion.div
        className="absolute bottom-0 w-full h-48 bg-gradient-to-t from-white via-white/50 to-transparent"
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ duration: 1 }}
      />
    </section>
  );
};

export default Hero;
