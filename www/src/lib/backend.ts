import type { Trainer } from "~/types";

const backend = "http://localhost:3247";

const BackendEndpoints = {
  GetTrainerMe: "/user/trainer/me",
} as const;

async function getTrainerMe(id: string): Promise<Trainer> {
  let res = await fetch(backend + BackendEndpoints.GetTrainerMe, {
    headers: { "X-Discord-ID": id },
  });

  return res.json();
}

export { BackendEndpoints, getTrainerMe };
