package main

import (
	"github.com/disintegration/imaging"
	"image"
	"path"
	"strconv"
)

type FragmentCoordinates struct {
	XLeft  int
	YLeft  int
	XRight int
	YRight int
}

type ProcessingOptions struct {
	grayscale           bool
	normalizedGrayscale bool
	flipHorizontally    bool
	noise               bool
}

func processImagesToFragments(img *ImageInfo, options *ProcessingOptions, resultsDir string) error {
	srcImg, err := imaging.Open(img.Path())
	if err != nil {
		return err
	}

	for i, fc := range img.Annotation().FragmentCoordinates {
		rect := image.Rect(fc.XLeft, fc.YLeft, fc.XRight, fc.YRight)

		fragment := imaging.Crop(srcImg, rect)

		if options.grayscale {
			fragment = convertToGrayscale(fragment)
			if options.normalizedGrayscale {
				fragment = normalizeGrayscaleIntensity(fragment)
			}
		}

		if options.flipHorizontally {
			fragment = flipHorizontally(fragment)
		}

		if options.noise {
			fragment = blur(fragment)
		}

		resultPath := path.Join(resultsDir, img.Name()[:len(img.Name())-4]+"_"+strconv.Itoa(i)+".png")

		err := imaging.Save(fragment, resultPath)
		if err != nil {
			return err
		}
	}

	return nil
}
