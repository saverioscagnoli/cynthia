import { CSV } from "../dist";
import path from "path";
import { test } from "vitest";

test("csv", () => {
  const csv = new CSV({
    src: path.join(__dirname, "test.csv"),
    header: ["name", "age"]
  });

  csv.write([
    { age: 20, name: "John" },
    { name: "Doe", age: 30 }
  ]);
});
