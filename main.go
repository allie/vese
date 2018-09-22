package main

import (
	"flag"
	"fmt"
	"github.com/allie/vese/hardware"
	"github.com/allie/vese/ui"
)

func main() {
	rom := flag.String("rom", "", "Path to ROM file")
	cfg := flag.String("config", "config.json", "Path to configuration file")

	flag.Parse()

	if *rom == "" {
		fmt.Println("Please specify a ROM file to emulate.")
		return
	}

	// TODO: parse the configuration file
	*cfg = *cfg

	cart, err := hardware.NewCartridge(*rom)

	if err != nil {
		panic(err)
	}

	fmt.Printf("ROM:   %s\n", cart.Title)
	fmt.Printf("Size:  %d\n", cart.Size)
	fmt.Printf("MD5:   %s\n", cart.Checksum)
	fmt.Printf("Known: %t\n", cart.Known)

	ves := hardware.NewVES()
	u := ui.NewUi()

	go ves.Emulate(cart)

	u.Run()
}
