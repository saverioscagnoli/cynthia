import { useEffect, useState } from "react";
import { CheckIcon } from "@radix-ui/react-icons";
import { Avatar, Button } from "~/components/ui";
import { cn, dsAvatar } from "~/lib/utils";
import { privateApi } from "~/lib/wrapper";
import { useAccount } from "~/contexts/account";
import { useAuth } from "~/contexts/auth";
import { AccountBadges } from "./badges";

const UserDisplay = () => {
  const { user, updateUser, isEditing, stopEdit } = useAccount();
  const [username, setUsername] = useState(user.username);
  const [error, setError] = useState(false);
  const { token, loggedUser, updateLoggedUser } = useAuth();

  useEffect(() => {
    // Stopped editing
    if (!isEditing && user.username) {
      if (username.length === 0) setUsername(user.username);
    }
  }, [isEditing]);

  useEffect(() => {
    if (error) {
      setTimeout(() => setError(false), 3000);
    }
  }, [error]);

  const onUsernameChange = async () => {
    if (!token) return;

    let res = await privateApi.updateUsername(token, username);

    if (!res.ok) {
      console.error(res.error.message);

      setError(true);
      setUsername(user.username);
      return;
    }

    stopEdit();
    updateUser({ username });

    if (loggedUser?.id === user.id) {
      updateLoggedUser({ username });
    }
  };

  return (
    <div className={cn("flex flex-col gap-4")}>
      <div className={cn("flex gap-4")}>
        <Avatar
          className={cn(
            "h-24 w-24",
            "-mt-10",
            "ring-8 ring-(--slate-3) dark:ring-(--slate-2)",
            "z-10"
          )}
          src={dsAvatar(user.id, user.avatar_hash)}
        />
        <div className={cn("mt-2 flex flex-col gap-px")}>
          {isEditing ? (
            <div className={cn("flex items-center gap-2")}>
              <input
                className={cn("w-50! text-2xl! font-bold!")}
                value={username}
                onChange={e => setUsername(e.target.value)}
              />
              <Button
                size="icon"
                variant="soft"
                colorScheme="green"
                onClick={onUsernameChange}
              >
                <CheckIcon />
              </Button>
              {error && (
                <p className={cn("text-sm text-red-500")}>
                  Failed to update username
                </p>
              )}
            </div>
          ) : (
            <header className={cn("text-2xl font-bold")}>{username}</header>
          )}
          <header className={cn("text-(--gray-10)")}>
            @ {user.discord_username}
          </header>
        </div>
      </div>
      <AccountBadges user={user} />
    </div>
  );
};

export { UserDisplay };
