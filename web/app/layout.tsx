import type { Metadata } from "next";
import { Inter, Roboto_Mono } from "next/font/google";
import "./globals.css";
import { Toaster } from "react-hot-toast";

const inter = Inter({
  variable: "--font-inter",
  subsets: ["latin"],
});

const robotoMono = Roboto_Mono({
  variable: "--font-roboto-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Linktok - Shorten Links Instantly | Free URL Shortener",
  description:
    "Linktok is a modern URL shortener that creates concise, trackable links. Get real-time analytics, custom short URLs, and QR codes for your marketing campaigns.",
  keywords: [
    "URL shortener",
    "link shortener",
    "free link shortening",
    "custom URLs",
    "link analytics",
    "QR code generator",
    "marketing tools",
  ],
  openGraph: {
    title: "Linktok - Professional URL Shortening Solution",
    description:
      "Transform long URLs into memorable, shareable links. Perfect for social media and digital marketing campaigns.",
    url: "https://linktok.com",
    siteName: "Linktok",
    locale: "en_US",
    type: "website",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${inter.variable} ${robotoMono.variable} antialiased`}>
        <Toaster />
        {children}
      </body>
    </html>
  );
}
