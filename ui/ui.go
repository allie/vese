package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Ui struct {
	video   *Video
	running bool
	last    uint64
	now     uint64
	dt      float64
}

func NewUi() *Ui {
	u := new(Ui)

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	u.video = NewVideo()
	return u
}

func (u *Ui) Run() {
	u.running = true
	u.now = sdl.GetPerformanceCounter()
	u.last = 0

	for u.running {
		u.last = u.now
		u.now = sdl.GetPerformanceCounter()
		u.dt = (float64((u.now-u.last)*1000) / float64(sdl.GetPerformanceFrequency()))

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				u.running = false
				break
			}
		}

		u.video.Render(u.dt)
	}
}

func (u *Ui) Quit() {
	u.video.Destroy()
	sdl.Quit()
}
