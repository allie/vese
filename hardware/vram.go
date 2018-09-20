package hardware

const (
	VRAMWidth  = 128
	VRAMHeight = 64
)

// RGB pixel
type Pixel struct {
	r byte
	g byte
	b byte
}

type VRAM struct {
	data [2048]byte
}

func NewVRAM() *VRAM {
	v := new(VRAM)
	return v
}

// Zero out all VRAM data
func (v *VRAM) Clear() {
	for i := 0; i < len(v.data); i++ {
		v.data[i] = 0
	}
}

// Generate an SDL-friendly framebuffer from VRAM data
func (v *VRAM) GetPixelBuffer() []Pixel {
	var buffer []Pixel
	buffer = make([]Pixel, VRAMWidth*VRAMHeight)

	// Do some stuff

	return buffer
}
