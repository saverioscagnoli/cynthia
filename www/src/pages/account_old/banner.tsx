import { useRef, useState } from "react";
import { Cross1Icon, Pencil1Icon } from "@radix-ui/react-icons";
import { Button } from "~/components/ui";
import { cn } from "~/lib/utils";
import { privateApi } from "~/lib/wrapper";
import { useAccount } from "~/contexts/account";
import { useAuth } from "~/contexts/auth";
import { useTheme } from "~/contexts/theme";

const UserBanner = () => {
  const { user, isEditing, stopEdit } = useAccount();
  const { token } = useAuth();
  const { theme } = useTheme();
  const fileInputRef = useRef<HTMLInputElement>(null);

  const [a, setA] = useState(1);

  const onEditClick = () => {
    fileInputRef.current?.click();
  };

  const onBannerChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file || !token) return;

    stopEdit();
    await privateApi.updateBanner(token, file);
    setA(a => a + 1); // bust cache
  };

  const onBannerDelete = async () => {
    if (!token) return;

    stopEdit();
    await privateApi.deleteBanner(token);
    setA(a => a + 1); // bust cache
  };

  return (
    <div className={cn("absolute inset-0")}>
      <img
        key={`${a}-${theme}`}
        className={cn("h-full w-full", "object-cover", "rounded-2xl")}
        src={`http://localhost:3247/api/user/${user.id}/banner?v=${a}`}
        onError={e => {
          e.currentTarget.src =
            theme === "dark" ? "/pattern-dark.svg" : "/pattern.svg";
        }}
      />
      {isEditing && (
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
