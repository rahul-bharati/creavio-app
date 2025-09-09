import type { Metadata } from "next";
import { Inter, Raleway } from "next/font/google";
import "./globals.css";

const interFont = Inter({
  variable: "--font-inter",
  subsets: ["latin"],
});

const ralewayFont = Raleway({
  variable: "--font-raleway",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Creavio",
  description: "Secure chat, voice, video, and streaming—all in one place. No tracking. No limits",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${interFont.variable} ${ralewayFont.variable} antialiased`}
      >
        {children}
      </body>
    </html>
  );
}
