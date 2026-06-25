import type { MatchStats } from "~/types";
import type { User } from "~/contexts/auth";
import { api } from "./request";

const privateApi = {
  getLoggedUser: async () =>
    await api<User>("/api/user/me", { credentials: "include" }),

  updateUsername: async (username: string) =>
    await api<void>("/api/user/username", {
      method: "PUT",
      credentials: "include",
      body: JSON.stringify({ username })
    }),

  updateTrainerSprite: async (spriteId: number) =>
    await api<void>("/api/user/sprite-id", {
      method: "PUT",
      credentials: "include",
      body: JSON.stringify({ id: spriteId })
    }),

  updateBanner: async (file: File) => {
    let fd = new FormData();
    fd.append("banner", file);
    return await api<void>("/api/user/banner", {
      method: "PUT",
      credentials: "include",
      body: fd
    });
  },

  deleteBanner: async () =>
    await api<void>("/api/user/banner", {
      method: "DELETE",
      credentials: "include"
    })
};

const publicApi = {
  getUser: async (id: string) => await api<User>(`/api/user/${id}`),
  getUserStats: async (id: string) =>
    await api<MatchStats>(`/api/user/${id}/matches`)
};

export { privateApi, publicApi };
