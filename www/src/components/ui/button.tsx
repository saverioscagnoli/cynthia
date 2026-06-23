import type React from "react";
import type { ComponentProps, ReactNode } from "react";
import { cva } from "class-variance-authority";
import type { VariantProps } from "class-variance-authority";
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
import "@radix-ui/colors/gray.css";
import "@radix-ui/colors/gray-dark.css";

const buttonVariants = cva(
  [
    "inline-flex items-center justify-center gap-2",
    "select-none",
    "cursor-pointer",
    "font-semibold",
    "transition-all",
    "focus-visible:outline-none",
    "focus-visible:ring-2",
    "focus-visible:ring-(--slate-12)",
    "focus-visible:ring-offset-2",
    "dark:focus-visible:[--tw-ring-offset-color:var(--slate-1)]",
    "disabled:opacity-50",
    "disabled:pointer-events-none",
    "disabled:cursor-not-allowed"
  ],
  {
    variants: {
      size: {
        sm: "h-7 px-2.5 py-1.5 rounded text-sm",
        md: "h-9 px-4 py-2 rounded-md text-sm",
        lg: "h-12 px-5 py-3 rounded-md",
        icon: "w-9 h-9 rounded-md"
      },
      colorScheme: {
        plum: "",
        red: "",
        blue: "",
        teal: "",
        cyan: "",
        green: "",
        amber: "",
        orange: "",
        gray: ""
      },
      variant: {
        solid: "",
        soft: "",
        outline: "",
        ghost: ""
      },
      bump: {
        true: "active:translate-y-px duration-100"
      }
    },
    compoundVariants: [
      // Plum
      {
        colorScheme: "plum",
        variant: "solid",
        class: [
          "bg-(--plum-9) hover:bg-(--plum-10)",
          "text-(--slate-1) dark:text-(--slate-12)"
        ]
      },
      {
        colorScheme: "plum",
        variant: "soft",
        class: ["bg-(--plum-5) hover:bg-(--plum-6)", "text-(--plum-11)"]
      },
      {
        colorScheme: "plum",
        variant: "outline",
        class: [
          "outline outline-1 outline-(--plum-7)",
          "bg-transparent hover:bg-(--plum-6)",
          "text-(--plum-11)"
        ]
      },
      {
        colorScheme: "plum",
        variant: "ghost",
        class: ["bg-transparent hover:bg-(--plum-6)", "text-(--plum-11)"]
      },

      // Red
      {
        colorScheme: "red",
        variant: "solid",
        class: [
          "bg-(--red-9) hover:bg-(--red-10)",
          "text-(--slate-1) dark:text-(--slate-12)"
        ]
      },
      {
        colorScheme: "red",
        variant: "soft",
        class: ["bg-(--red-5) hover:bg-(--red-6)", "text-(--red-11)"]
      },
      {
        colorScheme: "red",
        variant: "outline",
        class: [
          "outline outline-1 outline-(--red-7)",
          "bg-transparent hover:bg-(--red-6)",
          "text-(--red-11)"
        ]
      },
      {
        colorScheme: "red",
        variant: "ghost",
        class: ["bg-transparent hover:bg-(--red-6)", "text-(--red-11)"]
      },

      // Blue
      {
        colorScheme: "blue",
        variant: "solid",
        class: [
          "bg-(--blue-9) hover:bg-(--blue-10)",
          "text-(--slate-1) dark:text-(--slate-12)"
        ]
      },
      {
        colorScheme: "blue",
        variant: "soft",
        class: ["bg-(--blue-5) hover:bg-(--blue-6)", "text-(--blue-11)"]
      },
      {
        colorScheme: "blue",
        variant: "outline",
        class: [
          "outline outline-1 outline-(--blue-7)",
          "bg-transparent hover:bg-(--blue-6)",
          "text-(--blue-11)"
        ]
      },
      {
        colorScheme: "blue",
        variant: "ghost",
        class: ["bg-transparent hover:bg-(--blue-6)", "text-(--blue-11)"]
      },

      // Teal
      {
        colorScheme: "teal",
        variant: "solid",
        class: [
          "bg-(--teal-9) hover:bg-(--teal-10)",
          "text-(--slate-1) dark:text-(--slate-12)"
        ]
      },
      {
        colorScheme: "teal",
        variant: "soft",
        class: ["bg-(--teal-5) hover:bg-(--teal-6)", "text-(--teal-11)"]
      },
      {
        colorScheme: "teal",
        variant: "outline",
        class: [
          "outline outline-1 outline-(--teal-7)",
          "bg-transparent hover:bg-(--teal-6)",
          "text-(--teal-11)"
        ]
      },
      {
        colorScheme: "teal",
        variant: "ghost",
        class: ["bg-transparent hover:bg-(--teal-6)", "text-(--teal-11)"]
      },

      // Cyan
      {
        colorScheme: "cyan",
        variant: "solid",
        class: [
          "bg-(--cyan-9) hover:bg-(--cyan-10)",
          "text-(--slate-1) dark:text-(--slate-12)"
        ]
      },
      {
        colorScheme: "cyan",
        variant: "soft",
        class: ["bg-(--cyan-5) hover:bg-(--cyan-6)", "text-(--cyan-11)"]
      },
      {
        colorScheme: "cyan",
        variant: "outline",
        class: [
          "outline outline-1 outline-(--cyan-7)",
          "bg-transparent hover:bg-(--cyan-6)",
          "text-(--cyan-11)"
        ]
      },
      {
        colorScheme: "cyan",
        variant: "ghost",
        class: ["bg-transparent hover:bg-(--cyan-6)", "text-(--cyan-11)"]
      },

      // Green
      {
        colorScheme: "green",
        variant: "solid",
        class: [
          "bg-(--green-9) hover:bg-(--green-10)",
          "text-(--slate-1) dark:text-(--slate-12)"
        ]
      },
      {
        colorScheme: "green",
        variant: "soft",
        class: ["bg-(--green-5) hover:bg-(--green-6)", "text-(--green-11)"]
      },
      {
        colorScheme: "green",
        variant: "outline",
        class: [
          "outline outline-1 outline-(--green-7)",
          "bg-transparent hover:bg-(--green-6)",
          "text-(--green-11)"
        ]
      },
      {
        colorScheme: "green",
        variant: "ghost",
        class: ["bg-transparent hover:bg-(--green-6)", "text-(--green-11)"]
      },

      // Amber
      {
        colorScheme: "amber",
        variant: "solid",
        class: [
          "bg-(--amber-9) hover:bg-(--amber-10)",
          "text-(--slate-12) dark:text-(--slate-1)"
        ]
      },
      {
        colorScheme: "amber",
        variant: "soft",
        class: ["bg-(--amber-5) hover:bg-(--amber-6)", "text-(--amber-11)"]
      },
      {
        colorScheme: "amber",
        variant: "outline",
        class: [
          "outline outline-1 outline-(--amber-7)",
          "bg-transparent hover:bg-(--amber-6)",
          "text-(--amber-11)"
        ]
      },
      {
        colorScheme: "amber",
        variant: "ghost",
        class: ["bg-transparent hover:bg-(--amber-6)", "text-(--amber-11)"]
      },

      // Orange
      {
        colorScheme: "orange",
        variant: "solid",
        class: ["bg-(--orange-9) hover:bg-(--orange-10)", "text-(--slate-1)"]
      },
      {
        colorScheme: "orange",
        variant: "soft",
        class: ["bg-(--orange-5) hover:bg-(--orange-6)", "text-(--orange-11)"]
      },
      {
        colorScheme: "orange",
        variant: "outline",
        class: [
          "outline outline-1 outline-(--orange-7)",
          "bg-transparent hover:bg-(--orange-6)",
          "text-(--orange-11)"
        ]
      },
      {
        colorScheme: "orange",
        variant: "ghost",
        class: ["bg-transparent hover:bg-(--orange-6)", "text-(--orange-11)"]
      },

      // Gray
      {
        colorScheme: "gray",
        variant: "solid",
        class: ["bg-(--gray-12) hover:bg-(--gray-11)", "text-(--slate-1)"]
      },
      {
        colorScheme: "gray",
        variant: "soft",
        class: ["bg-(--gray-5) hover:bg-(--gray-6)", "text-(--slate-12)"]
      },
      {
        colorScheme: "gray",
        variant: "outline",
        class: [
          "outline outline-1 outline-(--gray-7)",
          "bg-transparent hover:bg-(--gray-6)",
          "text-(--gray-12)"
        ]
      },
      {
        colorScheme: "gray",
        variant: "ghost",
        class: ["bg-transparent hover:bg-(--gray-6)", "text-(--gray-12)"]
      }
    ],
    defaultVariants: {
      size: "md",
      colorScheme: "plum",
      variant: "solid",
      bump: true
    }
  }
);

type ButtonProps = ComponentProps<"button"> &
  VariantProps<typeof buttonVariants> & {
    leftIcon?: ReactNode;
    rightIcon?: ReactNode;
  };

export type { ButtonProps };

const Button: React.FC<ButtonProps> = ({
  size,
  colorScheme,
  variant,
  bump,
  leftIcon,
  rightIcon,
  children,
  className,
  disabled,
  ...props
}) => {
  return (
    <button
      className={buttonVariants({
        size,
        colorScheme,
        variant,
        bump: disabled ? false : bump,
        className
      })}
      disabled={disabled}
      {...props}
    >
      {leftIcon && leftIcon}
      {children}
      {rightIcon && rightIcon}
    </button>
  );
};

Button.displayName = "Button";

export { Button };
