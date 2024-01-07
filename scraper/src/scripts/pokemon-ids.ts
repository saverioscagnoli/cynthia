import { Config } from "@config";
import { Scraper } from "@structs";
import { Fandom } from "@urls";
import { CSV } from "csv-rw";
import { findMarker } from "find-marker";
import path from "path";

async function scrape() {
  const scraper = new Scraper({ url: Fandom.PokemonList });

  scraper.on("loaded", async $ => {
    const SEARCH_PARAM = "were introduced in";

    // Get the tables containing the generation lists.
    let tables = $(`*:contains('${SEARCH_PARAM}')`)
      .next("table")
      .toArray()
      .slice(0, Config.MaxGenerations);

    // Get the table bodies.
    let bodys = tables.map(t => $(t).find("tbody").toArray());

    // Remove the first row from each table body.
    let trs = bodys.map(b => $(b).find("tr").toArray().slice(1)).flat();

    const pokemonIDsCsv = new CSV({
      path: path.join(findMarker(), "api/data/pokemon-ids.csv"),
      headers: ["n:id", "s:name"],
      deletePrevious: true
    });

    // Iterate over each table row and get the ID and name.
    for (let ts of trs) {
      let tds = $(ts).find("td").toArray();

      let id = +$(tds[0]).text().trim();
      let name = $(tds[2])
        .text()
        .trim()
        .toLowerCase()
        .replace(/♀/g, "-f")
        .replace(/♂/g, "-m")
        .replace(/[^a-z0-9-]/g, "-")
        .replace(/--+/g, "-")
        .replace(/^-+|-+$/g, "");

      pokemonIDsCsv.store({ id, name });
    }

    await pokemonIDsCsv.flush();
  });
}

export { scrape as scrapePokemonIDs };
