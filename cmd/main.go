package main

import (
	"flag"

	"github.com/otherpirate/pdf-signature/internal/name"
	"github.com/otherpirate/pdf-signature/internal/signature"
)

func main() {
	fileImage := flag.String("f", "", "path with filename for base image file")
	fileName := flag.String("n", "", "path with filename for all names should be signed")
	folderResult := flag.String("o", "", "path to directory where is gonna by all signed files")
	posX := flag.Float64("x", -1, "X position where should be has the name")
	posY := flag.Float64("y", -1, "Y position where should be has the name")
	flag.Parse()

	names, err := name.GetNames(*fileName)
	if err != nil {
		panic(err)
	}
	err = signature.Signature(*fileImage, *folderResult, *posX, *posY, names)
	if err != nil {
		panic(err)
	}
}
