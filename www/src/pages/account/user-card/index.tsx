import { LuCheck, LuPencil, LuX } from "react-icons/lu";
import { Card } from "~/components/card";
import { Button } from "~/components/ui";
import { cn } from "~/lib/utils";
import { useAccount } from "~/contexts/account";
import { useAuth } from "~/contexts/auth";
import { Banner } from "./banner";
import { UserInfo } from "./info";

const UserCard = () => {
  const { loggedUser } = useAuth();
  const { user, isEditing, startEdit, stopEdit, onEditConfirm } = useAccount();

  return (
    <Card
      className={cn(
        "h-110 w-full",
        "relative",
        "overflow-hidden",
        "flex flex-col justify-end"
      )}
    >
      <Banner />
      <div
        className={cn(
          "h-2/5 w-full",
          "flex justify-between",
          "px-12 py-4",
          "bg-(--slate-3) dark:bg-(--slate-2)",
          "z-10"
        )}
      >
        <UserInfo />
        {loggedUser?.id === user.id ? (
          isEditing ? (
            <div className={cn("flex gap-4")}>
              <Button
                variant="soft"
                colorScheme="gray"
                leftIcon={<LuX />}
                onClick={stopEdit}
              >
                Cancel
              </Button>
              <Button
                variant="soft"
                colorScheme="green"
                leftIcon={<LuCheck />}
                onClick={onEditConfirm}
              >
                Confirm Edit
              </Button>
            </div>
          ) : (
            <Button
              variant="soft"
              colorScheme="gray"
              leftIcon={<LuPencil />}
              onClick={startEdit}
            >
              Edit Profile
            </Button>
          )
        ) : (
          <span></span>
        )}
      </div>
    </Card>
  );
};

export { UserCard };
