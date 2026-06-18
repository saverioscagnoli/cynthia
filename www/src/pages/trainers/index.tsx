import { Input, Loader, Modal } from "@mantine/core";
import { Suspense, useRef, useState } from "react";
import { TrainerList } from "./trainer-list";
import { cn } from "~/lib/utils";
import { useHotkey } from "@util-hooks/use-hotkey";
import { useDisclosure } from "@mantine/hooks";
import type { TrainerSpriteMapEntry } from "~/types";
import { useAuth } from "~/contexts/auth";

const SHEET_URL = "http://localhost:9000/sprites/trainer/sheet";
const SCALE = 2;

const TrainersPage = () => {
  const [query, setQuery] = useState<string>("");
  const inputRef = useRef<HTMLInputElement>(null);
  const [opened, { open, close }] = useDisclosure();
  const [title, setTitle] = useState<string>("Confirm sprite change");
  const { user, trainer } = useAuth();
  const [[_s, e], setSelected] = useState<
    [TrainerSpriteMapEntry, TrainerSpriteMapEntry]
  >([null, null]);

  useHotkey(window, ["ctrl"], "f", (e) => {
    e.preventDefault();

    if (inputRef.current) {
      inputRef.current.focus();
    }
  });

  return (
    <div className={cn("relative h-[calc(100vh-4rem)]")}>
      <Suspense
        fallback={
          <Loader
            className={cn("absolute top-1/2 left-1/2 -translate-x-1/2")}
            color="grape"
          />
        }
      >
        <TrainerList
          query={query}
          onClick={(_s, t) => {
            open();
            setSelected([_s, t]);

            if (user) {
              setTitle("Sprite change confirmation");
            } else {
              setTitle("Authentication Error");
            }
          }}
        />
      </Suspense>
      <div className={cn("absolute top-4 left-1/2 -translate-x-1/2")}>
        <Input
          className={cn("w-96")}
          placeholder="Search trainers..."
          value={query}
          onChange={(e) => setQuery(e.currentTarget.value)}
          ref={inputRef}
        />
      </div>
      <Modal
        title={title}
        centered
        opened={opened}
        size="auto"
        onClose={() => {
          setSelected([null, null]);
          close();
        }}
      >
        {_s && user && trainer ? (
          <div className={cn("w-full h-full", "p-4")}>
            <img src={`http://localhost:9000/sprites/trainer/${trainer.id}`} />
            to
            <div
              style={{
                width: e.w * SCALE,
                height: e.h * SCALE,
                backgroundImage: `url(${SHEET_URL})`,
                backgroundPosition: `-${e.x * SCALE}px -${e.y * SCALE}px`,
                backgroundRepeat: "no-repeat",
                backgroundSize: `${_s.w * SCALE}px ${_s.h * SCALE}px`,
                imageRendering: "pixelated",
              }}
            />
          </div>
        ) : (
          <p>You need to log in to change the sprite!</p>
        )}
      </Modal>
    </div>
  );
};

export { TrainersPage };
