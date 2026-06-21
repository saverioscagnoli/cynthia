import { useEffect, useState } from "react";
import { CheckIcon } from "@radix-ui/react-icons";
import { Avatar, Button } from "~/components/ui";
import { updateUsername } from "~/lib/backend";
import { cn, dsAvatar } from "~/lib/utils";
import { useAuth, type User } from "~/contexts/auth";
import { useProfileEdit } from "~/contexts/profile-edit";
import { AccountBadges } from "./badges";

type UserDisplayProps = {
  user: User;
};

const UserDisplay: React.FC<UserDisplayProps> = ({ user }) => {
  const [username, setUsername] = useState(user.username);
  const [error, setError] = useState(false);
  const { token, updateUser } = useAuth();
  const { editing, stopEditing } = useProfileEdit();

  useEffect(() => {
    // Stopped editing
    if (!editing) {
      if (username.length === 0) setUsername(user.username);
    }
  }, [editing]);

  useEffect(() => {
    if (error) {
      setTimeout(() => setError(false), 3000);
    }
  }, [error]);

  const onUsernameChange = async () => {
    let err = false;

    try {
      await updateUsername(token, username);
    } catch (e) {
      console.error(e);
      setError(true);
      err = true;
    }

    if (err) {
      setUsername(user.username);
    } else {
      stopEditing();
      updateUser({ username });
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
          {editing ? (
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
