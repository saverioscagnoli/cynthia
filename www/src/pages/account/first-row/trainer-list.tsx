import { use, type Dispatch, type SetStateAction } from "react";
import { VirtuosoGrid as VGrid } from "react-virtuoso";
import type { TrainerSpriteMapEntry, TrainerSpriteSheet } from "~/types";
import { Checkbox } from "~/components/ui/checkbox";
import { cn } from "~/lib/utils";

type TrainerSelectListProps = {
  query: string;
  selected: number | null;
  setSelected: Dispatch<SetStateAction<number | null>>;
};

const map = fetch("http://localhost:9000/sprites/trainer/sheet/map").then(r =>
  r.json()
);

const TrainerSelectList: React.FC<TrainerSelectListProps> = ({
  query,
  selected,
  setSelected
}) => {
  const { _sheet, sprites } = use<TrainerSpriteSheet>(map);
  const list = Object.values(sprites);
  const filtered = list.filter(e => e.name.includes(query));

  const onSpriteClick = (e: TrainerSpriteMapEntry) => () => {
    if (selected === e.id) setSelected(null);
    else setSelected(e.id);
  };

  return (
    <VGrid
      className={cn("h-full")}
      data={filtered}
      listClassName={cn("flex flex-wrap")}
      itemContent={(_, e) => (
        <div
          className={cn(
            "flex flex-col items-center justify-between gap-2",
            "relative",
            "p-2"
          )}
          onClick={onSpriteClick(e)}
        >
          <div
            className={cn(
              "cursor-poiner hover:animate-[bounce-sprite_0.4s_ease-in-out_infinite]"
            )}
            style={{
              width: e.w * 2,
              height: e.h * 2,
              backgroundImage:
                "url(http://localhost:9000/sprites/trainer/sheet)",
              backgroundPosition: `-${e.x * 2}px -${e.y * 2}px`,
              backgroundRepeat: "no-repeat",
              backgroundSize: `${_sheet.w * 2}px ${_sheet.h * 2}px`,
              imageRendering: "pixelated"
            }}
          />
          <p>{e.name}</p>
          {selected === e.id && (
            <Checkbox className={cn("absolute top-4 right-4")} checked />
          )}
        </div>
      )}
    />
  );
};

export { TrainerSelectList };
