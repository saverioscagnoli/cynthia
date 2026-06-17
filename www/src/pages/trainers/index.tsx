import { Input, Loader } from "@mantine/core";
import { Suspense, useRef, useState } from "react";
import { TrainerList } from "./trainer-list";
import { cn } from "~/lib/utils";
import { useHotkey } from "@util-hooks/use-hotkey";

const TrainersPage = () => {
    const [query, setQuery] = useState<string>("");
    const inputRef = useRef<HTMLInputElement>(null);

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
                        className={cn(
                            "absolute top-1/2 left-1/2 -translate-x-1/2",
                        )}
                        color="grape"
                    />
                }
            >
                <TrainerList query={query} />
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
        </div>
    );
};

export { TrainersPage };
