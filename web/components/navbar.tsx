import Link from "next/link";
import React from "react";

const Navbar = () => {
  return (
    <nav className="px-6 md:px-12 py-4 backdrop-blur-lg w-full fixed top-0 z-50">
      <div className="flex items-center justify-between max-w-7xl mx-auto">
        {/* Logo */}
        <Link
          href="/"
          className="text-2xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-purple-600 to-pink-600"
        >
          Linktok
        </Link>

        {/* Get Started Button */}
        <button className="bg-gradient-to-r from-purple-600 to-pink-600 text-white px-5 py-2 rounded-full text-sm hover:from-purple-700 hover:to-pink-700 transition-all shadow-sm hover:shadow-md">
          Get Started
        </button>
      </div>
    </nav>
  );
};

export default Navbar;
