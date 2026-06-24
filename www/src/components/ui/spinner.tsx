import type { ComponentProps } from "react";
import { cva, type VariantProps } from "class-variance-authority";
import "@radix-ui/colors/slate.css";
import "@radix-ui/colors/slate-dark.css";
import "@radix-ui/colors/plum.css";
import "@radix-ui/colors/plum-dark.css";
import "@radix-ui/colors/red.css";
import "@radix-ui/colors/red-dark.css";
import "@radix-ui/colors/blue.css";
import "@radix-ui/colors/blue-dark.css";
import "@radix-ui/colors/teal.css";
import "@radix-ui/colors/teal-dark.css";
import "@radix-ui/colors/cyan.css";
import "@radix-ui/colors/cyan-dark.css";
import "@radix-ui/colors/green.css";
import "@radix-ui/colors/green-dark.css";
import "@radix-ui/colors/amber.css";
import "@radix-ui/colors/amber-dark.css";
import "@radix-ui/colors/orange.css";
import "@radix-ui/colors/orange-dark.css";

const spinnerVariants = cva("animate-spin", {
  variants: {
    colorScheme: {
      plum: "[&_.spinner-arc]:stroke-(--plum-9)",
      red: "[&_.spinner-arc]:stroke-(--red-9)",
      blue: "[&_.spinner-arc]:stroke-(--blue-9)",
      teal: "[&_.spinner-arc]:stroke-(--teal-9)",
      cyan: "[&_.spinner-arc]:stroke-(--cyan-9)",
      green: "[&_.spinner-arc]:stroke-(--green-9)",
      amber: "[&_.spinner-arc]:stroke-(--amber-9)",
      orange: "[&_.spinner-arc]:stroke-(--orange-9)",
      gray: "[&_.spinner-arc]:stroke-(--slate-12)",
      white:
        "[&_.spinner-arc]:stroke-(--slate-1) [&_.spinner-arc]:dark:stroke-(--slate-12)",
      black:
        "[&_.spinner-arc]:stroke-(--slate-12) [&_.spinner-arc]:dark:stroke-(--slate-1)"
    },
    size: {
      sm: "w-3 h-3",
      md: "w-5 h-5",
      lg: "w-7 h-7"
    }
  }
});

type SpinnerProps = ComponentProps<"svg"> &
  VariantProps<typeof spinnerVariants> & {
    noTrack?: boolean;
  };

export type { SpinnerProps };

const Spinner: React.FC<SpinnerProps> = ({
  colorScheme = "plum",
  size = "md",
  noTrack = false,
  className,
  ...props
}) => {
  return (
    <svg
      viewBox="0 0 24 24"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      className={spinnerVariants({
        className,
        colorScheme,
        size: typeof size !== "number" ? size : null
      })}
      {...props}
    >
      {!noTrack && (
        <circle
          cx="12"
          cy="12"
          r="10"
          stroke="#BDBDBD"
          strokeOpacity="0.25"
          strokeWidth="3"
          fill="none"
        />
      )}

      <circle
        className="spinner-arc"
        cx="12"
        cy="12"
        r="10"
        fill="none"
        strokeWidth="3"
        strokeLinecap="round"
        strokeDasharray="31.416 31.416" // Circumference is 2π * 10 ≈ 62.832, so half is ~31.416
        strokeDashoffset="15.708" // Quarter turn offset
      />
    </svg>
  );
};

export { Spinner };
