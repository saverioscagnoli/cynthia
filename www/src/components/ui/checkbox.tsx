import { type ReactNode, type RefObject } from "react";
import * as RxCheckbox from "@radix-ui/react-checkbox";
import { CheckIcon } from "@radix-ui/react-icons";
import { cva, type VariantProps } from "class-variance-authority";
import "@radix-ui/colors/plum.css";
import "@radix-ui/colors/plum-dark.css";
import "@radix-ui/colors/teal.css";
import "@radix-ui/colors/teal-dark.css";
import "@radix-ui/colors/grass.css";
import "@radix-ui/colors/grass-dark.css";
import "@radix-ui/colors/red.css";
import "@radix-ui/colors/red-dark.css";
import "@radix-ui/colors/amber.css";
import "@radix-ui/colors/amber-dark.css";
import "@radix-ui/colors/blue.css";
import "@radix-ui/colors/blue-dark.css";
import "@radix-ui/colors/gray.css";
import "@radix-ui/colors/gray-dark.css";
import "@radix-ui/colors/slate.css";
import "@radix-ui/colors/slate-dark.css";
import "./checkbox.css";

const checkboxVariants = cva(
  [
    "inline-flex items-center justify-center",
    "bg-transparent",
    "border data-[state=checked]:border-none border-(--gray-7)",
    "rounded",
    "disabled:cursor-not-allowed",
    "disabled:opacity-50"
  ],
  {
    variants: {
      size: {
        sm: "w-4 h-4",
        md: "w-5 h-5",
        lg: "w-6 h-6"
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
        outline: "border"
      }
    },
    compoundVariants: [
      {
        colorScheme: "plum",
        variant: "solid",
        class: [
          "data-[state=checked]:bg-(--plum-9)",
          "text-(--slate-1) dark:text-(--slate-12)"
        ]
      },
      {
        colorScheme: "plum",
        variant: "soft",
        class: [
          "border-(--plum-7)",
          "data-[state=checked]:bg-(--plum-5)",
          "text-(--plum-11)"
        ]
      },
      {
        colorScheme: "plum",
        variant: "outline",
        class: [
          "border-(--plum-7) data-[state=checked]:border-solid",
          "bg-transparent hover:bg-(--plum-6)",
          "text-(--plum-11)"
        ]
      },
      {
        colorScheme: "red",
        variant: "solid",
        class: [
          "data-[state=checked]:bg-(--red-9)",
          "text-(--slate-1) dark:text-(--slate-12)"
        ]
      },
      {
        colorScheme: "red",
        variant: "soft",
        class: [
          "border-(--red-7)",
          "data-[state=checked]:bg-(--red-5)",
          "text-(--red-11)"
        ]
      },
      {
        colorScheme: "red",
        variant: "outline",
        class: [
          "border-(--red-7) data-[state=checked]:border-solid",
          "bg-transparent hover:bg-(--red-6)",
          "text-(--red-11)"
        ]
      },
      {
        colorScheme: "blue",
        variant: "solid",
        class: [
          "data-[state=checked]:bg-(--blue-9)",
          "text-(--slate-1) dark:text-(--slate-12)"
        ]
      },
      {
        colorScheme: "blue",
        variant: "soft",
        class: [
          "border-(--blue-7)!",
          "data-[state=checked]:bg-(--blue-5)",
          "text-(--blue-11)"
        ]
      },
      {
        colorScheme: "blue",
        variant: "outline",
        class: [
          "border-(--blue-7)! data-[state=checked]:border-solid",
          "bg-transparent hover:bg-(--blue-6)",
          "text-(--blue-11)"
        ]
      },
      {
        colorScheme: "teal",
        variant: "solid",
        class: [
          "data-[state=checked]:bg-(--teal-9)",
          "text-(--slate-1) dark:text-(--slate-12)"
        ]
      },
      {
        colorScheme: "teal",
        variant: "soft",
        class: [
          "border-(--teal-7)!",
          "data-[state=checked]:bg-(--teal-5)",
          "text-(--teal-11)"
        ]
      },
      {
        colorScheme: "teal",
        variant: "outline",
        class: [
          "border-(--teal-7)! data-[state=checked]:border-solid",
          "bg-transparent hover:bg-(--teal-6)",
          "text-(--teal-11)"
        ]
      },
      {
        colorScheme: "cyan",
        variant: "solid",
        class: [
          "data-[state=checked]:bg-(--cyan-9)",
          "text-(--slate-1) dark:text-(--slate-12)"
        ]
      },
      {
        colorScheme: "cyan",
        variant: "soft",
        class: [
          "border-(--cyan-7)!",
          "data-[state=checked]:bg-(--cyan-5)",
          "text-(--cyan-11)"
        ]
      },
      {
        colorScheme: "cyan",
        variant: "outline",
        class: [
          "border-(--cyan-7)! data-[state=checked]:border-solid",
          "bg-transparent hover:bg-(--cyan-6)",
          "text-(--cyan-11)"
        ]
      },
      {
        colorScheme: "green",
        variant: "solid",
        class: [
          "data-[state=checked]:bg-(--grass-9)",
          "text-(--slate-1) dark:text-(--slate-12)"
        ]
      },
      {
        colorScheme: "green",
        variant: "soft",
        class: [
          "border-(--grass-7)!",
          "data-[state=checked]:bg-(--grass-5)",
          "text-(--grass-11)"
        ]
      },
      {
        colorScheme: "green",
        variant: "outline",
        class: [
          "border-(--grass-7)! data-[state=checked]:border-solid",
          "bg-transparent hover:bg-(--grass-6)",
          "text-(--grass-11)"
        ]
      },
      {
        colorScheme: "amber",
        variant: "solid",
        class: [
          "data-[state=checked]:bg-(--amber-9)",
          "text-(--slate-12) dark:text-(--slate-1)"
        ]
      },
      {
        colorScheme: "amber",
        variant: "soft",
        class: [
          "border-(--amber-7)!",
          "data-[state=checked]:bg-(--amber-5)",
          "text-(--amber-11)"
        ]
      },
      {
        colorScheme: "amber",
        variant: "outline",
        class: [
          "border-(--amber-7)! data-[state=checked]:border-solid",
          "bg-transparent hover:bg-(--amber-6)",
          "text-(--amber-11)"
        ]
      },
      {
        colorScheme: "orange",
        variant: "solid",
        class: ["data-[state=checked]:bg-(--orange-9)", "text-(--slate-1)"]
      },
      {
        colorScheme: "orange",
        variant: "soft",
        class: [
          "border-(--orange-7)!",
          "data-[state=checked]:bg-(--orange-5)",
          "text-(--orange-11)"
        ]
      },
      {
        colorScheme: "orange",
        variant: "outline",
        class: [
          "border-(--orange-7)! data-[state=checked]:border-solid",
          "bg-transparent hover:bg-(--orange-6)",
          "text-(--orange-11)"
        ]
      },
      {
        colorScheme: "gray",
        variant: "solid",
        class: ["data-[state=checked]:bg-(--gray-12)", "text-(--slate-1)"]
      },
      {
        colorScheme: "gray",
        variant: "soft",
        class: [
          "border-(--gray-7)",
          "data-[state=checked]:bg-(--gray-5)",
          "text-(--slate-12)"
        ]
      },
      {
        colorScheme: "gray",
        variant: "outline",
        class: [
          "border-(--gray-7) data-[state=checked]:border-solid",
          "bg-transparent hover:bg-(--gray-6)",
          "text-(--slate-12)"
        ]
      }
    ],
    defaultVariants: {
      size: "md",
      colorScheme: "plum",
      variant: "solid"
    }
  }
);

const checkboxIconVariants = cva(null, {
  variants: {
    size: {
      sm: "w-4 h-4",
      md: "w-5 h-5",
      lg: "w-6 h-6"
    }
  },
  defaultVariants: {
    size: "md"
  }
});

type CheckboxProps = RxCheckbox.CheckboxProps &
  VariantProps<typeof checkboxVariants> & {
    icon?: ReactNode;
    ref?: RefObject<HTMLButtonElement>;
  };

export type { CheckboxProps };

const Checkbox: React.FC<CheckboxProps> = ({
  className,
  colorScheme = "plum",
  variant = "solid",
  size = "md",
  icon = <CheckIcon className={checkboxIconVariants({ size })} />,
  ...props
}) => {
  return (
    <RxCheckbox.Root
      className={checkboxVariants({ className, colorScheme, variant, size })}
      {...props}
    >
      <RxCheckbox.Indicator className="checkbox-indicator" children={icon} />
    </RxCheckbox.Root>
  );
};

Checkbox.displayName = "Checkbox";

export { Checkbox };
