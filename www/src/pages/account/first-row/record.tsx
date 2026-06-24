import { cn } from "~/lib/utils";

type StatsRecordProps = {
  wins: number;
  losses: number;
  draws: number;
};

const StatsRecord: React.FC<StatsRecordProps> = ({ wins, losses, draws }) => {
  return (
    <div
      className={cn(
        "flex flex-1 flex-col items-center justify-center gap-3",
        "border-r border-r-(--gray-4) pr-4"
      )}
    >
      <header
        className={cn("text-sm tracking-widest text-(--slate-11) uppercase")}
      >
        Record
      </header>
      <div className={cn("flex gap-6")}>
        <div className={cn("flex flex-col items-center gap-0.5")}>
          <header className={cn("text-2xl font-semibold text-(--green-9)")}>
            {wins}
          </header>
          <header className={cn("text-sm text-(--slate-11)")}>Wins</header>
        </div>
        <div className={cn("flex flex-col items-center gap-0.5")}>
          <header className={cn("text-2xl font-semibold text-(--red-9)")}>
            {losses}
          </header>
          <header className={cn("text-sm text-(--slate-11)")}>Losses</header>
        </div>
        <div className={cn("flex flex-col items-center gap-0.5")}>
          <header className={cn("text-2xl font-semibold text-(--blue-9)")}>
            {draws}
          </header>
          <header className={cn("text-sm text-(--slate-11)")}>Draws</header>
        </div>
      </div>
    </div>
  );
};
export { StatsRecord };
