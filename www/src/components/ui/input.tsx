import React, { type ComponentProps } from "react";
import { cva, type VariantProps } from "class-variance-authority";
import "@radix-ui/colors/slate.css";
import "@radix-ui/colors/slate-dark.css";
import "@radix-ui/colors/plum.css";
import "@radix-ui/colors/plum-dark.css";
import "@radix-ui/colors/red.css";
import "@radix-ui/colors/red-dark.css";
import "@radix-ui/colors/green.css";
import "@radix-ui/colors/green-dark.css";
import "@radix-ui/colors/blue.css";
import "@radix-ui/colors/blue-dark.css";
import "@radix-ui/colors/teal.css";
import "@radix-ui/colors/teal-dark.css";
import "@radix-ui/colors/cyan.css";
import "@radix-ui/colors/cyan-dark.css";
import "@radix-ui/colors/yellow.css";
import "@radix-ui/colors/yellow-dark.css";
import "@radix-ui/colors/orange.css";
import "@radix-ui/colors/orange-dark.css";
import "@radix-ui/colors/gray.css";
import "@radix-ui/colors/gray-dark.css";

const inputVariants = cva(
  [
    "bg-transparent border border-(--gray-7) outline-none",
    "caret-(--slate-12) text-(--slate-12)",
    "font-normal shadow transition-colors duration-100",
    "focus:ring-2 focus:ring-offset-2 focus:ring-offset-(--slate-1)",
    "disabled:cursor-not-allowed disabled:opacity-50"
  ],
  {
    variants: {
      colorScheme: {
        plum: "focus:ring-(--plum-9)",
        red: "focus:ring-(--red-9)",
        blue: "focus:ring-(--blue-9)",
        teal: "focus:ring-(--teal-9)",
        cyan: "focus:ring-(--cyan-9)",
        green: "focus:ring-(--green-9)",
        amber: "focus:ring-(--amber-9)",
        orange: "focus:ring-(--orange-9)",
        gray: "focus:ring-(--slate-12)"
      },
      size: {
        sm: "h-7 px-2 text-sm rounded",
        md: "h-9 px-3 text-sm rounded",
        lg: "h-12 px-3  text-lg rounded-md"
      },
      invalid: {
        true: "border-(--red-9)"
      }
    },
    defaultVariants: {
      colorScheme: "plum",
      size: "md"
    }
  }
);

type InputProps = Omit<ComponentProps<"input">, "size"> &
  VariantProps<typeof inputVariants> & {
    htmlSize?: ComponentProps<"input">["size"];
  };

const Input: React.FC<InputProps> = ({
  colorScheme,
  size,
  invalid,
  htmlSize,
  className,
  ...props
}) => {
  return (
    <input
      className={inputVariants({ colorScheme, size, invalid, className })}
      size={htmlSize}
      {...props}
    />
  );
};

Input.displayName = "Input";

export { Input };
