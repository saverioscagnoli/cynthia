type TrainerSpriteMapEntry = {
    x: number;
    y: number;
    w: number;
    h: number;
    name: string;
};

type TrainerSpriteSheet = { [r: string]: TrainerSpriteMapEntry };

export type { TrainerSpriteSheet, TrainerSpriteMapEntry };
