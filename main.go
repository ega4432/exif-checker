package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rwcarlsen/goexif/exif"
)

var filepath string

func init() {
	flag.StringVar(&filepath, "file", "", "Specify the path of the file whose exif you want to check.")
	flag.Parse()
}

func main() {
	if filepath == "" && len(os.Args) < 2 {
		panic("filepath is empty")
	}

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	x, err := exif.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	camModel, _ := x.Get(exif.Model)
	str, err := camModel.StringVal()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	log.Println(str)

	focal, _ := x.Get(exif.FocalLength)
	numer, denom, _ := focal.Rat2(0) // retrieve first (only) rat. value
	fmt.Printf("%v/%v\n", numer, denom)

	// Two convenience functions exist for date/time taken and GPS coords:
	tm, _ := x.DateTime()
	fmt.Println("Taken: ", tm)

	lat, long, _ := x.LatLong()
	fmt.Println("lat, long: ", lat, ", ", long)
}
