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

	for y := 0; y < newH; y++ {
		srcY := int(float64(y) / factor)

		for x := 0; x < newW; x++ {
			srcX := int(float64(x) / factor)
			dst.Set(x, y, src.At(srcX, srcY))
		}
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, dst); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
