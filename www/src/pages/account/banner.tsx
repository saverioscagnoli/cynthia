import { useEffect, useRef, useState } from "react";
import { Cross1Icon, Pencil1Icon } from "@radix-ui/react-icons";
import { Button } from "~/components/ui";
import { deleteBanner, uploadBanner } from "~/lib/backend";
import { cn } from "~/lib/utils";
import { useAuth } from "~/contexts/auth";
import { useProfileEdit } from "~/contexts/profile-edit";
import { useTheme } from "~/contexts/theme";

const UserBanner = () => {
  const { token } = useAuth();
  const { theme } = useTheme();
  const { editing, stopEditing } = useProfileEdit();
  const [bannerURL, setBannerURL] = useState<string | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    fetch("/user/banner", {
      headers: { Authorization: `Bearer ${token}` }
    })
      .then(res => {
        if (!res.ok) throw new Error("no banner");
        return res.blob();
      })
      .then(blob => setBannerURL(URL.createObjectURL(blob)))
      .catch(() => setBannerURL(null));
  }, []);

  useEffect(() => {
    return () => {
      if (bannerURL) URL.revokeObjectURL(bannerURL);
    };
  }, [bannerURL]);

  const onEditClick = () => {
    fileInputRef.current?.click();
  };

  const onBannerChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];

    if (!file || !token) return;

    const localURL = URL.createObjectURL(file);
    setBannerURL(localURL);
    stopEditing();
    uploadBanner(token, file);
  };

  const onBannerDelete = async () => {
    try {
      await deleteBanner(token);
    } catch (e) {
      console.error(e);
      return;
    }

    setBannerURL(null);
    stopEditing();
  };

  return (
    <div className={cn("absolute inset-0")}>
      <img
        className={cn("h-full w-full", "object-cover", "rounded-2xl")}
        src={
          bannerURL ?? (theme === "dark" ? "/pattern-dark.svg" : "/pattern.svg")
        }
      />
      {editing && (
        <>
          <input
            ref={fileInputRef}
            type="file"
            className="hidden"
            accept="image/*"
            onChange={onBannerChange}
          />
          <Button
            className="absolute top-4 right-4"
            size="icon"
            variant="soft"
            colorScheme="red"
            onClick={onEditClick}
          >
            <Pencil1Icon />
          </Button>
          <Button
            className={cn("absolute top-16 right-4")}
            size="icon"
            variant="soft"
            colorScheme="red"
            onClick={onBannerDelete}
          >
            <Cross1Icon />
          </Button>
        </>
      )}
    </div>
  );
};

export { UserBanner };
