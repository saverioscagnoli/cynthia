import { Database } from "bun:sqlite";
import type { MainClient, NamedAPIResourceList } from "pokenode-ts";
import { BASE_URL, concurrent } from "./main";

async function seedPokemonSpecies(client: MainClient, db: Database) {
  const res = await fetch(BASE_URL + "/pokemon-species");
  const list = (await res.json()) as NamedAPIResourceList;

  const insert = db.prepare(`
    INSERT OR IGNORE INTO pokemon_species (id, name, national_dex_number, gender_chance, capture_rate, base_happiness, is_legendary, is_mythic, hatch_counter, has_gender_differences, forms_switchable, color)
    VALUES ($id, $name, $national_dex_number, $gender_chance, $capture_rate, $base_happiness, $is_legendary, $is_mythic, $hatch_counter, $has_gender_differences, $forms_switchable, $color)
  `);

  await concurrent(
    50,
    Array.from({ length: list.count }, (_, i) => async () => {
      let s = await client.pokemon.getPokemonSpeciesById(i + 1);

      insert.run({
        $id: s.id,
        $name: s.name,
        $national_dex_number: s.pokedex_numbers[0]?.entry_number ?? 0,
        $gender_chance: s.gender_rate,
        $capture_rate: s.capture_rate,
        $base_happiness: s.base_happiness,
        $is_legendary: s.is_legendary ? 1 : 0,
        $is_mythic: s.is_mythical ? 1 : 0,
        $hatch_counter: s.hatch_counter,
        $has_gender_differences: s.has_gender_differences ? 1 : 0,
        $forms_switchable: s.forms_switchable ? 1 : 0,
        $color: s.color.name,
      });

      console.log(s.id, s.name);
    }),
  );
}

export { seedPokemonSpecies };
