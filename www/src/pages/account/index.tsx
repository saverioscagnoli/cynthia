import { Pencil1Icon } from "@radix-ui/react-icons";
import { Avatar, Button } from "~/components/ui";
import { cn, dsAvatar } from "~/lib/utils";
import { useAuth } from "~/contexts/auth";
import { useTheme } from "~/contexts/theme";
import { UserBanner } from "./banner";
import { UserDisplay } from "./display";

const AccountPage = () => {
  const { theme } = useTheme();
  const { user } = useAuth();
  let u = user!;

  console.log(user);

  return (
    <div
      className={cn(
        "relative h-100 w-220",
        "mx-auto my-12",
        "border border-(--gray-4)",
        "rounded-2xl",
        "overflow-hidden"
      )}
    >
      <UserBanner />
      <div
        className={cn(
          "absolute right-0 bottom-0 left-0",
          "h-2/5",
          "flex justify-between",
          "px-8 pt-4",
          "bg-(--slate-3) dark:bg-(--slate-2)",
          "rounded-b-2xl"
        )}
      >
        <UserDisplay user={u} />
        <Button
          className={cn("mt-4")}
          size="sm"
          variant="soft"
          colorScheme="gray"
          leftIcon={<Pencil1Icon />}
        >
          Edit Profile
        </Button>
      </div>
    </div>
  );
};

export { AccountPage };
