package util

import (
	"image/color"
	"log"
	"os"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
)

var img = "./images/exif7.jpg"

// Open the image -> Change this to transform input images later
func openImg(img string) *os.File {
	// open img
	file, err := os.Open(img)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

// Opens file and gets exif data
func exifGet(img string) string {
	// Get the opened file
	file := openImg(img)

	// decode exif data
	x, err := exif.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// get orientation
	orientation, _ := x.Get(exif.Orientation)

	// to string
	sOrientation := orientation.String()
	return sOrientation
}

// Transform and save function
func transform(img string) {
	// Get orientation number
	exifNum := exifGet(img)

	// open image
	myImage, err := imaging.Open(img)
	if err != nil {
		log.Fatal(err)
	}

	// conditionals and saves
	if exifNum == "3" {
		rotatedImg := imaging.Rotate(myImage, 180, color.NRGBA{0, 0, 0, 0})
		err = imaging.Save(rotatedImg, "./testingsave/etcetcetc.jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	} else if exifNum == "6" {
		rotatedImg := imaging.Rotate(myImage, 270, color.NRGBA{0, 0, 0, 0})
		err = imaging.Save(rotatedImg, "./testingsave/etcetcetc.jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	} else if exifNum == "8" {
		rotatedImg := imaging.Rotate(myImage, 90, color.NRGBA{0, 0, 0, 0})
		err = imaging.Save(rotatedImg, "./testingsave/etcetcetc.jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	} else if exifNum == "1" {
		rotatedImg := myImage
		err = imaging.Save(rotatedImg, "./testingsave/etcetcetc.jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	} else if exifNum == "7" {
		rotatedImg := imaging.Rotate(myImage, 90, color.NRGBA{0, 0, 0, 0})
		flippedImg := imaging.FlipH(rotatedImg)
		err = imaging.Save(flippedImg, "./testingsave/etcetcetc.jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	} else if exifNum == "2" {
		flippedImg := imaging.FlipH(myImage)
		err = imaging.Save(flippedImg, "./testingsave/etcetcetc.jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	} else if exifNum == "5" {
		rotatedImg := imaging.Rotate(myImage, 270, color.NRGBA{0, 0, 0, 0})
		flippedImg := imaging.FlipH(rotatedImg)
		err = imaging.Save(flippedImg, "./testingsave/etcetcetc.jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	} else if exifNum == "8" {
		flippedImg := imaging.FlipV(myImage)
		err = imaging.Save(flippedImg, "./testingsave/etcetcetc.jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	}
}
