package main

import (
	"flag"
	"log"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	// Parse command-line arguments
	outputFilename := flag.String("o", os.Args[1]+".png", "Output filename")
	size := flag.Int("s", 256, "QR code size (in pixels)")
	flag.Parse()

	// Generate QR code
	data := flag.Arg(0)
	if data == "" {
		log.Fatal("No data specified")
	}

	err := qrcode.WriteFile(data, qrcode.Medium, *size, *outputFilename)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("QR code written to %s", *outputFilename)
}
