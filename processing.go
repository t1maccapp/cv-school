package main

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
)

// TODO: Add this value to cli params
const BLUR_SIGMA = 3.5

func convertToGrayscale(img *image.NRGBA) *image.NRGBA {
	return imaging.Grayscale(img)
}

func flipHorizontally(img *image.NRGBA) *image.NRGBA {
	return imaging.FlipH(img)
}

func blur(img *image.NRGBA) *image.NRGBA {
	return imaging.Blur(img, BLUR_SIGMA)
}

func normalizeGrayscaleIntensity(img *image.NRGBA) *image.NRGBA {
	min := uint8(255)
	max := uint8(0)

	for x := 0; x <= img.Rect.Max.X; x++ {
		for y := 0; y <= img.Rect.Max.Y; y++ {
			nrgba := img.NRGBAAt(x, y)

			intensity := getIntensityFromNRGBA(nrgba)

			if intensity < min {
				min = intensity
			}

			if intensity > max {
				max = intensity
			}
		}
	}

	fn := func(c color.NRGBA) color.NRGBA {
		i := getIntensityFromNRGBA(c)

		iN := (float64(i) - float64(min)) * (255 / (float64(max) - float64(min)))
		iNR := round(iN)

		return color.NRGBA{iNR, iNR, iNR, c.A}
	}

	return imaging.AdjustFunc(img, fn)
}

func getIntensityFromNRGBA(c color.NRGBA) uint8 {
	return round(0.299*float64(c.R) + 0.587*float64(c.G) + 0.114*float64(c.B))
}

func round(x float64) uint8 {
	if x < 0 {
		return uint8(x - 0.5)
	}

	return uint8(x + 0.5)
}
