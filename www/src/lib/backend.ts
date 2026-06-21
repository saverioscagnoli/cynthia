import type { User } from "~/contexts/auth";

async function getLoggedUser(token: string): Promise<User> {
  let res = await fetch("/user/me", {
    headers: { Authorization: `Bearer ${token}` }
  });

  if (res.status === 401) {
    throw new Error("unauthorized");
  }

  return await res.json();
}

async function updateUsername(token: string, username: string): Promise<void> {
  let res = await fetch("/user/username", {
    method: "PUT",
    headers: { Authorization: `Bearer ${token}` },
    body: JSON.stringify({ username })
  });

  if (!res.ok) {
    throw new Error(await res.text());
  }
}

async function updateTrainerSprite(
  token: string,
  spriteId: number
): Promise<void> {
  let res = await fetch("/user/sprite-id", {
    method: "PUT",
    headers: { Authorization: `Bearer ${token}` },
    body: JSON.stringify({ id: spriteId })
  });

  if (!res.ok) {
    throw new Error(await res.text());
  }
}

async function uploadBanner(token: string, file: File): Promise<void> {
  let formData = new FormData();

  formData.append("banner", file);

  let res = await fetch("/user/banner", {
    method: "PUT",
    headers: {
      Authorization: `Bearer ${token}`
    },
    body: formData
  });

  if (!res.ok) {
    throw new Error(await res.text());
  }
}

async function deleteBanner(token: string): Promise<void> {
  let res = await fetch("/user/banner", {
    method: "DELETE",
    headers: { Authorization: `Bearer ${token}` }
  });

  if (!res.ok) {
    throw new Error("Failed to clear user banner");
  }
}

export {
  getLoggedUser,
  updateUsername,
  updateTrainerSprite,
  uploadBanner,
  deleteBanner
};
