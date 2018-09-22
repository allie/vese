package hardware

type VES struct {
	cpu *CPU
}

func NewVES() *VES {
	v := new(VES)
	v.cpu = NewCPU()
	return v
}

func (v *VES) Emulate(program []byte) {
	v.cpu.Execute(program)
}
