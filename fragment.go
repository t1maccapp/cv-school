package main

import (
	"github.com/disintegration/imaging"
	"image"
	"path"
	"strconv"
)

type FragmentCoordinates struct {
	XLeft int
	YLeft int
	XRight int
	YRight int
}

type ProcessingOptions struct {
	grayscale bool
	flipHorizontally bool
	noise bool
}

func processImagesToFragments(img *ImageInfo, options *ProcessingOptions, resultsDir string) error {
	srcImg, err := imaging.Open(img.Path())
	if err != nil {
		return err
	}

	for i, fc := range img.Annotation().FragmentCoordinates {
		rect := image.Rect(fc.XLeft, fc.YLeft, fc.XRight, fc.YRight)

		result := imaging.Crop(srcImg, rect)

		if options.grayscale {
			result = convertToGrayscale(result)
		}

		if options.flipHorizontally {
			result = flipHorizontally(result)
		}

		if options.noise {
			result = blur(result)
		}

		resultPath := path.Join(resultsDir, img.Name()[:len(img.Name())-4] + "_" + strconv.Itoa(i) + ".png")

		err := imaging.Save(result, resultPath)
		if err != nil {
			return err
		}
	}

	return nil
}