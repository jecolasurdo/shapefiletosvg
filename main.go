package main

import (
	"log"

	shapefiletosvg "github.com/jecolasurdo/shapefiletosvg/internal"
)

func main() {
	var err error
	var config *Configuration
	if config, err = GetConfiguration(); err == nil {
		err = shapefiletosvg.Convert(config.ShapefileLocation)
	}

	if err != nil {
		log.Fatal(err)
	}
}
