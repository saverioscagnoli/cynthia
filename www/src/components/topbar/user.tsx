import { useNavigate } from "react-router";
import { Avatar, Button } from "~/components/ui";
import { cn, dsAvatar } from "~/lib/utils";
import { useAuth } from "~/contexts/auth";
import { ProfileMenu } from "./profile-menu";

const UserDisplay = () => {
  const { loggedUser, logout } = useAuth();
  const user = loggedUser!;
  const nav = useNavigate();

  const onLogout = () => {
    logout();
    nav("/");
  };

  return (
    <div className={cn("w-fit", "flex items-center gap-4")}>
      <ProfileMenu>
        <div className={cn("flex items-center gap-4", "cursor-pointer")}>
          <Avatar size="sm" src={dsAvatar(user.id, user.avatar_hash)} />
          <p>{user.username}</p>
        </div>
      </ProfileMenu>
      <Button variant="soft" colorScheme="red" onClick={onLogout}>
        Logout
      </Button>
    </div>
  );
};

export { UserDisplay };
