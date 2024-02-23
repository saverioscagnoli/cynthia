import { DELIMITER } from "@consts";
import { Entry } from "@types";
import {
  existsSync,
  unlinkSync,
  writeFileSync,
  PathLike,
  createWriteStream
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
    this.header = header;

    this.init(src, !!deletePrevious);
  }

  private init(src: PathLike, deletePrevious: boolean) {
    if (existsSync(src) && deletePrevious) {
      unlinkSync(src);
    }

    writeFileSync(src, this.header.join(DELIMITER) + "\n");
  }

  public async write(data: Entry<T>[]) {
    let stream = createWriteStream(this.src, { flags: "a" });

    for (let i = 0; i < data.length; i++) {
      let entry = data[i];
      let values = this.header.map(header => entry[header as unknown as keyof Entry<T>]).join(DELIMITER);
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
