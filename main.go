package main

import (
	"github.com/allie/vese/hardware"
	"github.com/allie/vese/ui"
)

func main() {
	ves := hardware.NewVES()
	u := ui.NewUi()
	go ves.Emulate(make([]byte, 1))
	u.Run()
}
