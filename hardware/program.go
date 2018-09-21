package hardware

import "errors"

type Program struct {
	data []byte
	Pc   uint16
}

// Load a program and set the program counter to 0
func NewProgram(data []byte) *Program {
	p := new(Program)
	p.data = make([]byte, len(data))
	copy(data, p.data)
	p.Pc = 0
	return p
}

// Read one byte at the program counter, then increment the program counter
func (p *Program) ReadByte() (byte, error) {
	if int(p.Pc) >= len(p.data) {
		return 0, errors.New("Program data buffer is empty")
	}

	val := p.data[p.Pc]
	p.Pc++

	return val, nil
}

// Move the program counter to the specified address
func (p *Program) Seek(address uint16) error {
	if int(address) >= len(p.data) {
		return errors.New("Address out of bounds")
	}

	p.Pc = address
	return nil
}
