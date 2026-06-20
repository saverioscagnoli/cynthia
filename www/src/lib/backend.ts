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

export { getLoggedUser, uploadBanner };
