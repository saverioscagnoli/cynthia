import { BsDiscord, BsGithub } from "react-icons/bs";
import { LuMoon, LuSun } from "react-icons/lu";
import { useNavigate } from "react-router";
import { Button } from "~/components/ui";
import { cn } from "~/lib/utils";
import { useAuth } from "~/contexts/auth";
import { useTheme } from "~/contexts/theme";
import { TopbarCenterSection } from "./center";
import { UserDisplay } from "./user";

const Topbar = () => {
  const { logged } = useAuth();
  const { theme, toggle: toggleTheme } = useTheme();
  const nav = useNavigate();

  return (
    <div
      className={cn(
        "h-20 w-full",
        "flex items-center justify-between",
        "sm:px-[5%] md:px-[7%] lg:px-[10%] xl:px-[20%]",
        "sticky top-0 left-0 z-9998 backdrop-blur-md"
      )}
    >
      <img
        src="/cynthia-overworld.png"
        width={64}
        height={64}
        className={cn("mb-3", "cursor-pointer")}
        style={{ imageRendering: "pixelated" }}
        onClick={() => nav("/")}
      />
      <TopbarCenterSection />
      <div className={cn("flex items-center gap-4")}>
        {logged ? (
          <UserDisplay />
        ) : (
          <a href="/auth/login">
            <Button colorScheme="blue" leftIcon={<BsDiscord />}>
              Login with Discord
            </Button>
          </a>
        )}
        <a href="https://github.com/saverioscagnoli/cynthia" target="_blank">
          <Button size="icon" variant="ghost" colorScheme="gray">
            <BsGithub size={20} />
          </Button>
        </a>
        <Button
          size="icon"
          variant="ghost"
          colorScheme="gray"
          onClick={toggleTheme}
        >
          {theme === "dark" ? <LuSun size={20} /> : <LuMoon size={20} />}
        </Button>
      </div>
    </div>
  );
};

export { Topbar };
