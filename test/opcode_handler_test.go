package test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/Bo887/chip8-emulator"
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

    assert.Equal(t, nil, err)
    for _, elem := range cpu.Display {
        assert.Equal(t, uint8(0), elem)
    }
}

func Test00EEOpcode(t *testing.T) {
    cpu := chip8.CreateCpu()
    cpu.SP = 1
    cpu.Stack[0] = 0xBEEF
    next_instruction := uint16(0xBEEF + 2)
    cpu.PC = 0xDEAD

    assert.Equal(t, uint16(0xDEAD), cpu.PC)

    cpu.Opcode = 0x00EE
    err := cpu.HandleOpcode()

    assert.Equal(t, nil, err)
    assert.Equal(t, next_instruction, cpu.PC)
    assert.Equal(t, uint16(0), cpu.SP)
}

func TestInvalid00Opcode(t *testing.T) {
    cpu := chip8.CreateCpu()

    cpu.Opcode = 0x00CC
    err := cpu.HandleOpcode()

    assert.NotEqual(t, nil, err)
}

func Test1NNNOpcode(t *testing.T) {
    cpu := chip8.CreateCpu()
    cpu.PC = 0xDEAD

    cpu.Opcode = 0x1AAA
    err := cpu.HandleOpcode()
    assert.Equal(t, nil, err)

    assert.Equal(t, uint16(0x0AAA), cpu.PC)
    assert.Equal(t, uint16(0), cpu.SP)
}

func Test2NNNOpcode(t *testing.T) {
    cpu := chip8.CreateCpu()
    cpu.PC = 0xBEEF
    cpu.SP = 3

    cpu.Opcode = 0x2FFF
    err := cpu.HandleOpcode()
    assert.Equal(t, nil, err)

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

    assert.Equal(t, nil, err)

    cpu.Opcode = 0x2ABC
    err = cpu.HandleOpcode()

    assert.NotEqual(t, nil, err)
}

func Test3XNNOpcode(t *testing.T) {
    cpu := chip8.CreateCpu()
    cpu.PC = 0xF00D
    cpu.V[0xF] = 0x42
    target_pc := uint16(0xF00D + 4)

    cpu.Opcode = 0x3F42
    err := cpu.HandleOpcode()

    assert.Equal(t, nil, err)
    assert.Equal(t, target_pc, cpu.PC)

    cpu.V[0x0] = 0x00
    target_pc += 2

    cpu.Opcode = 0x3012
    err = cpu.HandleOpcode()

    assert.Equal(t, nil, err)
    assert.Equal(t, target_pc, cpu.PC)
}

func Test4XNNOpcode(t *testing.T) {
    cpu := chip8.CreateCpu()
    cpu.PC = 0xF00D
    cpu.V[0xF] = 0x42
    target_pc := uint16(0xF00D + 2)

    cpu.Opcode = 0x4F42
    err := cpu.HandleOpcode()

    assert.Equal(t, nil, err)
    assert.Equal(t, target_pc, cpu.PC)

    cpu.V[0x0] = 0x00
    target_pc += 4

    cpu.Opcode = 0x4012
    err = cpu.HandleOpcode()

    assert.Equal(t, nil, err)
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

    assert.Equal(t, nil, err)
    assert.Equal(t, target_pc, cpu.PC)

    cpu.V[0x3] = 0xAD
    target_pc += 2

    cpu.Opcode = 0x5360
    err = cpu.HandleOpcode()

    assert.Equal(t, nil, err)
    assert.Equal(t, target_pc, cpu.PC)
}

func TestInvalid5XY0Opcode(t *testing.T) {
    cpu := chip8.CreateCpu()
    cpu.PC = 0xBEEF
    target_pc := uint16(0xBEEF)

    cpu.Opcode = 0x5FF1
    err := cpu.HandleOpcode()

    assert.NotEqual(t, nil, err)
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

    assert.Equal(t, nil, err)
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

    assert.Equal(t, nil, err)
    assert.Equal(t, target_register_value, cpu.V[5])
    assert.Equal(t, target_pc, cpu.PC)
}
