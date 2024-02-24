import { CSV } from "../dist";
import path from "path";
import { test, expect, afterAll } from "vitest";
import { faker } from "@faker-js/faker";
import { describe } from "node:test";

describe("csv", () => {
  test("read", async () => {
    let csv = new CSV({
      src: path.join(__dirname, "read.csv"),
      header: ["s:name", "n:age", "s:sex", "s:hair_color", "s:zodiac_sign"]
    });

    let data = Array.from({ length: 100 }, () => ({
      name: faker.person.firstName(),
      sex: faker.person.sex(),
      hair_color: faker.color.rgb(),
      age: faker.number.int({ min: 1, max: 100 }),
      zodiac_sign: faker.person.zodiacSign()
    }));

    await csv.write(data);

    let read = await csv.read();

    expect(read).toEqual(data);
  });

  test("basic write", async () => {
    let csv = new CSV({
      src: path.join(__dirname, "basic-write.csv"),
      header: ["s:name", "n:age", "s:sex", "s:hair_color", "s:zodiac_sign"]
    });

    let data = Array.from({ length: 100 }, () => ({
      name: faker.person.firstName(),
      sex: faker.person.sex(),
      hair_color: faker.color.rgb(),
      age: faker.number.int({ min: 1, max: 100 }),
      zodiac_sign: faker.person.zodiacSign()
    }));

    await csv.write(data);

    let read = await csv.read();

    expect(read).toEqual(data);
  });

  test("write values with commas", async () => {
    let csv = new CSV({
      src: path.join(__dirname, "write-commas.csv"),
      header: ["s:name", "n:age"]
    });

    let data = [
      { name: "John, Doe", age: 20 },
      { name: "Jane, Doe", age: 20 }
    ];

    await csv.write(data);

    let read = await csv.read();

    expect(read).toEqual(data);
  });

  test("write values with quotes", async () => {
    let csv = new CSV({
      src: path.join(__dirname, "write-quotes.csv"),
      header: ["s:name", "n:age"]
    });

    let data = [
      { name: 'John "Doe"', age: 20 },
      { name: 'Jane "Doe"', age: 20 }
    ];

    await csv.write(data);

    let read = await csv.read();

    expect(read).toEqual(data);
  });

  test(
    "write large",
    async () => {
      let csv = new CSV({
        src: path.join(__dirname, "write-large.csv"),
        header: ["s:name", "n:age", "s:sex", "s:hair_color", "s:zodiac_sign"]
      });

      let data = Array.from({ length: 1000000 }, () => ({
        name: faker.person.firstName(),
        sex: faker.person.sex(),
        hair_color: faker.color.rgb(),
        age: faker.number.int({ min: 1, max: 100 }),
        zodiac_sign: faker.person.zodiacSign()
      }));

      await csv.write(data);

      let data2 = Array.from({ length: 1000000 }, () => ({
        name: faker.person.firstName(),
        sex: faker.person.sex(),
        hair_color: faker.color.rgb(),
        age: faker.number.int({ min: 1, max: 100 }),
        zodiac_sign: faker.person.zodiacSign()
      }));

      await csv.write(data2);

      let read = await csv.read();

      expect(read).toEqual(data.concat(data2));
    },
    { timeout: -1 }
  );

  afterAll(() => {
    let files = [
      "read.csv",
      "basic-write.csv",
      "write-large.csv",
      "write-commas.csv",
      "write-quotes.csv"
    ].map(file => path.join(__dirname, file));

    for (let file of files) {
      require("fs").unlinkSync(file);
    }
  });
});
