import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

function cn(...classes: ClassValue[]) {
  return twMerge(clsx(classes));
}

function dsAvatar(id: string, hash: string) {
  return `https://cdn.discordapp.com/avatars/${id}/${hash}.png`;
}

export { cn, dsAvatar };
