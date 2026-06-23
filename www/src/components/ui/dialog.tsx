import React, { type RefObject } from "react";
import * as RxDialog from "@radix-ui/react-dialog";
import { Cross1Icon } from "@radix-ui/react-icons";
import { cn } from "~/lib/utils";
import "@radix-ui/colors/gray.css";
import "@radix-ui/colors/gray-dark.css";
import "@radix-ui/colors/gray-alpha.css";
import "@radix-ui/colors/gray-dark-alpha.css";
import "@radix-ui/colors/slate.css";
import "@radix-ui/colors/slate-dark.css";
import "./dialog.css";

type DialogProps = RxDialog.DialogProps;

export type { DialogProps };

type DialogComponent = typeof RxDialog.Root & {
  Trigger: typeof DialogTrigger;
  Content: typeof DialogContent;
  Title: typeof DialogTitle;
  Description: typeof DialogDescription;
  Close: typeof DialogClose;
};

const Dialog = RxDialog.Root as DialogComponent;

type DialogTriggerProps = RxDialog.DialogTriggerProps;

export type { DialogTriggerProps };

const DialogTrigger = RxDialog.Trigger;

type DialogContentProps = RxDialog.DialogContentProps & {
  ref?: RefObject<HTMLDivElement>;
  overlayClassName?: string;
  overlayStyle?: React.CSSProperties;
};

export type { DialogContentProps };

const DialogContent: React.FC<DialogContentProps> = ({
  ref,
  className,
  overlayClassName,
  overlayStyle,
  ...props
}) => {
  return (
    <RxDialog.Portal>
      <RxDialog.Overlay
        className={cn(
          "fixed inset-0 z-9998",
          "bg-(--gray-a9) dark:bg-(--gray-a3)",
          "dialog-overlay",
          overlayClassName
        )}
        style={overlayStyle}
      />
      <RxDialog.Content
        className={cn(
          "max-h-[90vh] w-11/12 max-w-125 p-6",
          "fixed top-1/2 left-1/2 z-9999",
          "-translate-x-1/2 -translate-y-1/2 transform",
          "bg-(--slate-1)",
          "rounded-md shadow-md",
          "outline outline-(--gray-7) dark:outline",
          "dialog-content",
          className
        )}
        {...props}
        ref={ref}
      />
    </RxDialog.Portal>
  );
};

type DialogTitleProps = RxDialog.DialogTitleProps & {
  ref?: RefObject<HTMLHeadingElement>;
};

export type { DialogTitleProps };

const DialogTitle: React.FC<DialogTitleProps> = ({
  ref,
  className,
  ...props
}) => {
  return (
    <RxDialog.Title
      className={cn("text-xl font-semibold text-(--slate-12)", className)}
      ref={ref}
      {...props}
    />
  );
};

type DialogDescriptionProps = RxDialog.DialogDescriptionProps & {
  ref?: RefObject<HTMLParagraphElement>;
};

export type { DialogDescriptionProps };

const DialogDescription: React.FC<DialogDescriptionProps> = ({
  ref,
  className,
  ...props
}) => {
  return (
    <RxDialog.Description
      className={cn("mt-2", "font-normal text-(--slate-12)", className)}
      ref={ref}
      {...props}
    />
  );
};

type DialogCloseProps = RxDialog.DialogCloseProps & {
  ref?: RefObject<HTMLButtonElement>;
};

export type { DialogCloseProps };

const DialogClose: React.FC<DialogCloseProps> = ({
  className,
  children,
  ...props
}) => {
  if (children) {
    return (
      <RxDialog.Close className={className} {...props}>
        {children}
      </RxDialog.Close>
    );
  } else {
    return (
      <RxDialog.Close
        className={cn(
          "absolute top-4 right-4 h-7 w-7",
          "inline-flex items-center justify-center",
          "rounded-md",
          "bg-transparent text-(--slate-12)",
          "hover:bg-(--gray-5) active:translate-y-px",
          "transition-all duration-100",
          className
        )}
      >
        <Cross1Icon />
      </RxDialog.Close>
    );
  }
};

Dialog.displayName = "Dialog";
DialogTrigger.displayName = "DialogTrigger";
DialogContent.displayName = "DialogContent";
DialogTitle.displayName = "DialogTitle";
DialogDescription.displayName = "DialogDescription";
DialogClose.displayName = "DialogClose";

Dialog.Trigger = DialogTrigger;
Dialog.Content = DialogContent;
Dialog.Title = DialogTitle;
Dialog.Description = DialogDescription;
Dialog.Close = DialogClose;

export { Dialog };
