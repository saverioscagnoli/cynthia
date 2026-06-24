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

type MatchStats = {
  wins: number;
  losses: number;
  draws: number;
  winrate: number;

  current_streak: number;
  best_streak: number;

  single_wins: number;
  double_wins: number;
  total_damage_dealt: number;
  total_damage_received: number;
  pokemon_fainted: number;
  pokemon_lost: number;
};

export type { TrainerSpriteSheet, TrainerSpriteMapEntry, MatchStats };
