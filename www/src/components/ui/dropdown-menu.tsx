import React, {
  createContext,
  useContext,
  type ReactNode,
  type RefObject
} from "react";
import * as RxDropdownMenu from "@radix-ui/react-dropdown-menu";
import {
  CheckIcon,
  ChevronRightIcon,
  DotFilledIcon
} from "@radix-ui/react-icons";
import { cva } from "class-variance-authority";
import { cn } from "~/lib/utils";
import "@radix-ui/colors/slate.css";
import "@radix-ui/colors/slate-dark.css";
import "@radix-ui/colors/plum.css";
import "@radix-ui/colors/plum-dark.css";
import "./dropdown-menu.css";

type DropdownMenuColorScheme =
  | "plum"
  | "red"
  | "blue"
  | "teal"
  | "cyan"
  | "green"
  | "amber"
  | "orange"
  | "gray";

export type { DropdownMenuColorScheme };

const ColorSchemeContext = createContext<DropdownMenuColorScheme | null>(null);

const useColorScheme = () => {
  const context = useContext(ColorSchemeContext);

  if (context === null) {
    throw new Error(
      "useColorScheme must be used within a DropdownMenuProvider"
    );
  }

  return context;
};

type DropdownMenuProps = RxDropdownMenu.DropdownMenuProps & {
  colorScheme?: DropdownMenuColorScheme;
};

export type { DropdownMenuProps };

type DropdownMenuComponent = React.FC<DropdownMenuProps> & {
  Trigger: typeof DropdownMenuTrigger;
  Content: typeof DropdownMenuContent;
  Label: typeof DropdownMenuLabel;
  Item: typeof DropdownMenuItem;
  CheckboxItem: typeof DropdownMenuCheckboxItem;
  RadioGroup: typeof DropdownMenuRadioGroup;
  RadioItem: typeof DropdownMenuRadioItem;
  Sub: typeof DropdownMenuSub;
  SubTrigger: typeof DropdownMenuSubTrigger;
  SubContent: typeof DropdownMenuSubContent;
  Separator: typeof DropdownMenuSeparator;
};

const LeftIcon: React.FC<{ icon: ReactNode }> = ({ icon }) => {
  return (
    <div
      className={cn("pr-2 pl-1.5", "inline-flex items-center justify-center")}
    >
      {icon}
    </div>
  );
};

const DropdownMenu: DropdownMenuComponent = ({
  colorScheme = "plum",
  ...props
}) => {
  return (
    <ColorSchemeContext.Provider value={colorScheme}>
      <RxDropdownMenu.Root {...props} />
    </ColorSchemeContext.Provider>
  );
};

type DropdownMenuTriggerProps = RxDropdownMenu.DropdownMenuTriggerProps & {
  ref?: RefObject<HTMLButtonElement>;
};

export type { DropdownMenuTriggerProps };

const DropdownMenuTrigger: React.FC<DropdownMenuTriggerProps> = props => (
  <RxDropdownMenu.Trigger {...props} />
);

type DropdownMenuContentProps = RxDropdownMenu.DropdownMenuContentProps & {
  ref?: RefObject<HTMLDivElement>;
};

export type { DropdownMenuContentProps };

const DropdownMenuContent: React.FC<DropdownMenuContentProps> = ({
  className,
  ...props
}) => {
  return (
    <RxDropdownMenu.Portal>
      <RxDropdownMenu.Content
        className={cn(
          "h-fit min-w-fit",
          "p-1",
          "border border-(--gray-7)",
          "rounded shadow-md",
          "bg-(--slate-1)",
          "text-(--slate-12)",
          "select-none",
          "dropdown-menu-content",
          className
        )}
        {...props}
      />
    </RxDropdownMenu.Portal>
  );
};

type DropdownMenuLabelProps = RxDropdownMenu.DropdownMenuLabelProps & {
  ref?: RefObject<HTMLDivElement>;
};

export type { DropdownMenuLabelProps };

const DropdownMenuLabel: React.FC<DropdownMenuLabelProps> = ({
  className,
  ...props
}) => {
  return (
    <RxDropdownMenu.Label className={cn("px-3 py-1", "text-xs")} {...props} />
  );
};

const dropdownMenuItemVariants = cva(
  [
    "flex items-center",
    "pl-1 pr-4 py-[1px]",
    "text-sm",
    "rounded-xs",
    "outline-none font-normal",
    "focus:cursor-default",
    "data-[disabled]:opacity-30",
    "data-[disabled]:pointer-events-none"
  ],
  {
    variants: {
      colorScheme: {
        plum: "focus:bg-(--plum-9) focus:text-(--slate-1)",
        red: "focus:bg-(--red-9) focus:text-(--slate-1)",
        blue: "focus:bg-(--blue-9) focus:text-(--slate-1)",
        teal: "focus:bg-(--teal-9) focus:text-(--slate-1)",
        cyan: "focus:bg-(--cyan-9) focus:text-(--slate-1)",
        green: "focus:bg-(--green-9) focus:text-(--slate-1)",
        amber:
          "focus:bg-(--amber-9) focus:text-(--slate-12) focus:dark:text-(--slate-1)",
        orange: "focus:bg-(--orange-9) focus:text-(--slate-1)",
        gray: "focus:bg-(--gray-4) focus:text-(--slate-12)"
      }
    }
  }
);

type DropdownMenuItemProps = RxDropdownMenu.DropdownMenuItemProps & {
  colorScheme?: DropdownMenuColorScheme;
  leftIcon?: ReactNode;
  ref?: RefObject<HTMLDivElement>;
};

export type { DropdownMenuItemProps };

const DropdownMenuItem: React.FC<DropdownMenuItemProps> = ({
  colorScheme,
  leftIcon,
  children,
  className,
  ...props
}) => {
  return (
    <RxDropdownMenu.Item
      className={dropdownMenuItemVariants({
        className,
        colorScheme: colorScheme ?? useColorScheme()
      })}
      {...props}
    >
      <LeftIcon icon={leftIcon} />
      {children}
    </RxDropdownMenu.Item>
  );
};

type DropdownMenuCheckboxItemProps =
  RxDropdownMenu.DropdownMenuCheckboxItemProps & {
    colorScheme?: DropdownMenuColorScheme;
    icon?: ReactNode;
    ref?: RefObject<HTMLDivElement>;
  };

export type { DropdownMenuCheckboxItemProps };

const DropdownMenuCheckboxItem: React.FC<DropdownMenuCheckboxItemProps> = ({
  checked,
  colorScheme,
  icon = <CheckIcon />,
  className,
  children,
  ...props
}) => {
  return (
    <RxDropdownMenu.CheckboxItem
      className={dropdownMenuItemVariants({
        className,
        colorScheme: colorScheme ?? useColorScheme()
      })}
      checked={checked}
      {...props}
    >
      <LeftIcon
        icon={
          checked ? (
            <RxDropdownMenu.ItemIndicator>{icon}</RxDropdownMenu.ItemIndicator>
          ) : (
            <span className="px-1.5" />
          )
        }
      />
      {children}
    </RxDropdownMenu.CheckboxItem>
  );
};

type DropdownMenuRadioGroupProps = RxDropdownMenu.DropdownMenuRadioGroupProps;

export type { DropdownMenuRadioGroupProps };

const DropdownMenuRadioGroup = RxDropdownMenu.DropdownMenuRadioGroup;

type DropdownMenuRadioItemProps = RxDropdownMenu.DropdownMenuRadioItemProps & {
  colorScheme?: DropdownMenuColorScheme;
  icon?: ReactNode;
  ref?: RefObject<HTMLDivElement>;
};

export type { DropdownMenuRadioItemProps };

const DropdownMenuRadioItem: React.FC<DropdownMenuRadioItemProps> = ({
  colorScheme,
  icon = <DotFilledIcon />,
  children,
  className,
  ...props
}) => {
  return (
    <RxDropdownMenu.RadioItem
      className={dropdownMenuItemVariants({
        className,
        colorScheme: colorScheme ?? useColorScheme()
      })}
      {...props}
    >
      <div className={cn("w-7", "flex items-center justify-center")}>
        <RxDropdownMenu.ItemIndicator>{icon}</RxDropdownMenu.ItemIndicator>
      </div>
      {children}
    </RxDropdownMenu.RadioItem>
  );
};

type DropdownMenuSubProps = RxDropdownMenu.DropdownMenuSubProps;

export type { DropdownMenuSubProps };

const DropdownMenuSub = RxDropdownMenu.DropdownMenuSub;

type DropdownMenuSubTriggerProps =
  RxDropdownMenu.DropdownMenuSubTriggerProps & {
    colorScheme?: DropdownMenuColorScheme;
    leftIcon?: ReactNode;
    icon?: ReactNode;
  };

export type { DropdownMenuSubTriggerProps };

const DropdownMenuSubTrigger: React.FC<DropdownMenuSubTriggerProps> = ({
  colorScheme,
  leftIcon,
  icon = <ChevronRightIcon />,
  children,
  className,
  ...props
}) => {
  return (
    <RxDropdownMenu.SubTrigger
      className={dropdownMenuItemVariants({
        className: cn("relative", className),
        colorScheme: colorScheme ?? useColorScheme()
      })}
      {...props}
    >
      <LeftIcon icon={leftIcon} />
      {children}
      <span className={cn("absolute right-1")}>{icon}</span>
    </RxDropdownMenu.SubTrigger>
  );
};

type DropdownMenuSubContentProps =
  RxDropdownMenu.DropdownMenuSubContentProps & {
    ref?: RefObject<HTMLDivElement>;
  };

export type { DropdownMenuSubContentProps };

const DropdownMenuSubContent: React.FC<DropdownMenuSubContentProps> = ({
  className,
  ...props
}) => {
  return (
    <RxDropdownMenu.Portal>
      <RxDropdownMenu.SubContent
        className={cn(
          "h-fit min-w-fit",
          "p-1",
          "border border-(--gray-7)",
          "rounded shadow-md",
          "bg-(--slate-1)",
          "text-(--slate-12)",
          "select-none",
          "dropdown-menu-subcontent",
          className
        )}
        {...props}
      />
    </RxDropdownMenu.Portal>
  );
};

type DropdownMenuSeparatorProps = RxDropdownMenu.DropdownMenuSeparatorProps & {
  ref?: RefObject<HTMLDivElement>;
};

export type { DropdownMenuSeparatorProps };

const DropdownMenuSeparator: React.FC<DropdownMenuSeparatorProps> = ({
  className,
  ...props
}) => {
  return (
    <RxDropdownMenu.Separator
      className={cn("my-1 h-px", "bg-(--gray-7)", className)}
      {...props}
    />
  );
};

DropdownMenu.displayName = "DropdownMenu";
DropdownMenuTrigger.displayName = "DropdownMenuTrigger";
DropdownMenuContent.displayName = "DropdownMenuContent";
DropdownMenuLabel.displayName = "DropdownMenuLabel";
DropdownMenuItem.displayName = "DropdownMenuItem";
DropdownMenuCheckboxItem.displayName = "DropdownMenuCheckboxItem";
DropdownMenuRadioGroup.displayName = "DropdownMenuRadioGroup";
DropdownMenuRadioItem.displayName = "DropdownMenuRadioItem";
DropdownMenuSub.displayName = "DropdownMenuSub";
DropdownMenuSubTrigger.displayName = "DropdownMenuSubTrigger";
DropdownMenuSubContent.displayName = "DropdownMenuSubContent";
DropdownMenuSeparator.displayName = "DropdownMenuSeparator";

DropdownMenu.Trigger = DropdownMenuTrigger;
DropdownMenu.Content = DropdownMenuContent;
DropdownMenu.Label = DropdownMenuLabel;
DropdownMenu.Item = DropdownMenuItem;
DropdownMenu.CheckboxItem = DropdownMenuCheckboxItem;
DropdownMenu.RadioGroup = DropdownMenuRadioGroup;
DropdownMenu.RadioItem = DropdownMenuRadioItem;
DropdownMenu.Sub = DropdownMenuSub;
DropdownMenu.SubTrigger = DropdownMenuSubTrigger;
DropdownMenu.SubContent = DropdownMenuSubContent;
DropdownMenu.Separator = DropdownMenuSeparator;

export { DropdownMenu };
