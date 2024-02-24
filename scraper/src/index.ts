import { randomUUID } from "crypto";
import { CSV } from "@csv";

import path from "path";

const csv = new CSV({
  src: path.join(__dirname, "read.csv"),
  header: ["s:name", "n:age"]
});

let data = Array.from({ length: 100 }, () => ({
  name: randomUUID(),
  age: Math.floor(Math.random() * 100)
}));

csv.write(data).then(async () => {
  let read = await csv.read();
  console.log(read);
});
