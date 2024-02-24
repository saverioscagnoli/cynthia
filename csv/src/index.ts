import { DELIMITER } from "@consts";
import { Entry, Value } from "@types";
import csv from "csv-parser";
import {
  existsSync,
  unlinkSync,
  writeFileSync,
  PathLike,
  createWriteStream,
  createReadStream
} from "fs";

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

  /**
   * Function to initialize the CSV file.
   * @param src The source of the CSV file.
   * @param deletePrevious Whether to delete the file with the same source if it exists.
   * @private
   */
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

  /**
   * Function to read the CSV file.
   * @param options The options to use when reading the CSV file.
   * @returns A promise that resolves with the data from the CSV file.
   */
  public async read(options: { limit?: number } = {}) {
    let data = [] as Entry<T>[];
    let limit = options.limit || Infinity;

    return new Promise<Entry<T>[]>((res, rej) => {
      let first = false;

      createReadStream(this.src)
        .pipe(csv({ headers: this.header, separator: DELIMITER }))
        .on("data", row => {
          if (!first) {
            first = true;
            return;
          }

          data.push(row as Entry<T>);

          for (let k in row) {
            let value = row[k] as string;
            let parsed: Value = null;

            if (!isNaN(+value!)) {
              parsed = +value;
            } else if (value === "true") {
              parsed = true;
            } else if (value === "false") {
              parsed = false;
            } else if (value === "null") {
              parsed = null;
            } else {
              parsed = value;
            }

            row[k] = parsed;
          }

          if (data.length === limit) {
            res(data);
          }
        })
        .on("end", () => res(data))
        .on("error", err => rej(err));
    });
  }

  /**
   * Function to write to the CSV file.
   * @param data The data to write to the CSV file.
   * @returns A promise that resolves when the data has been written to the CSV file.
   */
  public async write(data: Entry<T>[]) {
    let stream = createWriteStream(this.src, { flags: "a" });

    for (let i = 0; i < data.length; i++) {
      let entry = data[i] as Record<string, unknown>;
      let values = this.header
        .map(header => {
          let value = String(entry[header]);
          return value.includes(",") ? `"${value}"` : value;
        })
        .join(DELIMITER);

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
