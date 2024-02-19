import { MainNavItem } from "@/lib/types";
import { useSelectedLayoutSegment } from "next/navigation";
import * as React from "react";
import Link from "next/link";
import { Icons } from "@/components/icons";
import { siteConfig } from "@/config/site";
import { cn } from "@/lib/utils";
import { MobileNav } from "@/components/mobile-nav";

interface MainNavProps {
  items?: MainNavItem[];
  children?: React.ReactNode;
}

export function MainNav({ items, children }: MainNavProps) {
  const segment = useSelectedLayoutSegment();
  const [showMobileMenu, setShowMobileMenu] = React.useState<boolean>(false);

  return (
    <div className="flex h-16 items-center w-full">
      {/* LEFT */}
      <div className="flex w-6/12 space-x-8">
        <Link href="/" passHref className="flex space-x-3">
          <Icons.logo />
          <span className="font-bold">{siteConfig.name}</span>
        </Link>
        {items?.length ? (
          <nav className="hidden gap-6 md:flex">
            {items?.map((item, index) => (
              <Link
                key={index}
                href={item.disabled ? "#" : item.href}
                className={cn(
                  "flex items-center transition-colors hover:text-foreground/90 sm:text-sm",
                  item.href.startsWith(`/${segment}`)
                    ? "text-foreground"
                    : "text-foreground/60",
                  item.disabled && "cursor-not-allowed opacity-80",
                )}
              >
                {item.title}
              </Link>
            ))}
          </nav>
        ) : null}
      </div>

      {/* RIGHT */}
      <div className="md:hidden w-6/12 flex items-center justify-end">
        <button
          className="flex items-center space-x-2"
          onClick={() => setShowMobileMenu((prev) => !prev)}
        >
          {showMobileMenu ? <Icons.close /> : <Icons.hamburger />}
        </button>
        {showMobileMenu && items && (
          <MobileNav items={items}>{children}</MobileNav>
        )}
      </div>
    </div>
  );
}
