import type { ComponentProps } from "react";
import { cn } from "~/lib/utils";

type CardProps = ComponentProps<"div">;

const Card: React.FC<CardProps> = ({ className, ...props }) => {
  return (
    <div
      className={cn(
        "border border-(--gray-4)",
        "rounded-2xl",
        "bg-(--slate-3) dark:bg-(--slate-2)",
        className
      )}
      {...props}
    />
  );
};

export { Card };
