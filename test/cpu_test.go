package test

import (
    "testing"
    "github.com/Bo887/chip8-emulator"
)

var cpu = chip8.CreateCpu()

func TestInitialMemory(t *testing.T) {
    AssertEqual(t, len(cpu.Memory), 4096)

    for i := range chip8.Fontset {
        AssertEqual(t, cpu.Memory[i], chip8.Fontset[i])
    }

    for i := 0x200; i < len(cpu.Memory); i++ {
        AssertEqual(t, cpu.Memory[i], uint8(0))
    }
}

func TestInitialPC(t *testing.T) {
    AssertEqual(t, cpu.PC, uint16(0x200))
}

func TestDefaultValues(t *testing.T) {
    AssertEqual(t, cpu.Opcode, uint16(0))

    AssertEqual(t, len(cpu.Stack), 16)
    for _, elem := range cpu.Stack {
        AssertEqual(t, elem, uint16(0))
    }

    AssertEqual(t, len(cpu.V), 16)
    for _, elem := range cpu.V {
        AssertEqual(t, elem, uint8(0))
    }

    AssertEqual(t, cpu.I, uint16(0))
    AssertEqual(t, cpu.SP, uint16(0))
    AssertEqual(t, cpu.DelayTimer, uint8(0))
    AssertEqual(t, cpu.SoundTimer, uint8(0))

    AssertEqual(t, len(cpu.Display), 2048)
    for _, elem := range cpu.Display {
        AssertEqual(t, elem, uint8(0))
    }

    AssertEqual(t, len(cpu.Keypad), 16)
    for _, elem := range cpu.Keypad{
        AssertEqual(t, elem, uint8(0))
    }
}
