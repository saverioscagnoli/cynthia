import { use } from "react";
import { cn } from "~/lib/utils";
import type { TrainerSpriteMapEntry, TrainerSpriteSheetMap } from "~/types";
import { VirtuosoGrid as VGrid } from "react-virtuoso";

const SHEET_URL = "http://localhost:9000/sprites/trainer/sheet";
const MAP_URL = "http://localhost:9000/sprites/trainer/sheet/map";
const SCALE = 2;

const mapPromise = fetch(MAP_URL).then((res) => res.json());

type TrainerListProps = {
  query: string;
  onClick?: (_s: TrainerSpriteMapEntry, t: TrainerSpriteMapEntry) => void;
};

const TrainerList: React.FC<TrainerListProps> = ({ query, onClick }) => {
  const { _sheet, ...sprites } = use<TrainerSpriteSheetMap>(mapPromise);

  return (
    <VGrid
      style={{ height: "100%" }}
      data={Object.values(sprites).filter((e) => e.name.includes(query))}
      listClassName={cn("flex flex-wrap")}
      components={{
        Header: () => <div className={cn("h-16")} />,
      }}
      itemContent={(_, e) => (
        <div
          className={cn("flex flex-col items-center justify-center gap-2 p-3")}
        >
          <div
            className={cn(
              "cursor-pointer hover:animate-[bounce-sprite_0.4s_ease-in-out_infinite]",
            )}
            style={{
              width: e.w * SCALE,
              height: e.h * SCALE,
              backgroundImage: `url(${SHEET_URL})`,
              backgroundPosition: `-${e.x * SCALE}px -${e.y * SCALE}px`,
              backgroundRepeat: "no-repeat",
              backgroundSize: `${_sheet.w * SCALE}px ${_sheet.h * SCALE}px`,
              imageRendering: "pixelated",
            }}
            onClick={() => onClick(_sheet, e)}
          />
          <p>{e.name}</p>
        </div>
      )}
    />
  );
};

export { TrainerList };
