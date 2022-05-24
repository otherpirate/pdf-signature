package main

import (
	"flag"
	"fmt"

	"github.com/otherpirate/pdf-signature/internal/signature"
)

func main() {
	fileImage := flag.String("f", "", "path with filename for base image file")
	fileName := flag.String("n", "", "path with filename for all names should be signed")
	folderResult := flag.String("o", "", "path to directory where is gonna by all signed files")
	posX := flag.Float64("x", -1, "X position where should be has the name")
	posY := flag.Float64("y", -1, "Y position where should be has the name")
	flag.Parse()

	fmt.Printf("%s - %s - %s - %f - %f", *fileImage, *fileName, *folderResult, *posX, *posY)
	err := signature.Magic4(*fileImage, *folderResult, *posX, *posY)
	if err != nil {
		panic(err)
	}
}
