package chip8

import (
//    "fmt"
    "os"
)

type Cpu struct {
    Opcode uint16
    Memory [kMemorySize]uint8
    Stack [kStackSize]uint16
    V [kNumRegisters]uint8 //main registers (V0 -> VF)
    I uint16 //index register
    PC uint16 //program counter
    SP uint16 //stack pointer
    DelayTimer uint8
    SoundTimer uint8
    Display [kDisplaySize]uint8
    Keypad [kKeypadSize]uint8
}

func CreateCpu() Cpu {
    rv := Cpu{}
    for i := range Fontset {
        rv.Memory[i] = Fontset[i]
    }
    //.text starts at 0x200
    rv.PC = 0x200
    return rv
}

func (cpu *Cpu) LoadProgram(path string) {
    file, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer file.Close()
}
