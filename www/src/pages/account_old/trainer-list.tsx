import { use, type Dispatch } from "react";
import { VirtuosoGrid as VGrid } from "react-virtuoso";
import type { TrainerSpriteSheet } from "~/types";
import { Checkbox } from "~/components/ui/checkbox";
import { cn } from "~/lib/utils";

const SHEET_URL = "http://localhost:9000/sprites/trainer/sheet";
const MAP_URL = "http://localhost:9000/sprites/trainer/sheet/map";
const SCALE = 2;

const mapPromise = fetch(MAP_URL).then(res => res.json());

type TrainerListProps = {
  query: string;
  selected: number | null;
  setSelected: Dispatch<number | null>;
};

const TrainerList: React.FC<TrainerListProps> = ({
  query,
  selected,
  setSelected
}) => {
  const { _sheet, sprites } = use<TrainerSpriteSheet>(mapPromise);

  return (
    <VGrid
      style={{ height: "100%" }}
      data={Object.values(sprites).filter(e => e.name.includes(query))}
      listClassName={cn("flex flex-wrap")}
      itemContent={(_, e) => (
        <div
          className={cn(
            "flex flex-col items-center justify-center gap-2",
            "relative",
            "p-3"
          )}
          onClick={() => {
            if (selected === e.id) setSelected(null);
            else setSelected(e.id);
          }}
        >
          <div
            className={cn(
              "cursor-pointer hover:animate-[bounce-sprite_0.4s_ease-in-out_infinite]"
            )}
            style={{
              width: e.w * SCALE,
              height: e.h * SCALE,
              backgroundImage: `url(${SHEET_URL})`,
              backgroundPosition: `-${e.x * SCALE}px -${e.y * SCALE}px`,
              backgroundRepeat: "no-repeat",
              backgroundSize: `${_sheet.w * SCALE}px ${_sheet.h * SCALE}px`,
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

export { TrainerList };
