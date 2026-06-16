package assets

import "embed"

//go:embed types/*.png
var TypeSprites embed.FS

//go:embed trainers/*.png
var TrainerSprites embed.FS

//go:embed trainer-sheet.json
var TrainerSheetJSON embed.FS

//go:embed trainer-sheet.png
var TrainerSheet embed.FS
