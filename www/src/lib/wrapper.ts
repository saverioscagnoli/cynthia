import type { MatchStats } from "~/types";
import type { User } from "~/contexts/auth";
import { api } from "./request";
import { authHeader } from "./utils";

const privateApi = {
  getLoggedUser: async (token: string) =>
    await api<User>("/api/user/me", { headers: authHeader(token) }),

  updateUsername: async (token: string, username: string) =>
    await api<void>("/api/user/username", {
      method: "PUT",
      headers: authHeader(token),
      body: JSON.stringify({ username })
    }),

  updateTrainerSprite: async (token: string, spriteId: number) =>
    await api<void>("/api/user/sprite-id", {
      method: "PUT",
      headers: authHeader(token),
      body: JSON.stringify({ id: spriteId })
    }),

  updateBanner: async (token: string, file: File) => {
    let fd = new FormData();

    fd.append("banner", file);

    return await api<void>("/api/user/banner", {
      method: "PUT",
      headers: authHeader(token),
      body: fd
    });
  },

  deleteBanner: async (token: string) =>
    await api<void>("/api/user/banner", {
      method: "DELETE",
      headers: authHeader(token)
    })
};

const publicApi = {
  getUser: async (id: string) => await api<User>(`/api/user/${id}`),
  getUserStats: async (id: string) =>
    await api<MatchStats>(`/api/user/${id}/matches`)
};

export { privateApi, publicApi };
