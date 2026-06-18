type TrainerSpriteMapEntry = {
  x: number;
  y: number;
  w: number;
  h: number;
  name: string;
};

type TrainerSpriteSheetMap = { [r: string]: TrainerSpriteMapEntry };

type Trainer = {
  id: string;
  money: number;
  sprite_id: number;
};

export type { TrainerSpriteSheetMap, TrainerSpriteMapEntry, Trainer };
