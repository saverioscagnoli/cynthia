import { DELIMITER } from "@consts";
import { Entry, Value } from "@types";
import {
  existsSync,
  unlinkSync,
  writeFileSync,
  PathLike,
  createWriteStream,
  createReadStream
} from "fs";
import split2 from "split2";

interface CSVOptions<T extends string> {
  /**
   * The source of the CSV file.
   */
  src: PathLike;

  /**
   * The header of the CSV file.
   */
  header: T[];

  /**
   * Whether to delete the file with the same source if it exists.
   * @default false
   */
  deletePrevious?: boolean;
}

class CSV<T extends string> {
  private src: PathLike;
  private header: T[];

  public constructor({ src, header, deletePrevious }: CSVOptions<T>) {
    this.src = src;
    this.header = this.stripHeader(header);

    this.init(src, !!deletePrevious);
  }

  private init(src: PathLike, deletePrevious: boolean) {
    if (existsSync(src) && deletePrevious) {
      console.log("Deleting previous file...");
      unlinkSync(src);
    }

    writeFileSync(src, this.header.join(DELIMITER) + "\n");
  }

  /**
   * Function to strip the headers of any prefixes or suffixes used for typing.
   * @param headers The array of headers to use when parsing the CSV file.
   * @private
   */
  private stripHeader(headers: T[]): T[] {
    if (!headers || headers.length === 0) return [];

    return headers.map(h => h.replace(/^(n|b|s):|\?$/g, "").trim()) as T[];
  }

  public async read(options: { limit?: number } = {}) {
    let data = [] as Entry<T>[];

    let stream = createReadStream(this.src, { encoding: "utf-8" });

    let limit = options.limit || Infinity;
    let first = true;

    stream.pipe(split2()).on("data", (line: string) => {
      if (first) {
        first = false;
        return;
      }

      if (data.length >= limit) {
        stream.destroy();
        return;
      }

      let values = line.split(DELIMITER);
      let entry = {} as Record<string, unknown>;

      for (let i = 0; i < values.length; i++) {
        let value = values[i];
        let valueLower = value.toLowerCase();
        let parsed: Value;

        if (!isNaN(+value)) {
          parsed = +value;
        } else if (valueLower === "true" || valueLower === "false") {
          parsed = valueLower === "true";
        } else if (["null", "undefined", ""].includes(valueLower)) {
          parsed = null;
        } else {
          parsed = value;
        }

        entry[this.header[i]] = parsed;
      }

      data.push(entry as Entry<T>);
    });

    return new Promise<Entry<T>[]>((res, rej) => {
      stream.on("close", () => res(data));
      stream.on("error", err => rej(err));
    });
  }

  public async write(data: Entry<T>[]) {
    let stream = createWriteStream(this.src, { flags: "a" });

    for (let i = 0; i < data.length; i++) {
      let entry = data[i] as Record<string, unknown>;
      let values = this.header.map(header => entry[header]).join(DELIMITER);

      stream.write(values + "\n");
    }

    stream.end();

    return new Promise<void>((res, rej) => {
      stream.on("finish", () => res());
      stream.on("error", err => rej(err));
    });
  }
}

export { CSV };
