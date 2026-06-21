package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type SpriteEntry struct {
	ID   int    `json:"id"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	W    int    `json:"w"`
	H    int    `json:"h"`
	Name string `json:"name"`
}

type SpriteSheet struct {
	Sheet   SpriteEntry            `json:"_sheet"`
	Sprites map[string]SpriteEntry `json:"sprites"`
}

func main() {
	dir := flag.String("dir", "./sprites", "directory of PNGs")
	out := flag.String("out", "./assets", "output directory")
	cols := flag.Int("cols", 40, "sprites per row")
	flag.Parse()

	entries, err := os.ReadDir(*dir)
	if err != nil {
		panic(err)
	}

	type loaded struct {
		name string
		img  image.Image
	}

	var sprites []loaded
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".png") {
			continue
		}
		f, err := os.Open(filepath.Join(*dir, e.Name()))
		if err != nil {
			panic(err)
		}
		img, err := png.Decode(f)
		f.Close()
		if err != nil {
			panic(err)
		}
		name := strings.TrimSuffix(e.Name(), ".png")
		sprites = append(sprites, loaded{name, img})
	}

	if len(sprites) == 0 {
		fmt.Println("no sprites found")
		return
	}

	sort.Slice(sprites, func(i, j int) bool {
		return sprites[i].name < sprites[j].name
	})

	cellW, cellH := 0, 0
	for _, s := range sprites {
		b := s.img.Bounds()
		if b.Dx() > cellW {
			cellW = b.Dx()
		}
		if b.Dy() > cellH {
			cellH = b.Dy()
		}
	}

	rows := (len(sprites) + *cols - 1) / *cols
	sheet := image.NewRGBA(image.Rect(0, 0, cellW**cols, cellH*rows))
	spriteMap := make(map[string]SpriteEntry)

	for i, s := range sprites {
		col := i % *cols
		row := i / *cols
		x, y := col*cellW, row*cellH
		dst := image.Rect(x, y, x+cellW, y+cellH)
		draw.Draw(sheet, dst, s.img, s.img.Bounds().Min, draw.Over)
		spriteMap[s.name] = SpriteEntry{ID: i, X: x, Y: y, W: cellW, H: cellH, Name: s.name}
	}

	result := SpriteSheet{
		Sheet: SpriteEntry{
			ID:   -1,
			X:    0,
			Y:    0,
			W:    sheet.Bounds().Dx(),
			H:    sheet.Bounds().Dy(),
			Name: "_sheet",
		},
		Sprites: spriteMap,
	}

	os.MkdirAll(*out, 0755)

	sheetFile, _ := os.Create(filepath.Join(*out, "trainer-sheet.png"))
	png.Encode(sheetFile, sheet)
	sheetFile.Close()

	mapFile, _ := os.Create(filepath.Join(*out, "trainer-sheet.json"))
	enc := json.NewEncoder(mapFile)
	enc.SetIndent("", "  ")
	enc.Encode(result)
	mapFile.Close()

	fmt.Printf("✓ %d sprites → trainer-sheet.png + trainer-sheet.json\n", len(sprites))
}
