import { useEffect, useRef, useState, type ComponentProps } from "react";
import { Avatar } from "~/components/ui";
import { cn, dsAvatar } from "~/lib/utils";
import { privateApi } from "~/lib/wrapper";
import { useAccount } from "~/contexts/account";
import { useAuth } from "~/contexts/auth";

const UserInfo: React.FC<ComponentProps<"div">> = ({ className, ...props }) => {
  const { user, isEditing, stopEdit, registerOnEditConfirm, updateUser } =
    useAccount();
  const { updateLoggedUser, token } = useAuth();
  const [username, setUsername] = useState<string>(user.username);
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    return registerOnEditConfirm(async () => {
      if (username.length < 3 || username.length > 23) {
        setUsername(user.username);
        stopEdit();
        return;
      }

      updateUser({ username });
      updateLoggedUser({ username });

      if (!token) {
        return;
      }

      await privateApi.updateUsername(token, username);
    });
  }, [username]);

  useEffect(() => {
    if (inputRef.current && isEditing) {
      inputRef.current.focus();
    }
  }, [isEditing]);

  return (
    <div className={cn("flex gap-4", className)} {...props}>
      <Avatar
        className={cn(
          "h-24 w-24",
          "-mt-10",
          "ring-8 ring-(--slate-3) dark:ring-(--slate-2)"
        )}
        src={dsAvatar(user.id, user.avatar_hash)}
      />
      <div className={cn("flex flex-col gap-px", "-mt-2")}>
        <input
          className={cn("text-2xl! font-bold!")}
          disabled={!isEditing}
          value={username}
          onChange={e => setUsername(e.target.value)}
          ref={inputRef}
        />
        <header className={cn("font-semibold", "text-(--gray-10)")}>
          @ {user.discord_username}
        </header>
      </div>
    </div>
  );
};

export { UserInfo };
