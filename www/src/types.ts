type TrainerSpriteMapEntry = {
  id: number;
  x: number;
  y: number;
  w: number;
  h: number;
  name: string;
};

type TrainerSpriteSheet = {
  _sheet: TrainerSpriteMapEntry;
  sprites: { [name: string]: TrainerSpriteMapEntry };
};

export type { TrainerSpriteSheet, TrainerSpriteMapEntry };
