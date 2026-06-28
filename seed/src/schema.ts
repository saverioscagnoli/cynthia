import { Database } from "bun:sqlite";

function createAllSchemas(db: Database) {
  db.run("PRAGMA foreign_keys = ON");

  db.run(`
    CREATE TABLE IF NOT EXISTS pokemon_species (
      id INTEGER PRIMARY KEY,
      name TEXT NOT NULL,
      national_dex_number INTEGER NOT NULL,
      gender_chance INTEGER NOT NULL,
      capture_rate INTEGER NOT NULL,
      base_happiness INTEGER NOT NULL,
      is_legendary INTEGER NOT NULL,
      is_mythic INTEGER NOT NULL,
      hatch_counter INTEGER NOT NULL,
      has_gender_differences INTEGER NOT NULL,
      forms_switchable INTEGER NOT NULL,
      color TEXT NOT NULL
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS pokemons (
      id INTEGER PRIMARY KEY,
      name TEXT NOT NULL,
      base_exp INTEGER NOT NULL,
      height INTEGER NOT NULL,
      weight INTEGER NOT NULL,
      is_default_form INTEGER NOT NULL,
      species_id INTEGER NOT NULL,
      FOREIGN KEY (species_id) REFERENCES pokemon_species(id)
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS pokemon_sprites (
      pokemon_id INTEGER PRIMARY KEY,
      front BLOB,
      back BLOB,
      front_shiny BLOB,
      back_shiny BLOB,
      front_female BLOB,
      back_female BLOB,
      front_female_shiny BLOB,
      back_female_shiny BLOB,
      FOREIGN KEY (pokemon_id) REFERENCES pokemons(id)
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS types (
      id INTEGER PRIMARY KEY,
      name TEXT NOT NULL
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS type_sprites (
      type_id INTEGER PRIMARY KEY,
      sprite BLOB NOT NULL,
      FOREIGN KEY (type_id) REFERENCES types(id)
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS pokemon_types (
      pokemon_id INTEGER NOT NULL,
      type_id INTEGER NOT NULL,
      slot INTEGER NOT NULL,
      PRIMARY KEY (pokemon_id, slot),
      FOREIGN KEY (pokemon_id) REFERENCES pokemons(id),
      FOREIGN KEY (type_id) REFERENCES types(id)
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS moves (
      id INTEGER PRIMARY KEY,
      name TEXT NOT NULL,
      accuracy INTEGER,
      power INTEGER,
      base_pp INTEGER NOT NULL,
      priority INTEGER NOT NULL,
      effect_chance INTEGER,
      damage_class TEXT,
      type_id INTEGER,
      min_hits INTEGER,
      max_hits INTEGER,
      min_turns INTEGER,
      max_turns INTEGER,
      healing INTEGER,
      drain_or_recoil INTEGER,
      crit_rate_bonus INTEGER NOT NULL,
      flinch_chance INTEGER NOT NULL,
      stat_change_chance INTEGER NOT NULL,
      FOREIGN KEY (type_id) REFERENCES types(id)
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS pokemon_moves (
      pokemon_id INTEGER NOT NULL,
      move_id INTEGER NOT NULL,
      learn_method TEXT NOT NULL,
      min_level INTEGER NOT NULL,
      PRIMARY KEY (pokemon_id, move_id),
      FOREIGN KEY (pokemon_id) REFERENCES pokemons(id),
      FOREIGN KEY (move_id) REFERENCES moves(id)
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS stats (
      id INTEGER PRIMARY KEY,
      name TEXT NOT NULL,
      is_battle_only INTEGER NOT NULL
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS pokemon_stats (
      stat_id INTEGER NOT NULL,
      pokemon_id INTEGER NOT NULL,
      base_stat INTEGER NOT NULL,
      PRIMARY KEY (stat_id, pokemon_id),
      FOREIGN KEY (stat_id) REFERENCES stats(id),
      FOREIGN KEY (pokemon_id) REFERENCES pokemons(id)
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS items (
      id INTEGER PRIMARY KEY,
      name TEXT NOT NULL,
      cost INTEGER NOT NULL,
      fling_power INTEGER,
      fling_effect TEXT
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS item_sprites (
      item_id INTEGER PRIMARY KEY,
      sprite BLOB NOT NULL,
      FOREIGN KEY (item_id) REFERENCES items(id)
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS held_items (
      item_id INTEGER NOT NULL,
      pokemon_id INTEGER NOT NULL,
      rarity INTEGER NOT NULL,
      PRIMARY KEY (item_id, pokemon_id),
      FOREIGN KEY (item_id) REFERENCES items(id),
      FOREIGN KEY (pokemon_id) REFERENCES pokemons(id)
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS locations (
      id INTEGER PRIMARY KEY,
      name TEXT NOT NULL,
      region_name TEXT NOT NULL
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS location_areas (
      id INTEGER PRIMARY KEY,
      location_id INTEGER NOT NULL,
      FOREIGN KEY (location_id) REFERENCES locations(id)
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS encounters (
      pokemon_id INTEGER NOT NULL,
      location_area_id INTEGER NOT NULL,
      min_level INTEGER NOT NULL,
      max_level INTEGER NOT NULL,
      method TEXT NOT NULL,
      PRIMARY KEY (pokemon_id, location_area_id),
      FOREIGN KEY (pokemon_id) REFERENCES pokemons(id),
      FOREIGN KEY (location_area_id) REFERENCES location_areas(id)
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS abilities (
      id INTEGER PRIMARY KEY,
      name TEXT NOT NULL,
      long_effect TEXT NOT NULL,
      short_effect TEXT NOT NULL
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS pokemon_abilities (
      id INTEGER NOT NULL,
      pokemon_id INTEGER NOT NULL,
      is_hidden INTEGER NOT NULL,
      slot INTEGER NOT NULL,
      PRIMARY KEY (id, pokemon_id),
      FOREIGN KEY (id) REFERENCES abilities(id),
      FOREIGN KEY (pokemon_id) REFERENCES pokemons(id)
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS natures (
      id INTEGER PRIMARY KEY,
      name TEXT NOT NULL,
      increased_stat_id INTEGER NOT NULL,
      decreased_stat_id INTEGER NOT NULL,
      FOREIGN KEY (increased_stat_id) REFERENCES stats(id),
      FOREIGN KEY (decreased_stat_id) REFERENCES stats(id)
    )
  `);

  db.run(`
    CREATE TABLE IF NOT EXISTS evolution_details (
      species_id INTEGER NOT NULL,
      evolves_to_species_id INTEGER NOT NULL,
      pokemon_id INTEGER NOT NULL,
      trigger TEXT NOT NULL,
      gender TEXT,
      held_item_id INTEGER,
      known_move_id INTEGER,
      known_move_type_id INTEGER,
      location_id INTEGER,
      min_level INTEGER NOT NULL,
      min_happiness INTEGER,
      min_beauty INTEGER,
      min_affection INTEGER,
      needs_multiplayer INTEGER NOT NULL,
      needs_rain INTEGER NOT NULL,
      party_species_id INTEGER,
      party_type_id INTEGER,
      relative_physical_stats INTEGER,
      time_of_day INTEGER,
      trade_species_id INTEGER,
      turn_upside_down INTEGER NOT NULL,
      used_move_id INTEGER,
      min_move_count INTEGER NOT NULL,
      min_steps INTEGER,
      min_damage_taken INTEGER,
      PRIMARY KEY (species_id, evolves_to_species_id, pokemon_id),
      FOREIGN KEY (species_id) REFERENCES pokemon_species(id),
      FOREIGN KEY (evolves_to_species_id) REFERENCES pokemon_species(id),
      FOREIGN KEY (pokemon_id) REFERENCES pokemons(id),
      FOREIGN KEY (held_item_id) REFERENCES items(id),
      FOREIGN KEY (known_move_id) REFERENCES moves(id),
      FOREIGN KEY (known_move_type_id) REFERENCES types(id),
      FOREIGN KEY (location_id) REFERENCES locations(id),
      FOREIGN KEY (used_move_id) REFERENCES moves(id)
    )
  `);
}

export { createAllSchemas };
