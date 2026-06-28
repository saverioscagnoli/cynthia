import { Database } from "bun:sqlite";
import type { MainClient, NamedAPIResourceList } from "pokenode-ts";
import { BASE_URL, concurrent } from "./main";

async function fetchSprite(url: string | null): Promise<Buffer | null> {
  if (!url) return null;
  let res = await fetch(url);
  return Buffer.from(await res.arrayBuffer());
}

async function seedPokemons(client: MainClient, db: Database) {
  let res = await fetch(`${BASE_URL}/pokemon?limit=1000000`);
  let list = (await res.json()) as NamedAPIResourceList;

  const insertPokemon = db.prepare(`
    INSERT OR IGNORE INTO pokemons (id, name, base_exp, height, weight, is_default_form, species_id)
    VALUES ($id, $name, $base_exp, $height, $weight, $is_default_form, $species_id)
  `);

  const insertSprites = db.prepare(`
    INSERT OR IGNORE INTO pokemon_sprites (pokemon_id, front, back, front_shiny, back_shiny, front_female, back_female, front_female_shiny, back_female_shiny)
    VALUES ($pokemon_id, $front, $back, $front_shiny, $back_shiny, $front_female, $back_female, $front_female_shiny, $back_female_shiny)
  `);

  await concurrent(
    50,
    list.results.map((entry) => async () => {
      let p = await client.pokemon.getPokemonByName(entry.name);

      const pokemonInserted =
        insertPokemon.run({
          $id: p.id,
          $name: p.name,
          $base_exp: p.base_experience,
          $height: p.height,
          $weight: p.weight,
          $is_default_form: p.is_default ? 1 : 0,
          $species_id: Number(p.species.url.split("/").at(-2) ?? 0),
        }).changes > 0;

      let s = p.sprites;

      if (pokemonInserted) {
        insertSprites.run({
          $pokemon_id: p.id,
          $front: await fetchSprite(s.front_default),
          $back: await fetchSprite(s.back_default),
          $front_shiny: await fetchSprite(s.front_shiny),
          $back_shiny: await fetchSprite(s.back_shiny),
          $front_female: await fetchSprite(s.front_female),
          $back_female: await fetchSprite(s.back_female),
          $front_female_shiny: await fetchSprite(s.front_shiny_female),
          $back_female_shiny: await fetchSprite(s.back_shiny_female),
        });
      }

      console.log(p.id, p.name);
    }),
  );
}

export { seedPokemons };
