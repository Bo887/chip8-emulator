package test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/Bo887/chip8-emulator"
)

var cpu = chip8.CreateCpu()

func TestInitialMemory(t *testing.T) {
    assert.Equal(t, 4096, len(cpu.Memory))

    for i := range chip8.Fontset {
        assert.Equal(t, chip8.Fontset[i], cpu.Memory[i])
    }

    for i := 0x200; i < len(cpu.Memory); i++ {
        assert.Equal(t, uint8(0), cpu.Memory[i])
    }
}

func TestInitialPC(t *testing.T) {
    assert.Equal(t, uint16(0x200), cpu.PC)
}

func TestDefaultValues(t *testing.T) {
    assert.Equal(t, uint16(0), cpu.Opcode)

    assert.Equal(t, 16, len(cpu.Stack))
    for _, elem := range cpu.Stack {
        assert.Equal(t, uint16(0), elem)
    }

    assert.Equal(t, len(cpu.V), 16)
    for _, elem := range cpu.V {
        assert.Equal(t, uint8(0), elem)
    }

    assert.Equal(t, uint16(0), cpu.I)
    assert.Equal(t, uint16(0), cpu.SP)
    assert.Equal(t, uint8(0), cpu.DelayTimer)
    assert.Equal(t, uint8(0), cpu.SoundTimer)

    assert.Equal(t, 2048, len(cpu.Display))
    for _, elem := range cpu.Display {
        assert.Equal(t, uint8(0), elem)
    }

    assert.Equal(t, 16, len(cpu.Keypad))
    for _, elem := range cpu.Keypad{
        assert.Equal(t, uint8(0), elem)
    }
}
