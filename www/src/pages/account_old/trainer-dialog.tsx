import { Suspense, useRef, useState } from "react";
import { useHotkey } from "@util-hooks/use-hotkey";
import { Button, Input } from "~/components/ui";
import { Dialog } from "~/components/ui/dialog";
import { cn } from "~/lib/utils";
import { privateApi } from "~/lib/wrapper";
import { useAccount } from "~/contexts/account";
import { useAuth } from "~/contexts/auth";
import { TrainerList } from "./trainer-list";

type TrainerSelectDialogProps = {
  children: React.ReactNode;
};

const TrainerSelectDialog: React.FC<TrainerSelectDialogProps> = ({
  children
}) => {
  const { user, updateUser } = useAccount();
  const { token, loggedUser, updateLoggedUser } = useAuth();
  const [search, setSearch] = useState("");
  const [selected, setSelected] = useState<number | null>(null);

  const modalRef = useRef<HTMLDivElement>(null);
  const inputRef = useRef<HTMLInputElement>(null);

  useHotkey(window, ["ctrl"], "f", e => {
    e.preventDefault();
    if (inputRef.current) inputRef.current.focus();
  });

  const onTrainerSpriteChange = async () => {
    if (!token || !selected) return;

    await privateApi.updateTrainerSprite(token, selected);

    updateUser({ sprite_id: selected });

    if (loggedUser?.id === user.id) {
      updateLoggedUser({ sprite_id: selected });
    }

    setSelected(null);
  };

  return (
    <Dialog
      onOpenChange={o => {
        if (!o) {
          setSelected(null);
          setSearch("");
        }
      }}
    >
      <Dialog.Trigger asChild>{children}</Dialog.Trigger>
      <Dialog.Content
        className={cn("h-full w-[75vw]! max-w-none!", "flex flex-col gap-6")}
        ref={modalRef as React.RefObject<HTMLDivElement>}
      >
        <Dialog.Close />
        <Dialog.Title>Select new trainer sprite</Dialog.Title>
        <Input
          size="lg"
          value={search}
          onChange={e => setSearch(e.target.value)}
          ref={inputRef}
        />
        <Suspense>
          <TrainerList
            query={search}
            selected={selected}
            setSelected={setSelected}
          />
        </Suspense>
        <div className={cn("flex justify-end gap-4")}>
          <Button variant="soft" colorScheme="gray">
            Cancel
          </Button>
          <Dialog.Close asChild>
            <Button
              disabled={selected === null}
              variant="soft"
              colorScheme="green"
              onClick={onTrainerSpriteChange}
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
