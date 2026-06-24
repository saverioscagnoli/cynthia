import { LuHouse, LuLogOut, LuUser } from "react-icons/lu";
import { useNavigate } from "react-router";
import { DropdownMenu } from "~/components/ui";
import { cn } from "~/lib/utils";
import { useAuth } from "~/contexts/auth";

type ProfileMenuProps = {
  children: React.ReactNode;
};

const ProfileMenu: React.FC<ProfileMenuProps> = ({ children }) => {
  const { loggedUser, logout } = useAuth();
  const user = loggedUser!;
  const nav = useNavigate();

  const onLogout = () => {
    nav("/");
    logout();
  };

  return (
    <DropdownMenu colorScheme="gray">
      <DropdownMenu.Trigger>{children}</DropdownMenu.Trigger>
      <DropdownMenu.Content
        className={cn("w-44", "mt-2", "z-9999")}
        onCloseAutoFocus={e => e.preventDefault()}
      >
        <DropdownMenu.Item leftIcon={<LuHouse />} onClick={() => nav("/")}>
          Home
        </DropdownMenu.Item>
        <DropdownMenu.Separator />
        <DropdownMenu.Label>Profile</DropdownMenu.Label>
        <DropdownMenu.Item
          leftIcon={<LuUser />}
          onClick={() => nav(`/user/${user.id}`)}
        >
          Account
        </DropdownMenu.Item>
        <DropdownMenu.Item
          colorScheme="red"
          leftIcon={<LuLogOut />}
          onClick={onLogout}
        >
          Logout
        </DropdownMenu.Item>
      </DropdownMenu.Content>
    </DropdownMenu>
  );
};

export { ProfileMenu };
