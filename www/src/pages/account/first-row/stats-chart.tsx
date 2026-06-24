import type { MatchStats } from "~/types";
import { Pie, PieChart, Sector } from "recharts";

type StatsChartProps = MatchStats;

const StatsChart: React.FC<StatsChartProps> = stats => {
  const data = [
    { name: "Wins", value: stats.wins, color: "var(--green-9)" },
    { name: "Losses", value: stats.losses, color: "var(--red-9)" },
    { name: "Draws", value: stats.draws, color: "var(--blue-9)" }
  ];

  const total = stats.wins + stats.losses + stats.draws;

  return (
    <PieChart width={180} height={180}>
      <Pie
        data={data}
        dataKey="value"
        nameKey="name"
        cx="50%"
        cy="50%"
        innerRadius={50}
        outerRadius={70}
        shape={(props: any) => {
          const {
            cx,
            cy,
            innerRadius,
            outerRadius,
            startAngle,
            endAngle,
            index
          } = props;
          return (
            <Sector
              cx={cx}
              cy={cy}
              innerRadius={innerRadius}
              outerRadius={outerRadius}
              startAngle={startAngle}
              endAngle={endAngle}
              fill={data[index].color}
            />
          );
        }}
      />
      <text
        x={90}
        y={82}
        textAnchor="middle"
        dominantBaseline="middle"
        fill="var(--slate-12)"
        fontSize={22}
        fontWeight="bold"
      >
        {total}
      </text>
      <text
        x={90}
        y={102}
        textAnchor="middle"
        dominantBaseline="middle"
        fill="var(--slate-11)"
        fontSize={14}
      >
        matches
      </text>
    </PieChart>
  );
};

export { StatsChart };
