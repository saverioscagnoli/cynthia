package assets

import "embed"

//go:embed types/*.png
var TypeSprites embed.FS

//go:embed trainers/*.png
var TrainerSprites embed.FS
