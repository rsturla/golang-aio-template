"use client";

import { ThemeProvider } from "@/components/theme-provider";
import "@/styles/globals.css";
import { MainNav } from "@/components/main-nav";
import { mainNavConfig } from "@/config/site";
import { cn } from "@/lib/utils";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={cn("min-h-screen bg-background font-sans antialiased")}>
        <ThemeProvider
          attribute="class"
          defaultTheme="dark"
          enableSystem
          disableTransitionOnChange
        >
          <header className="container z-40">
            <div className="h-20">
              <MainNav items={mainNavConfig} />
            </div>
          </header>
          <main className="container py-2">{children}</main>
        </ThemeProvider>
      </body>
    </html>
  );
}
