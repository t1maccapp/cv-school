package main

import (
	"io/ioutil"
	"path"
)

const SUPPORTED_IMAGE_EXT = ".png"
const EXT_LEN = 4

type ImageInfo struct {
	name       string
	path       string
	annotation *Annotation
}

func (image *ImageInfo) Name() string {
	return image.name
}

func (image *ImageInfo) Path() string {
	return image.path
}

func (image *ImageInfo) SetAnnotation(annotation *Annotation) {
	image.annotation = annotation
}

func (image *ImageInfo) Annotation() *Annotation {
	return image.annotation
}

func getImagesInfo(dir string) ([]*ImageInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	paths := make([]*ImageInfo, 0, len(files))

	for _, fileInfo := range files {
		if imageNameIsValid(fileInfo.Name()) {
			image := &ImageInfo{name: fileInfo.Name(), path: path.Join(dir, fileInfo.Name())}

			paths = append(paths, image)
		}
	}

	return paths, nil
}

func imageNameIsValid(imageName string) bool {
	extLenIsValid := len(imageName) > EXT_LEN
	extIsSupported := SUPPORTED_IMAGE_EXT == imageName[len(imageName)-EXT_LEN:]

	return extLenIsValid && extIsSupported
}
