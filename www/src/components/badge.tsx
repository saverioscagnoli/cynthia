import type React from "react";
import type { ComponentProps } from "react";
import { cn } from "~/lib/utils";

type BadgeProps = ComponentProps<"div"> & {
  label: string;
  leftIcon?: React.ReactNode;
  rightIcon?: React.ReactNode;
};

const Badge: React.FC<BadgeProps> = ({
  label,
  leftIcon,
  rightIcon,
  className,
  ...props
}) => {
  return (
    <div
      className={cn(
        "inline-flex items-center gap-2",
        "px-3 py-1",
        "bg-(--gray-5)",
        "text-xs text-(--gray-12)",
        "rounded-xl",
        className
      )}
      {...props}
    >
      {leftIcon}
      {label}
      {rightIcon}
    </div>
  );
};

export { Badge };
