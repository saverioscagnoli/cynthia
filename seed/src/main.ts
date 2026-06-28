//import { PokemonClient } from "pokenode-ts";
//
//async function main() {
//  const client = new PokemonClient({ baseURL: "http://localhost/api/v2" });
//  const BATCH_SIZE = 10;
//
//  for (let i = 1; i <= 1350; i += BATCH_SIZE) {
//    const ids = Array.from({ length: BATCH_SIZE }, (_, j) => i + j).filter(
//      (id) => id <= 1350,
//    );
//    const results = await Promise.all(
//      ids.map((id) => client.getPokemonById(id)),
//    );
//    results.forEach((res) => console.log(res.name));
//  }
//}
//
//main();

import { Database } from "bun:sqlite";
import { seedPokemons } from "./pokemon";
import { MainClient } from "pokenode-ts";
import { createAllSchemas } from "./schema";
import { seedPokemonSpecies } from "./species";

export const BASE_URL = "http://localhost/api/v2";

export async function concurrent<T>(
  concurrency: number,
  tasks: (() => Promise<T>)[],
): Promise<T[]> {
  const results: T[] = new Array(tasks.length);
  let index = 0;

  async function worker() {
    while (index < tasks.length) {
      let i = index++;
      results[i] = await tasks[i]!();
    }
  }

  await Promise.all(Array.from({ length: concurrency }, worker));
  return results;
}

let client = new MainClient({ baseURL: "http://localhost/api/v2" });
let db = new Database("pokemon.db");

createAllSchemas(db);

await seedPokemonSpecies(client, db);
await seedPokemons(client, db);

db.close();
