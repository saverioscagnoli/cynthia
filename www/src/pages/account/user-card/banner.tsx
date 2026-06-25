import { useEffect, useRef, useState, type ChangeEvent } from "react";
import { LuImage, LuTrash } from "react-icons/lu";
import { Button } from "~/components/ui";
import { cn } from "~/lib/utils";
import { privateApi } from "~/lib/wrapper";
import { useAccount } from "~/contexts/account";
import { useTheme } from "~/contexts/theme";

const Banner = () => {
  const [hasEdited, setHasEdited] = useState<boolean>(false);
  const { user, updateUser, isEditing, onEditConfirm, registerOnEditConfirm } =
    useAccount();
  const { theme } = useTheme();
  const [localPreview, setLocalPreview] = useState<string | null>(null);

  const inputRef = useRef<HTMLInputElement>(null);
  const typeRef = useRef<"edit" | "clear" | null>(null);
  const fileRef = useRef<File | null>(null);

  const onEditBannerClick = () => {
    typeRef.current = "edit";
    inputRef.current?.click();
  };

  const onBannerFileChange = async (e: ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0] ?? null;
    fileRef.current = file;

    if (file) setLocalPreview(URL.createObjectURL(file)); // show instantly
    onEditConfirm();
  };

  const onClearBannerClick = () => {
    typeRef.current = "clear";
    onEditConfirm();
  };

  useEffect(() => {
    if (isEditing) setHasEdited(true);
  }, [isEditing]);

  useEffect(() => {
    return registerOnEditConfirm(async () => {
      if (typeRef.current === "clear") {
        await privateApi.deleteBanner();
        updateUser({ banner: null });
      } else if (typeRef.current === "edit" && fileRef.current) {
        await privateApi.updateBanner(fileRef.current);
        updateUser({ banner: [1] });
      }

      typeRef.current = null;
    });
  }, []);

  const bannerUrl = localPreview
    ? localPreview
    : user.banner
      ? `${__API_URL__}/api/user/${user.id}/banner`
      : theme === "dark"
        ? "/pattern-dark.svg"
        : "/pattern.svg";

  return (
    <div className={cn("absolute inset-0")}>
      <img
        className={cn("h-full w-full", "object-cover", "-z-1")}
        src={bannerUrl}
      />
      <div
        className={cn("flex flex-col gap-4", "absolute top-4 right-4")}
        style={{
          opacity: 0,
          pointerEvents: isEditing ? "all" : "none",
          animation: isEditing
            ? "fadeIn 0.2s ease forwards"
            : hasEdited
              ? "fadeOut 0.2s ease forwards"
              : "none"
        }}
      >
        <Button
          variant="soft"
          colorScheme="red"
          leftIcon={<LuImage />}
          onClick={onEditBannerClick}
        >
          Edit Banner
        </Button>
        <input
          type="file"
          accept="image/*"
          className={cn("absolute hidden")}
          onChange={onBannerFileChange}
          ref={inputRef}
        />
        <Button
          variant="soft"
          colorScheme="red"
          leftIcon={<LuTrash />}
          onClick={onClearBannerClick}
        >
          Clear Banner
        </Button>
      </div>
    </div>
  );
};

export { Banner };
