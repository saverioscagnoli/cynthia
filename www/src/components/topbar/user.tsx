import { Avatar, Button } from "~/components/ui";
import { cn, dsAvatar } from "~/lib/utils";
import { useAuth } from "~/contexts/auth";
import { ProfileMenu } from "./profile-menu";

const UserDisplay = () => {
  const { user, logout } = useAuth();
  let u = user!;

  return (
    <div className={cn("w-fit", "flex items-center gap-4")}>
      <ProfileMenu>
        <div className={cn("flex items-center gap-4", "cursor-pointer")}>
          <Avatar size="sm" src={dsAvatar(u.id, u.avatar_hash)} />
          <p>{u.username}</p>
        </div>
      </ProfileMenu>
      <Button variant="soft" colorScheme="red" onClick={logout}>
        Logout
      </Button>
    </div>
  );
};

export { UserDisplay };
