import { cn } from "~/lib/utils";
import { useAccount } from "~/contexts/account";
import { TrainerSelectDialog } from "./trainer-dialog";

const TrainerDisplay = () => {
  const { user, isEditing } = useAccount();

  return (
    <TrainerSelectDialog>
      <div
        className={cn(
          "h-full w-full",
          "flex items-center justify-center",
          "relative",
          {
            "pointer-events-none": !isEditing
          }
        )}
      >
        <header className={cn("absolute top-4 left-4", "text-xl font-bold")}>
          Trainer
        </header>
        {user.sprite_id ? (
          <img
            className={cn({
              "animate-[bounce-sprite_0.4s_ease-in-out_infinite] cursor-pointer":
                isEditing
            })}
            src={`http://localhost:9000/sprites/trainer/${user.sprite_id}`}
            width={160}
            height={160}
            style={{ imageRendering: "pixelated" }}
          />
        ) : (
          <header>No sprite</header>
        )}
      </div>
    </TrainerSelectDialog>
  );
};

export { TrainerDisplay };
