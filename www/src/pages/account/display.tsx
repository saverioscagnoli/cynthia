import { Avatar } from "~/components/ui";
import { cn, dsAvatar } from "~/lib/utils";
import type { User } from "~/contexts/auth";
import { AccountBadges } from "./badges";

type UserDisplayProps = {
  user: User;
};

const UserDisplay: React.FC<UserDisplayProps> = ({ user }) => {
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
          <header className={cn("text-2xl font-bold")}>{user.username}</header>
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
