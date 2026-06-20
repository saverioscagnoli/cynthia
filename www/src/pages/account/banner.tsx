import { useEffect, useState } from "react";
import { Pencil1Icon } from "@radix-ui/react-icons";
import { Button } from "~/components/ui";
import { uploadBanner } from "~/lib/backend";
import { cn } from "~/lib/utils";
import { useAuth, type User } from "~/contexts/auth";
import { useTheme } from "~/contexts/theme";

const UserBanner = () => {
  const { token } = useAuth();
  const { theme } = useTheme();
  const [bannerURL, setBannerURL] = useState<string | null>(null);

  useEffect(() => {
    fetch("/user/banner", {
      headers: { Authorization: `Bearer ${token}` }
    })
      .then(res => res.blob())
      .then(blob => setBannerURL(URL.createObjectURL(blob)))
      .catch(() => setBannerURL(null));
  }, []);

  useEffect(() => {
    return () => {
      if (bannerURL) URL.revokeObjectURL(bannerURL);
    };
  }, [bannerURL]);

  const onBannerChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];

    if (!file || !token) return;

    uploadBanner(token, file);
  };

  return (
    <div className={cn("absolute inset-0", "group")}>
      <img
        className={cn("h-full w-full", "object-cover", "rounded-2xl")}
        src={
          bannerURL ?? (theme === "dark" ? "/pattern-dark.svg" : "/pattern.svg")
        }
      />
      <label className="absolute top-4 right-4 cursor-pointer">
        <input
          type="file"
          className="hidden"
          accept="image/*"
          onChange={onBannerChange}
        />
        <Button
          className={cn(
            "opacity-0 transition-opacity duration-200 group-hover:opacity-100",
            "pointer-events-none"
          )}
          size="icon"
          variant="ghost"
          colorScheme="gray"
        >
          <Pencil1Icon />
        </Button>
      </label>
    </div>
  );
};

export { UserBanner };
