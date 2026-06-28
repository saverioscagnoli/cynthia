package util

import (
	"bytes"
	"image"
	"image/png"
)

func ScaleSprite(data []byte, factor float64) ([]byte, error) {
	src, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	b := src.Bounds()
	newW := int(float64(b.Dx()) * factor)
	newH := int(float64(b.Dy()) * factor)
	dst := image.NewRGBA(image.Rect(0, 0, newW, newH))

	scaleX := b.Dx()
	scaleY := b.Dy()

	for y := 0; y < newH; y++ {
		srcY := (y * scaleY) / newH
		for x := 0; x < newW; x++ {
			srcX := (x * scaleX) / newW
			dst.Set(x, y, src.At(b.Min.X+srcX, b.Min.Y+srcY))
		}
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, dst); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
