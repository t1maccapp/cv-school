package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Annotation struct {
	FragmentCoordinates []*FragmentCoordinates
}

func readAnnotationFromFile(path string) (*Annotation, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)

	fragmentCoordinates := make([]*FragmentCoordinates, 0)

	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		fc, err := getFragmentCoordinatesFromString(line)
		if err != nil {
			return nil, err
		}
		fragmentCoordinates = append(fragmentCoordinates, fc)
	}

	if err != io.EOF {
		return nil, err
	}

	annotation := &Annotation{
		FragmentCoordinates: fragmentCoordinates,
	}

	return annotation, nil
}

func getFragmentCoordinatesFromString(s string) (*FragmentCoordinates, error) {
	coordinates := strings.Split(s, ",")

	for i, c := range coordinates {
		coordinates[i] = strings.TrimSuffix(strings.TrimSuffix(c, "\n"), "\r")
	}

	xLeft, err := strconv.Atoi(coordinates[0])
	if err != nil {
		return nil, err
	}

	yLeft, err := strconv.Atoi(coordinates[1])
	if err != nil {
		return nil, err
	}

	xRight, err := strconv.Atoi(coordinates[2])
	if err != nil {
		return nil, err
	}

	yRight, err := strconv.Atoi(coordinates[3])
	if err != nil {
		return nil, err
	}

	fc := &FragmentCoordinates{
		XLeft:  xLeft,
		YLeft:  yLeft,
		XRight: xRight,
		YRight: yRight,
	}

	return fc, nil
}
