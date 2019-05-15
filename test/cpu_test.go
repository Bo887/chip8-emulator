package test

import (
    "testing"
    "github.com/Bo887/chip8-emulator"
)

func TestDefaultValues(t *testing.T){
    cpu := chip8.CreateCpu()
    AssertEqual(t, cpu.Opcode, uint16(0))

    AssertEqual(t, len(cpu.Memory), 4096)
    for _, elem := range cpu.Memory {
        AssertEqual(t, elem, uint8(0))
    }

    AssertEqual(t, len(cpu.Stack), 16)
    for _, elem := range cpu.Stack {
        AssertEqual(t, elem, uint16(0))
    }

    AssertEqual(t, len(cpu.V), 16)
    for _, elem := range cpu.V {
        AssertEqual(t, elem, uint8(0))
    }

    AssertEqual(t, cpu.I, uint16(0))
    AssertEqual(t, cpu.PC, uint16(0x200))
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
