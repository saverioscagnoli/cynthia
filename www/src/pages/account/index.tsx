import { Pencil1Icon } from "@radix-ui/react-icons";
import { useHotkey } from "@util-hooks/use-hotkey";
import { Button } from "~/components/ui";
import { cn } from "~/lib/utils";
import { useAuth } from "~/contexts/auth";
import { useProfileEdit } from "~/contexts/profile-edit";
import { UserBanner } from "./banner";
import { UserDisplay } from "./display";
import { TrainerSelectDialog } from "./trainer-dialog";

const AccountPage = () => {
  const { editing, stopEditing, toggleEditing } = useProfileEdit();
  const { user } = useAuth();
  let u = user!;

  useHotkey(window, [], "escape", () => stopEditing());

  return (
    <div className={cn("h-full w-220", "flex flex-col gap-4", "mx-auto my-12")}>
      <div
        className={cn(
          "relative h-100 w-full",
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
            colorScheme={editing ? "red" : "gray"}
            leftIcon={<Pencil1Icon />}
            onClick={toggleEditing}
            bump={false}
          >
            {editing ? "Stop Editing" : "Edit Profile"}
          </Button>
        </div>
      </div>

      <div className={cn("w-full", "flex gap-4")}>
        <div
          className={cn(
            "h-60 w-1/2",
            "bg-(--slate-3) dark:bg-(--slate-2)",
            "border border-(--gray-4)",
            "rounded-2xl"
          )}
        >
          <div
            className={cn(
              "h-full w-full",
              "p-6",
              "flex items-center justify-between"
            )}
          >
            <div className={cn("h-full", "flex flex-col gap-16")}>
              <header className={cn("text-xl font-bold")}>Trainer</header>
              <TrainerSelectDialog>
                <Button variant="soft">Change Sprite</Button>
              </TrainerSelectDialog>
            </div>
            {u.sprite_id !== null ? (
              <img
                width={160}
                height={160}
                style={{ imageRendering: "pixelated" }}
                src={`http://localhost:9000/sprites/trainer/${user.sprite_id}`}
              />
            ) : (
              <div
                className={cn(
                  "h-50 w-50",
                  "inline-flex items-center justify-center",
                  "rounded-2xl",
                  "border border-(--gray-4)"
                )}
              >
                <p>No sprite</p>
              </div>
            )}
          </div>
        </div>
        <div
          className={cn(
            "h-60 w-1/2",
            "bg-(--slate-3) dark:bg-(--slate-2)",
            "border border-(--gray-4)",
            "rounded-2xl"
          )}
        ></div>
      </div>
    </div>
  );
};

export { AccountPage };
