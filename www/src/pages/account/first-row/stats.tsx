import { useEffect, useState } from "react";
import type { MatchStats } from "~/types";
import { Spinner } from "~/components/ui";
import { cn } from "~/lib/utils";
import { publicApi } from "~/lib/wrapper";
import { useAccount } from "~/contexts/account";
import { StatsRecord } from "./record";
import { StatsChart } from "./stats-chart";
import { Streaks } from "./streaks";
import { Winrate } from "./winrate";

const TrainerStats = () => {
  const { user } = useAccount();
  const [stats, setStats] = useState<MatchStats | null>(null);

  const fetchStats = async () => {
    if (!user.id) return;

    let res = await publicApi.getUserStats(user.id);

    if (!res.ok) {
      console.error(res.error);
      return;
    }

    setStats(res.data);
  };

  useEffect(() => {
    fetchStats();
  }, [user]);

  if (!stats) {
    return (
      <div className={cn("h-full w-full", "flex items-center justify-center")}>
        <Spinner size="lg" noTrack />
      </div>
    );
  }

  return (
    <div className={cn("h-full w-full", "flex items-center gap-8", "px-6")}>
      <StatsChart {...stats} />

      <div className="flex flex-1 gap-6">
        <Winrate winrate={stats.winrate} />
        <StatsRecord
          wins={stats.wins}
          losses={stats.losses}
          draws={stats.draws}
        />

        <Streaks current={stats.current_streak} best={stats.best_streak} />
      </div>
    </div>
  );
};

export { TrainerStats };
