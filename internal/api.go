package shapefiletosvg

import (
	"fmt"
	"reflect"

	shp "github.com/jonas-p/go-shp"
)

// Convert opens a shapefile and converts it to an svg
func Convert(shapefileName string) error {
	shape, err := shp.Open(shapefileName)
	if err != nil {
		return err
	}
	defer shape.Close()
	fields := shape.Fields()
	for shape.Next() {
		n, p := shape.Shape()
		fmt.Println(reflect.TypeOf(p).Elem(), p.BBox())
		for k, f := range fields {
			val := shape.ReadAttribute(n, k)
			fmt.Printf("\t%v: %v\n", f, val)
		}
		fmt.Println()
	}
	return nil
}
