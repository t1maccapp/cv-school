# cv-school cli app
[![Go Report Card](https://goreportcard.com/badge/github.com/t1maccapp/cv-school)](https://goreportcard.com/report/github.com/t1maccapp/cv-school)

Cli app for extracting fragments with people from images

### Installing

On OS X:

```
brew install go dep
git clone `this repo`
dep ensure
go build -o executable-bin `app folder`
```

### Params
* **-images, -i** -- Images folder
* **-annotations, -a** -- Annotations folder
* **-out, -o** -- Folder for results
* **-flip-horizontally, -f** -- Flip images horizontally
* **-grayscale, -g** -- Grayscale
* **-normalized-grayscale, -ng** -- Normalise grayscale intensity (only works with **-grayscale**)
* **-noise, -n** -- Add Gaussian noise

### Examples

```
executable-bin -i ./images -a ./annotations -o ./fragments -f -g -n
```

Check **./out** for sample outputs