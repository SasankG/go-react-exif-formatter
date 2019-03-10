package util

import (
	"image/color"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
)

// testing image
var img = "./images/exif7.jpg"

// Open the image -> Change this to transform input images later
func openImg(imageFile string) *os.File {
	// open img
	file, err := os.Open(imageFile)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

// Opens file and gets exif data
func exifGet(imgs string) string {
	// Get the opened file
	file := openImg(imgs)

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

// Generate randome name
func nameGen(length int) string {
	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}

// Transform and save function
// @params imges -> the path of the image
func transform(imges string) {
	// Get orientation number
	exifNum := exifGet(imges)

	// open image
	myImage, err := imaging.Open(imges)
	if err != nil {
		log.Fatal(err)
	}

	// Generate a different sequence every time the function is run
	rand.Seed(time.Now().UnixNano())

	// conditionals and saves
	if exifNum == "3" {
		rotatedImg := imaging.Rotate(myImage, 180, color.NRGBA{0, 0, 0, 0})
		err = imaging.Save(rotatedImg, "./testingsave/"+nameGen(10)+".jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	} else if exifNum == "6" {
		rotatedImg := imaging.Rotate(myImage, 270, color.NRGBA{0, 0, 0, 0})
		err = imaging.Save(rotatedImg, "./testingsave/"+nameGen(10)+".jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	} else if exifNum == "8" {
		rotatedImg := imaging.Rotate(myImage, 90, color.NRGBA{0, 0, 0, 0})
		err = imaging.Save(rotatedImg, "./testingsave/"+nameGen(10)+".jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	} else if exifNum == "1" {
		rotatedImg := myImage
		err = imaging.Save(rotatedImg, "./testingsave/"+nameGen(10)+".jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	} else if exifNum == "7" {
		rotatedImg := imaging.Rotate(myImage, 90, color.NRGBA{0, 0, 0, 0})
		flippedImg := imaging.FlipH(rotatedImg)
		err = imaging.Save(flippedImg, "./testingsave/"+nameGen(10)+".jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	} else if exifNum == "2" {
		flippedImg := imaging.FlipH(myImage)
		err = imaging.Save(flippedImg, "./testingsave/"+nameGen(10)+".jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	} else if exifNum == "5" {
		rotatedImg := imaging.Rotate(myImage, 270, color.NRGBA{0, 0, 0, 0})
		flippedImg := imaging.FlipH(rotatedImg)
		err = imaging.Save(flippedImg, "./testingsave/"+nameGen(10)+".jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	} else if exifNum == "8" {
		flippedImg := imaging.FlipV(myImage)
		err = imaging.Save(flippedImg, "./testingsave/"+nameGen(10)+".jpg")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	}
}
