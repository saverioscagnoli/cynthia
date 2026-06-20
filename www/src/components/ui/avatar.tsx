import type { ReactNode, RefObject } from "react";
import type React from "react";
import * as RxAvatar from "@radix-ui/react-avatar";
import { PersonIcon } from "@radix-ui/react-icons";
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
import "@radix-ui/colors/gray.css";
import "@radix-ui/colors/gray-dark.css";

const avatarVariants = cva(
  [
    "inline-flex justify-center items-center",
    "align-middle",
    "overflow-hidden",
    "select-none",
    "rounded-full",
    "font-medium"
  ],
  {
    variants: {
      size: {
        sm: "w-10 h-10 text-sm",
        md: "w-12 h-12 text-md",
        lg: "w-14 h-14 text-lg",
        xl: "w-16 h-16 text-xl"
      },
      colorScheme: {
        plum: "bg-(--plum-3)   text-(--plum-11)",
        red: "bg-(--red-3)    text-(--red-11)",
        blue: "bg-(--blue-3)   text-(--blue-11)",
        teal: "bg-(--teal-3)   text-(--teal-11)",
        cyan: "bg-(--cyan-3)   text-(--cyan-11)",
        green: "bg-(--green-3)  text-(--green-11)",
        amber: "bg-(--amber-3)  text-(--amber-11)",
        orange: "bg-(--orange-3) text-(--orange-11)",
        gray: "bg-(--gray-3)   text-(--slate-11)"
      }
    },
    defaultVariants: {
      size: "md",
      colorScheme: "plum"
    }
  }
);

const iconSizeMap: Record<NonNullable<AvatarProps["size"]>, number> = {
  sm: 14,
  md: 18,
  lg: 22,
  xl: 28
};

type AvatarProps = RxAvatar.AvatarProps &
  VariantProps<typeof avatarVariants> & {
    ref?: RefObject<HTMLSpanElement>;
    src?: string;
    fallback?: ReactNode | string;
    imageProps?: Omit<RxAvatar.AvatarImageProps, "src">;
    fallbackProps?: Omit<RxAvatar.AvatarFallbackProps, "children">;
  };

export type { AvatarProps };

const Avatar: React.FC<AvatarProps> = ({
  src,
  fallback,
  size = "md",
  colorScheme,
  imageProps,
  fallbackProps,
  className,
  ...props
}) => {
  const iconSize = iconSizeMap[size ?? "md"];

  return (
    <RxAvatar.Root
      {...props}
      className={avatarVariants({ className, size, colorScheme })}
    >
      <RxAvatar.Image
        className="h-full w-full rounded-[inherit] object-cover"
        src={src}
        {...imageProps}
      />
      <RxAvatar.Fallback {...fallbackProps}>
        {fallback ?? <PersonIcon width={iconSize} height={iconSize} />}
      </RxAvatar.Fallback>
    </RxAvatar.Root>
  );
};

Avatar.displayName = "Avatar";

export { Avatar };
