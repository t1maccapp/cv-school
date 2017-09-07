package main

import (
	"github.com/disintegration/imaging"
	"image"
)

// TODO: Add this value to cli params
const BLUR_SIGMA = 3.5

func convertToGrayscale(image *image.NRGBA) *image.NRGBA{
	return imaging.Grayscale(image)
}

func flipHorizontally(image *image.NRGBA) *image.NRGBA {
	return imaging.FlipH(image)
}

func blur(image *image.NRGBA) *image.NRGBA {
	return imaging.Blur(image, BLUR_SIGMA)
}