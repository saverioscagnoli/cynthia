import { cn } from "~/lib/utils";

type StreaksProps = {
  current: number;
  best: number;
};

const Streaks: React.FC<StreaksProps> = ({ current, best }) => {
  return (
    <div
      className={cn("flex flex-1 flex-col items-center justify-center gap-3")}
    >
      <header
        className={cn("text-sm tracking-widest text-(--slate-11) uppercase")}
      >
        streaks
      </header>
      <div className={cn("flex gap-6")}>
        <div className={cn("flex flex-col items-center gap-0.5")}>
          <header className={cn("text-2xl font-semibold")}>{current}</header>
          <header className={cn("text-sm text-(--slate-11)")}>Current</header>
        </div>
        <div className={cn("flex flex-col items-center gap-0.5")}>
          <header className={cn("text-2xl font-semibold")}>{best}</header>
          <header className={cn("text-sm text-(--slate-11)")}>Best</header>
        </div>
      </div>
    </div>
  );
};

export { Streaks };
