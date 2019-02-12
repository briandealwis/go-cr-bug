package main

import (
	"fmt"

	"github.com/google/go-containerregistry/pkg/v1/tarball"
)

func main() {
	img, err := tarball.ImageFromPath("basic_windows_image.tar", nil)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	} else {
		fmt.Printf("Opened image: %v\n", img)
	}
}
