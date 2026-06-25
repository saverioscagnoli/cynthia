import { Suspense, useEffect, useRef, useState } from "react";
import { LuCheck, LuX } from "react-icons/lu";
import { useHotkey } from "@util-hooks/use-hotkey";
import { Button, Input, Spinner } from "~/components/ui";
import { Dialog } from "~/components/ui/dialog";
import { cn } from "~/lib/utils";
import { privateApi } from "~/lib/wrapper";
import { useAccount } from "~/contexts/account";
import { useAuth } from "~/contexts/auth";
import { TrainerSelectList } from "./trainer-list";

const TrainerSelectDialog: React.FC<{ children: React.ReactNode }> = ({
  children
}) => {
  const { updateLoggedUser } = useAuth();
  const { onEditConfirm, registerOnEditConfirm, updateUser } = useAccount();
  const [open, setOpen] = useState<boolean>(false);
  const [query, setQuery] = useState<string>("");
  const [selected, setSelected] = useState<number | null>(null);
  const inputRef = useRef<HTMLInputElement>(null);

  useHotkey(window, ["ctrl"], "f", e => {
    if (open) {
      e.preventDefault();
      inputRef.current?.focus();
    }
  });

  const onCancel = () => {
    setQuery("");
    setSelected(null);
  };

  useEffect(() => {
    return registerOnEditConfirm(async () => {
      if (!selected) return;

      await privateApi.updateTrainerSprite(selected);

      updateLoggedUser({ sprite_id: selected });
      updateUser({ sprite_id: selected });

      setQuery("");
      setSelected(null);
    });
  }, [selected]);

  return (
    <Dialog open={open} onOpenChange={o => setOpen(o)}>
      <Dialog.Trigger asChild>{children}</Dialog.Trigger>
      <Dialog.Content
        className={cn("h-[75vh]! w-[65vw]! max-w-none!", "flex flex-col gap-6")}
      >
        <Dialog.Close />
        <Dialog.Title>Change trainer sprite</Dialog.Title>
        <Input
          size="lg"
          value={query}
          onChange={e => setQuery(e.target.value)}
          ref={inputRef}
        />
        <Suspense
          fallback={
            <div
              className={cn(
                "h-full w-full",
                "flex items-center justify-center"
              )}
            >
              <Spinner className={cn("h-16 w-16")} noTrack />
            </div>
          }
        >
          <TrainerSelectList
            query={query}
            selected={selected}
            setSelected={setSelected}
          />
        </Suspense>
        <div className={cn("w-full", "flex justify-end gap-4")}>
          <Dialog.Close asChild>
            <Button
              size="lg"
              variant="soft"
              colorScheme="gray"
              leftIcon={<LuX size={20} />}
              onClick={onCancel}
            >
              Cancel
            </Button>
          </Dialog.Close>
          <Dialog.Close asChild>
            <Button
              disabled={!selected}
              size="lg"
              variant="soft"
              colorScheme="green"
              leftIcon={<LuCheck size={20} />}
              onClick={onEditConfirm}
            >
              Confirm
            </Button>
          </Dialog.Close>
        </div>
      </Dialog.Content>
    </Dialog>
  );
};

export { TrainerSelectDialog };
