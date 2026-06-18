import { type TrainerSpriteSheetMap } from "~/types";

const api = "http://localhost:9000";

const ApiEndpoints = {
  GetTrainersSpriteSheetMap: "/sprites/trainer/sheet/map",
} as const;

async function getTrainersSpritesheetMap(): Promise<TrainerSpriteSheetMap> {
  let res = await fetch(api + ApiEndpoints.GetTrainersSpriteSheetMap);
  return res.json();
}

export { ApiEndpoints, getTrainersSpritesheetMap };
