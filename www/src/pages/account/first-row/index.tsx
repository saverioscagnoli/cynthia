import type { ComponentProps } from "react";
import { Card } from "~/components/card";
import { cn } from "~/lib/utils";
import { TrainerStats } from "./stats";
import { TrainerDisplay } from "./trainer-display";

const FirstCardRow: React.FC<ComponentProps<"div">> = ({
  className,
  ...props
}) => {
  return (
    <div className={cn("h-64 w-full", "flex gap-4", className)} {...props}>
      <Card className={cn("h-full w-1/4")}>
        <TrainerDisplay />
      </Card>
      <Card className={cn("h-full w-3/4")}>
        <TrainerStats />
      </Card>
    </div>
  );
};

export { FirstCardRow };
