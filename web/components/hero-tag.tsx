import { ArrowRight, Link } from "lucide-react";
import React from "react";

const HeroTag = () => {
  return (
    <div className="group w-64 bg-gradient-to-r from-purple-50 to-pink-50 p-3 gap-3 rounded-xl shadow-md hover:shadow-lg transition-all duration-300 flex items-center justify-between border border-purple-100">
      <Link
        size={20}
        className="text-purple-300 group-hover:text-purple-500 transition-colors"
      />
      <p className="text-sm text-gray-600 group-hover:text-gray-900 transition-colors font-medium">
        Linktok on the web!
      </p>
      <ArrowRight
        size={20}
        className="text-purple-300 group-hover:text-purple-500 transition-colors"
      />
    </div>
  );
};

export default HeroTag;
