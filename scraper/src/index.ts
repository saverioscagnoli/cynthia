import { Scraper } from "@structs";
import { Fandom } from "@urls";

const scraper = new Scraper({ url: Fandom.PokemonList });


scraper.on("loaded", $ => {
  let title =  $(".mw-page-title-main").text();

  console.log(title);
});