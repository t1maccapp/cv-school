package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"path"
)

func validateInputFlags(c *cli.Context) error {
	if c.String("images") == "" {
		return errors.New("images (i) should be specified")
	}

	if c.String("annotations") == "" {
		return errors.New("annotations (a) should be specified")
	}

	if c.String("out") == "" {
		return errors.New("out (a) should be specified")
	}

	return nil
}

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "images, i",
			Usage: "Images dir path",
		},
		cli.StringFlag{
			Name:  "annotations, a",
			Usage: "Annotations dir path",
		},
		cli.StringFlag{
			Name:  "out, o",
			Usage: "Results dir path",
		},
		cli.BoolFlag{
			Name:  "grayscale, g",
			Usage: "Add Grayscaling to fragments",
		},
		cli.BoolFlag{
			Name:  "flip-horizontally, f",
			Usage: "Flip fragments",
		},
		cli.BoolFlag{
			Name:  "bright-normalization, b",
			Usage: "Flip fragments",
		},
		cli.BoolFlag{
			Name:  "noise, n",
			Usage: "Flip fragments",
		},
	}

	app.Action = func(c *cli.Context) {
		err := validateInputFlags(c)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		imagesDir := c.String("images")
		annotationsDir := c.String("annotations")
		resultsDir := c.String("out")

		imagesInfo, err := getImagesInfo(imagesDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, imageInfo := range imagesInfo {
			annotationName := imageInfo.Name()[:len(imageInfo.Name())-4] + ".txt"

			annotation, err := readAnnotationFromFile(path.Join(annotationsDir, annotationName))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			imageInfo.SetAnnotation(annotation)
		}

		processingOptions := &ProcessingOptions{
			grayscale:        c.Bool("grayscale"),
			flipHorizontally: c.Bool("flip-horizontally"),
			noise:            c.Bool("noise"),
		}

		for _, imageInfo := range imagesInfo {
			err := processImagesToFragments(imageInfo, processingOptions, resultsDir)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	app.Run(os.Args)
}
