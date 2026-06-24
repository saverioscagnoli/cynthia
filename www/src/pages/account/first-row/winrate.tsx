import { cn } from "~/lib/utils";

type WinrateProps = {
  winrate: number;
};

const Winrate: React.FC<WinrateProps> = ({ winrate }) => {
  return (
    <div
      className={cn(
        "flex flex-1 flex-col items-center justify-center gap-1",
        "border-r border-r-(--gray-4)"
      )}
    >
      <header
        className={cn("text-sm tracking-widest text-(--slate-11) uppercase")}
      >
        winrate
      </header>
      <header className={cn("text-4xl font-bold")}>{winrate * 100}%</header>
    </div>
  );
};

export { Winrate };
