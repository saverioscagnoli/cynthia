import * as cheerio from "cheerio";
import fs from "fs";
import { request, fetch } from "undici";
import { EventEmitter } from "./event-emitter";
import { Config } from "@config";

interface ScraperOptions {
  /**
   * The URL to scrape.
   */
  url: string;
}

interface ScraperEvents {
  loaded: [selector: cheerio.CheerioAPI];
}

class Scraper extends EventEmitter<ScraperEvents> {
  public constructor({ url }: ScraperOptions) {
    super();

    this.init(url);
  }

  private async init(url: string) {
    let html = await this.getHtml(url);
    let $ = this.get$(html);

    this.emit("loaded", $);
  }

  protected async getHtml(url: string) {
    let controller = new AbortController();
    let timeout = setTimeout(() => controller.abort(), Config.RequestTimeout);

    let res = await request(url, {
      method: "GET",
      headers: { "User-Agent": "*" },
      signal: controller.signal
    });

    clearTimeout(timeout);
    return await res.body.text();
  }

  protected get$(html: string) {
    return cheerio.load(html);
  }

  protected async download(url: string, output: string): Promise<boolean> {
    try {
      let res = await fetch(url);
      if (!res.ok) return false;

      let bytes = await res.arrayBuffer();
      fs.writeFileSync(output, Buffer.from(bytes));
    } catch (err) {
      console.error(err);
      return false;
    }
  }
}

export { Scraper };
export type { ScraperOptions };
