package ui

import (
	"github.com/allie/vese/hardware"
	"github.com/veandco/go-sdl2/sdl"
)

type video struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	windowX  int32
	windowY  int32
	width    int32
	height   int32
	title    string
}

func newVideo() *video {
	v := new(video)

	v.title = "vese"

	// TODO: replace these with config variables
	v.windowX = sdl.WINDOWPOS_UNDEFINED
	v.windowY = sdl.WINDOWPOS_UNDEFINED
	v.width = 512
	v.height = 256

	var err error

	v.window, err = sdl.CreateWindow(v.title, v.windowX, v.windowY, v.width, v.height, sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}

	v.renderer, err = sdl.CreateRenderer(v.window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		panic(err)
	}

	v.renderer.SetLogicalSize(hardware.VRAMWidth, hardware.VRAMHeight)

	return v
}

func (v *video) render(dt float64) {
	dt = dt
	v.renderer.SetDrawColor(0, 0, 0, 255)
	v.renderer.Clear()

	v.renderer.SetDrawColor(255, 0, 0, 255)
	rect := sdl.Rect{0, 0, 128, 64}
	v.renderer.FillRect(&rect)

	v.renderer.Present()
}

func (v *video) destroy() {
	v.renderer.Destroy()
	v.window.Destroy()
}
