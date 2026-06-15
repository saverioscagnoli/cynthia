import { Input } from "@mantine/core";
import { useEffect, useState } from "react";
import { LazyImage } from "~/components/lazy-image";
import { Topbar } from "~/components/topbar";

const TrainersPage = () => {
    const [trainerSpriteCount, setTrainerSpriteCount] = useState<number>(0);
    const [search, setSearch] = useState("");

    useEffect(() => {
        fetch("http://localhost:9000/sprites/trainer/count")
            .then((res) => res.text())
            .then((count) => setTrainerSpriteCount(parseInt(count)));
    }, []);

    const indices = Array.from(
        { length: trainerSpriteCount },
        (_, i) => i,
    ).filter((i) => i.toString().includes(search));

    return (
        <div className="flex flex-col h-screen">
            {trainerSpriteCount == 0 ? (
                <div>Loading...</div>
            ) : (
                <div className="flex flex-1 min-h-0">
                    {/* min-h-0 prevents flex overflow */}
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
                        <p className="text-xs mt-3">{indices.length} results</p>
                    </div>
                    {/* Grid */}
                    <div className="flex-1 overflow-y-auto p-6">
                        <div className="grid grid-cols-[repeat(8,160px)] gap-4">
                            {indices.map((i) => (
                                <LazyImage
                                    key={i}
                                    width={160}
                                    height={160}
                                    className="[image-rendering:pixelated] hover:animate-[bounce-sprite_0.4s_ease-in-out_infinite] cursor-pointer"
                                    src={`http://localhost:9000/sprites/trainer/${i}`}
                                />
                            ))}
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
};

export { TrainersPage };
