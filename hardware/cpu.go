package hardware

type CPU struct {
	acc        byte
	isar       ISAR
	scratchpad Scratchpad
	status     StatusFlags
	opcode     byte
	running    bool
	vram       *VRAM
	program    *Program
}

type ISAR struct {
	lo byte
	hi byte
}

type StatusFlags struct {
	icb      bool
	overflow bool
	zero     bool
	carry    bool
	sign     bool
}

type Scratchpad = [64]byte

func NewCPU() *CPU {
	c := new(CPU)
	c.acc = 0
	c.isar.Unpack(0)
	c.status.Unpack(0)
	c.vram = NewVRAM()
	c.program = nil
	return c
}

// TODO: Try to figure out actual reset and power-on state for the registers and memory
func (c *CPU) Reset() {
	c.acc = 0
	c.isar.Unpack(0)
	c.status.Unpack(0)
	c.opcode = 0
	c.scratchpad.Clear()
	c.vram.Clear()
	c.program.Seek(0)
}

// Execute the instruction for the given opcode
func (c *CPU) performInstruction(opcode byte) int {
	cycles := 1

	switch {
	// LR A, Ku
	case opcode == 0x00:

	// LR A, Kl
	case opcode == 0x01:

	// LR A, Qu
	case opcode == 0x02:

	// LR A, Ql
	case opcode == 0x03:

	// LR Ku, A
	case opcode == 0x04:

	// LR Kl, A
	case opcode == 0x05:

	// LR Qu, A
	case opcode == 0x06:

	// LR Ql, A
	case opcode == 0x07:

	// LR K, P
	case opcode == 0x08:

	// LR P, K
	case opcode == 0x09:

	// LR A, IS
	case opcode == 0x0A:

	// LR IS, A
	case opcode == 0x0B:

	// PK
	case opcode == 0x0C:

	// LR P0, Q
	case opcode == 0x0D:

	// LR Q, DC
	case opcode == 0x0E:

	// LR DC, Q
	case opcode == 0x0F:

	// LR DC, H
	case opcode == 0x10:

	// LR H, DC
	case opcode == 0x11:

	// SR 1
	case opcode == 0x12:

	// SL 1
	case opcode == 0x13:

	// SR 4
	case opcode == 0x14:

	// SL 4
	case opcode == 0x15:

	// LM
	case opcode == 0x16:

	// ST
	case opcode == 0x17:

	// COM
	case opcode == 0x18:

	// LNK
	case opcode == 0x19:

	// DI
	case opcode == 0x1A:

	// EI
	case opcode == 0x1B:

	// POP
	case opcode == 0x1C:

	// LR W, J
	case opcode == 0x1D:

	// LR J, W
	case opcode == 0x1E:

	// INC
	case opcode == 0x1F:

	// LI n
	case opcode == 0x20:

	// NI n
	case opcode == 0x21:

	// OI n
	case opcode == 0x22:

	// XI n
	case opcode == 0x23:

	// AI n
	case opcode == 0x24:

	// CI n
	case opcode == 0x25:

	// IN n
	case opcode == 0x26:

	// OUT n
	case opcode == 0x27:

	// PI mn
	case opcode == 0x28:

	// JMP mn
	case opcode == 0x29:

	// DCI mn
	case opcode == 0x2A:

	// NOP
	case opcode == 0x2B:

	// XDC
	case opcode == 0x2C:

	// DS r
	case opcode&0xF0 == 0x30:

	// LR A, r
	case opcode&0xF0 == 0x40:

	// LR r, A
	case opcode&0xF0 == 0x50:

	// LISU i
	case opcode&0xF8 == 0x60:

	// LISL i
	case opcode&0xF8 == 0x68:

	// CLR
	case opcode == 0x70:

	// LIS i
	case opcode&0xF0 == 0x70:

	// BT t, n
	// 001 = S
	// 010 = C
	// 100 = Z
	case opcode&0xF8 == 0x80:

	// AM
	case opcode == 0x88:

	// AMD
	case opcode == 0x89:

	// NM
	case opcode == 0x8A:

	// OM
	case opcode == 0x8B:

	// XM
	case opcode == 0x8C:

	// CM
	case opcode == 0x8D:

	// ADC
	case opcode == 0x8E:

	// BR7 n
	case opcode == 0x8F:

	// BR n
	case opcode == 0x90:

	// BF i, n
	// 0001 = S
	// 0010 = C
	// 0100 = Z
	// 1000 = O
	case opcode&0xF0 == 0x90:

	// INS i
	case opcode&0xF0 == 0xA0:

	// OUTS i
	case opcode&0xF0 == 0xB0:

	// AS r
	case opcode&0xF0 == 0xC0:

	// ASD r
	case opcode&0xF0 == 0xD0:

	// XS r
	case opcode&0xF0 == 0xE0:

	// NS r
	case opcode&0xF0 == 0xF0:
	}

	return cycles
}

// Perform one logical CPU step
func (c *CPU) Step() {
	// Fetch an opcode
	c.opcode = c.Program.ReadByte()

	// Execute the instruction
	c.cycles += performInstruction(c.opcode)

	// TODO: Interrupts etc
}

// Load a program into the CPU and run it forever
func (c *CPU) Execute(program []byte) {
	c.program = NewProgram(program)
	c.running = true

	for c.running {
		// TODO: this should be stepping by a certain number of cycles so we can limit the speed
		c.Step()
	}
}

// Get ISAR hi and lo as one packed value
func (i ISAR) Pack() byte {
	// TODO
	return 0
}

// Set ISAR hi and lo from one packed value
func (i ISAR) Unpack(val byte) {
	// TODO
}

// Get status flags as one packed value
func (s StatusFlags) Pack() byte {
	// TODO
	return 0
}

// Set status flags from one packed value
func (s StatusFlags) Unpack(val byte) {
	// TODO
}

// Enable a particular flag
func (s StatusFlag) Set(flag int) {

}

// Clear a particular flag
func (s StatusFlag) Clear(flag int) {

}

// Set status flags as one packed value

// Clear all scratchpad registers
func (s Scratchpad) Clear() {
	for i := 0; i < len(s); i++ {
		s[i] = 0
	}
}
