package test

import (
	"github.com/Bo887/chip8-emulator/chip8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test00E0Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	for i := range cpu.Display {
		cpu.Display[i] = 32
	}

	for _, elem := range cpu.Display {
		assert.Equal(t, uint8(32), elem)
	}

	cpu.Opcode = 0x00E0
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	for _, elem := range cpu.Display {
		assert.Equal(t, uint8(0), elem)
	}
}

func Test00EEOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.SP = 0
	cpu.Opcode = 0x00EE
	err := cpu.HandleOpcode()

	assert.NotNil(t, err)

	cpu.SP = 1
	cpu.Stack[0] = 0xBEEF
	next_instruction := uint16(0xBEEF + 2)
	cpu.PC = 0xDEAD

	assert.Equal(t, uint16(0xDEAD), cpu.PC)

	cpu.Opcode = 0x00EE
	err = cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, next_instruction, cpu.PC)
	assert.Equal(t, uint16(0), cpu.SP)
}

func TestInvalid00Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()

	cpu.Opcode = 0x00CC
	err := cpu.HandleOpcode()

	assert.NotNil(t, err)
}

func Test1NNNOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xDEAD

	cpu.Opcode = 0x1AAA
	err := cpu.HandleOpcode()
	assert.Nil(t, err)

	assert.Equal(t, uint16(0x0AAA), cpu.PC)
	assert.Equal(t, uint16(0), cpu.SP)
}

func Test2NNNOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	cpu.SP = 3

	cpu.Opcode = 0x2FFF
	err := cpu.HandleOpcode()
	assert.Nil(t, err)

	assert.Equal(t, uint16(0xFFF), cpu.PC)
	assert.Equal(t, uint16(4), cpu.SP)
	assert.Equal(t, uint16(0xBEEF), cpu.Stack[3])
}

func Test2NNNStackOverflow(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	cpu.SP = 15

	cpu.Opcode = 0x2FFF
	err := cpu.HandleOpcode()

	assert.Nil(t, err)

	cpu.Opcode = 0x2ABC
	err = cpu.HandleOpcode()

	assert.NotNil(t, err)
}

func Test3XNNOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xF00D
	cpu.V[0xF] = 0x42
	target_pc := uint16(0xF00D + 4)

	cpu.Opcode = 0x3F42
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)

	cpu.V[0x0] = 0x00
	target_pc += 2

	cpu.Opcode = 0x3012
	err = cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
}

func Test4XNNOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xF00D
	cpu.V[0xF] = 0x42
	target_pc := uint16(0xF00D + 2)

	cpu.Opcode = 0x4F42
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)

	cpu.V[0x0] = 0x00
	target_pc += 4

	cpu.Opcode = 0x4012
	err = cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
}

func Test5XY0Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	cpu.V[0x3] = 0xAC
	cpu.V[0x6] = 0xAC
	target_pc := uint16(0xBEEF) + 4

	cpu.Opcode = 0x5360
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)

	cpu.V[0x3] = 0xAD
	target_pc += 2

	cpu.Opcode = 0x5360
	err = cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
}

func TestInvalid5XY0Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	target_pc := uint16(0xBEEF)

	cpu.Opcode = 0x5FF1
	err := cpu.HandleOpcode()

	assert.NotNil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
}

func Test6XNNOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	target_pc := uint16(0xBEEF) + 2
	cpu.V[5] = 0x00

	assert.Equal(t, uint8(0x00), cpu.V[5])

	cpu.Opcode = 0x65FF
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, uint8(0xFF), cpu.V[5])
	assert.Equal(t, target_pc, cpu.PC)
}

func Test7XNNOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	cpu.V[5] = 0x23
	target_pc := uint16(0xBEEF) + 2
	target_register_value := uint8(0x23 + 0xAB)

	assert.Equal(t, cpu.V[5], uint8(0x23))

	cpu.Opcode = 0x75AB
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_register_value, cpu.V[5])
	assert.Equal(t, target_pc, cpu.PC)
}

func Test8XY0Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	cpu.V[5] = 0x00
	cpu.V[6] = 0x20
	target_pc := uint16(0xBEEF + 2)
	target_register_value := uint8(0x20)

	cpu.Opcode = 0x8560
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_register_value, cpu.V[5])
}

func Test8XY1Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	cpu.V[5] = 0x24
	cpu.V[6] = 0xEF
	target_pc := uint16(0xBEEF + 2)
	target_register_value := uint8(0x24 | 0xEF)

	cpu.Opcode = 0x8561
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_register_value, cpu.V[5])
}

func Test8XY2Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	cpu.V[5] = 0x24
	cpu.V[6] = 0xEF
	target_pc := uint16(0xBEEF + 2)
	target_register_value := uint8(0x24 & 0xEF)

	cpu.Opcode = 0x8562
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_register_value, cpu.V[5])
}

func Test8XY3Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	cpu.V[5] = 0x0D
	cpu.V[6] = 0xAB
	target_pc := uint16(0xBEEF + 2)
	target_register_value := uint8(0x0D ^ 0xAB)

	cpu.Opcode = 0x8563
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_register_value, cpu.V[5])
}

func Test8XY4Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	cpu.V[5] = 0x0A
	cpu.V[6] = 0x02
	target_pc := uint16(0xBEEF + 2)
	target_register_value := uint8(0x0A + 0x02)
	target_vf := uint8(0)

	cpu.Opcode = 0x8564
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_register_value, cpu.V[5])
	assert.Equal(t, target_vf, cpu.V[0xF])

	cpu.V[5] = 0xFF
	cpu.V[6] = 0x01
	target_pc += 2
	target_register_value = uint8(0x00)
	target_vf = 1

	cpu.Opcode = 0x8564
	err = cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_register_value, cpu.V[5])
	assert.Equal(t, target_vf, cpu.V[0xF])
}

func Test8XY5Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	cpu.V[5] = 0x0A
	cpu.V[6] = 0x02
	target_pc := uint16(0xBEEF + 2)
	target_register_value := uint8(0x0A - 0x02)
	target_vf := uint8(1)

	cpu.Opcode = 0x8565
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_register_value, cpu.V[5])
	assert.Equal(t, target_vf, cpu.V[0xF])

	cpu.V[5] = 0x00
	cpu.V[6] = 0x01
	target_pc += 2
	target_register_value = uint8(0xFF)
	target_vf = 0

	cpu.Opcode = 0x8565
	err = cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_register_value, cpu.V[5])
	assert.Equal(t, target_vf, cpu.V[0xF])
}

func Test8XY6Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xDEAD
	cpu.V[8] = 0x04
	target_pc := uint16(0xDEAD + 2)
	target_register_value := uint8(0x02)
	target_vf := uint8(0)

	cpu.Opcode = 0x8866
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_register_value, cpu.V[8])
	assert.Equal(t, target_vf, cpu.V[0xF])

	cpu.V[8] = 0x05
	target_pc += 2
	target_vf = uint8(1)

	cpu.Opcode = 0x8866
	err = cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_register_value, cpu.V[8])
	assert.Equal(t, target_vf, cpu.V[0xF])
}

func Test8XY7Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	cpu.V[5] = 0x02
	cpu.V[6] = 0x0A
	target_pc := uint16(0xBEEF + 2)
	target_register_value := uint8(0x0A - 0x02)
	target_vf := uint8(1)

	cpu.Opcode = 0x8567
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_register_value, cpu.V[5])
	assert.Equal(t, target_vf, cpu.V[0xF])

	cpu.V[5] = 0x01
	cpu.V[6] = 0x00
	target_pc += 2
	target_register_value = uint8(0xFF)
	target_vf = 0

	cpu.Opcode = 0x8567
	err = cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_register_value, cpu.V[5])
	assert.Equal(t, target_vf, cpu.V[0xF])
}

func Test8XYEOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xDEAD
	cpu.V[8] = 0x04
	target_pc := uint16(0xDEAD + 2)
	target_register_value := uint8(0x08)
	target_vf := uint8(0)

	cpu.Opcode = 0x886E
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_register_value, cpu.V[8])
	assert.Equal(t, target_vf, cpu.V[0xF])

	cpu.V[8] = 0x05
	target_pc += 2
	target_register_value = uint8(0x0A)
	target_vf = uint8(1)

	cpu.Opcode = 0x886E
	err = cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_register_value, cpu.V[8])
	assert.Equal(t, target_vf, cpu.V[0xF])
}

func TestInvalid8XYOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xDEAD
	target_pc := uint16(0xDEAD)

	cpu.Opcode = 0x8FFF
	err := cpu.HandleOpcode()

	assert.NotNil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
}

func Test9XY0Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	cpu.V[0x3] = 0xAC
	cpu.V[0x6] = 0xAC
	target_pc := uint16(0xBEEF) + 2

	cpu.Opcode = 0x9360
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)

	cpu.V[0x3] = 0xAD
	target_pc += 4

	cpu.Opcode = 0x9360
	err = cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
}

func TestANNNOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	target_pc := uint16(0xBEEF) + 2
	target_I := uint16(0xABC)

	cpu.Opcode = 0xAABC
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_I, cpu.I)
}

func TestBNNNOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	cpu.V[0] = 0xAB
	target_pc := uint16(0xAB + 0xCDE)

	cpu.Opcode = 0xBCDE
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
}

func TestCXNNOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xCAFE
	expected := uint8(250)
	target_pc := uint16(0xCAFE + 2)

	//test with preset seed twice: should be same values
	cpu.Opcode = 0xC5FF
	err := cpu.HandleCXNNOpcode(0)
	assert.Nil(t, err)
	assert.Equal(t, expected, cpu.V[5])
	assert.Equal(t, target_pc, cpu.PC)

	cpu.Opcode = 0xC5FF
	target_pc += 2
	err = cpu.HandleCXNNOpcode(0)
	assert.Nil(t, err)
	assert.Equal(t, expected, cpu.V[5])
	assert.Equal(t, target_pc, cpu.PC)

	//test with random seed twice: should be different values
	cpu.Opcode = 0xC5FF
	target_pc += 2
	err = cpu.HandleOpcode()
	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	first_res := cpu.V[5]

	cpu.Opcode = 0xC5FF
	target_pc += 2
	err = cpu.HandleOpcode()
	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	second_res := cpu.V[5]

	assert.NotEqual(t, first_res, second_res)
}

func TestDXYNOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.I = 0x0BAD
	cpu.Memory[cpu.I] = 0x3C
	cpu.Memory[cpu.I+1] = 0xC3
	cpu.Memory[cpu.I+2] = 0xFF
	cpu.PC = 0xCAFE

	cpu.Opcode = 0xD003
	err := cpu.HandleOpcode()
	assert.Nil(t, err)
	assert.Equal(t, uint8(0), cpu.V[0xF])
	for i := 0; i < 8; i++ {
		if i < 2 || i >= 6 {
			assert.Equal(t, cpu.Display[i], uint8(0))
			assert.Equal(t, cpu.Display[chip8.NUM_COLS+i], uint8(1))
		} else {
			assert.Equal(t, cpu.Display[i], uint8(1))
			assert.Equal(t, cpu.Display[chip8.NUM_COLS+i], uint8(0))
		}
	}
	for i := 0; i < 8; i++ {
		assert.Equal(t, cpu.Display[chip8.NUM_COLS*2+i], uint8(1))
	}

	//test for a collision
	cpu.Opcode = 0xD113
	err = cpu.HandleOpcode()
	assert.Nil(t, err)
	assert.Equal(t, uint8(1), cpu.V[0xF])
}

func TestEX9EOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xDEAD
	cpu.V[4] = 0xA
	cpu.Keypad[0xA] = 1
	target_pc := uint16(0xDEAD + 4)

	cpu.Opcode = 0xE49E
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)

	cpu.Keypad[0xA] = 0
	target_pc += 2

	cpu.Opcode = 0xE49E
	err = cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
}

func TestEXA1Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xDEAD
	cpu.V[4] = 0xA
	cpu.Keypad[0xA] = 1
	target_pc := uint16(0xDEAD + 2)

	cpu.Opcode = 0xE4A1
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)

	cpu.Keypad[0xA] = 0
	target_pc += 4

	cpu.Opcode = 0xE4A1
	err = cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
}

func TestInvalidEXOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xDEAD
	target_pc := uint16(0xDEAD)

	cpu.Opcode = 0xEF24
	err := cpu.HandleOpcode()

	assert.NotNil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
}

func TestFX07Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xDEAD
	cpu.DelayTimer = 0xFC
	target_pc := uint16(0xDEAD + 2)
	target_vx := uint8(0xFC)

	assert.Equal(t, uint8(0), cpu.V[3])
	cpu.Opcode = 0xF307
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_vx, cpu.V[3])
}

func TestFX0AOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xDEAD
	target_pc := uint16(0xDEAD)

	cpu.Opcode = 0xF40A
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)

	cpu.Keypad[0xC] = 1
	target_pc += 2
	target_keypad_idx := uint8(0xC)

	cpu.Opcode = 0xF40A
	err = cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_keypad_idx, cpu.V[4])
}

func TestFX15Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xDEAD
	cpu.V[0xC] = 0xAB
	target_pc := uint16(0xDEAD + 2)
	target_delay_timer := uint8(0xAB)

	cpu.Opcode = 0xFC15
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_delay_timer, cpu.DelayTimer)
}

func TestFX18Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xDEAD
	cpu.V[0x2] = 0xFF
	target_pc := uint16(0xDEAD + 2)
	target_sound_timer := uint8(0xFF)

	cpu.Opcode = 0xF218
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_sound_timer, cpu.SoundTimer)
}

func TestFX1EOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	cpu.V[5] = 0x0A
	cpu.I = 0x02
	target_pc := uint16(0xBEEF + 2)
	target_i := uint16(0x0A + 0x02)
	target_vf := uint8(0)

	cpu.Opcode = 0xF51E
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_i, cpu.I)
	assert.Equal(t, target_vf, cpu.V[0xF])

	cpu.V[5] = 0x01
	cpu.I = 0xFFFF
	target_pc += 2
	target_i = uint16(0x00)
	target_vf = 1

	cpu.Opcode = 0xF51E
	err = cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_i, cpu.I)
	assert.Equal(t, target_vf, cpu.V[0xF])
}

func TestFX29Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	cpu.V[0] = 0xE
	target_i := uint16(0xE * 5)
	target_pc := uint16(0xBEEF + 2)

	cpu.Opcode = 0xF029
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, target_i, cpu.I)
}

func TestFX33Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xBEEF
	cpu.V[3] = 254
	cpu.I = 0x0034
	target_pc := uint16(0xBEEF + 2)

	cpu.Opcode = 0xF333
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	assert.Equal(t, uint8(2), cpu.Memory[cpu.I])
	assert.Equal(t, uint8(5), cpu.Memory[cpu.I+1])
	assert.Equal(t, uint8(4), cpu.Memory[cpu.I+2])

	cpu.I = 0xF00D

	cpu.Opcode = 0xF333
	err = cpu.HandleOpcode()

	assert.NotNil(t, err)
}

func TestFX55Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	for i := 0; i < chip8.NUM_REGISTERS; i++ {
		cpu.V[i] = uint8(i)
	}
	cpu.I = 0x0ABC
	cpu.PC = 0xF00D
	target_pc := uint16(0xF00D + 2)

	cpu.Opcode = 0xF455
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	for i := 0; i < chip8.NUM_REGISTERS; i++ {
		if i <= 4 {
			assert.Equal(t, uint8(i), cpu.Memory[cpu.I+uint16(i)])
		} else {
			assert.Equal(t, uint8(0), cpu.Memory[cpu.I+uint16(i)])
		}
	}

	cpu.I = chip8.MEMORY_SIZE - 8
	cpu.Opcode = 0xF955
	err = cpu.HandleOpcode()

	assert.NotNil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
}

func TestFX65Opcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	for i := 118; i < 150; i++ {
		cpu.Memory[i] = uint8(i)
	}
	cpu.I = 118
	cpu.PC = 0xF00D
	target_pc := uint16(0xF00D + 2)

	cpu.Opcode = 0xFC65
	err := cpu.HandleOpcode()

	assert.Nil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
	for i := 0; i < chip8.NUM_REGISTERS; i++ {
		if i <= 0xC {
			assert.Equal(t, uint8(i)+118, cpu.V[i])
		} else {
			assert.Equal(t, uint8(0), cpu.V[i])
		}
	}

	cpu.I = chip8.MEMORY_SIZE - 8
	cpu.Opcode = 0xF965
	err = cpu.HandleOpcode()

	assert.NotNil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
}

func TestInvalidFXOpcode(t *testing.T) {
	cpu := chip8.CreateCpu()
	cpu.PC = 0xDAD

	target_pc := uint16(0xDAD)

	cpu.Opcode = 0xFCAB
	err := cpu.HandleOpcode()

	assert.NotNil(t, err)
	assert.Equal(t, target_pc, cpu.PC)
}
