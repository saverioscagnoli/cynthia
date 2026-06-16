import { Input } from "@mantine/core";
import { useEffect, useState } from "react";

type SpriteEntry = {
    x: number;
    y: number;
    w: number;
    h: number;
    name: string;
};

const SHEET_URL = "http://localhost:9000/sprites/trainer/sheet";
const MAP_URL = "http://localhost:9000/sprites/trainer/sheet/map";
const SCALE = 2;

const TrainersPage = () => {
    const [spriteMap, setSpriteMap] = useState<Record<string, SpriteEntry>>({});
    const [sheetSize, setSheetSize] = useState<{ w: number; h: number } | null>(
        null,
    );
    const [search, setSearch] = useState("");

    useEffect(() => {
        fetch(MAP_URL)
            .then((res) => res.json())
            .then((data) => {
                const { _sheet, ...sprites } = data;
                setSheetSize({ w: _sheet.w, h: _sheet.h });
                setSpriteMap(sprites);
            });
    }, []);

    const names = Object.keys(spriteMap).filter((name) =>
        name.toLowerCase().includes(search.toLowerCase()),
    );

    const loading = !sheetSize || Object.keys(spriteMap).length === 0;

    return (
        <div className="flex flex-col h-[100vh-4rem]">
            {loading ? (
                <div>Loading...</div>
            ) : (
                <div className="flex flex-1 min-h-0">
                    {/* Sidebar */}
                    <div className="w-1/4 shrink-0 p-6 overflow-y-auto">
                        <p className="text-sm font-medium mb-2">
                            Search trainer
                        </p>
                        <Input
                            variant="filled"
                            value={search}
                            onChange={(e) => setSearch(e.currentTarget.value)}
                            rightSection={
                                search ? (
                                    <Input.ClearButton
                                        onClick={() => setSearch("")}
                                    />
                                ) : null
                            }
                        />
                        <p className="text-xs mt-3">{names.length} results</p>
                    </div>
                    {/* Grid */}
                    <div className="flex-1 overflow-y-auto p-6">
                        <div className="grid grid-cols-[repeat(8,160px)] gap-4">
                            {names.map((name) => {
                                const { x, y, w, h } = spriteMap[name];
                                return (
                                    <div
                                        key={name}
                                        title={name}
                                        className="flex flex-col items-center gap-1 cursor-pointer hover:animate-[bounce-sprite_0.4s_ease-in-out_infinite]"
                                    >
                                        <div
                                            style={{
                                                width: w * SCALE,
                                                height: h * SCALE,
                                                backgroundImage: `url(${SHEET_URL})`,
                                                backgroundPosition: `-${x * SCALE}px -${y * SCALE}px`,
                                                backgroundRepeat: "no-repeat",
                                                backgroundSize: `${sheetSize.w * SCALE}px ${sheetSize.h * SCALE}px`,
                                                imageRendering: "pixelated",
                                            }}
                                        />
                                        <p className="text-xs text-center truncate w-full">
                                            {name}
                                        </p>
                                    </div>
                                );
                            })}
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
};

export { TrainersPage };
