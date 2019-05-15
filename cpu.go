package chip8

import (
//    "fmt"
)

type Cpu struct {
    Opcode uint16
    Memory [4096]uint8
    Stack [16]uint16
    V [16]uint8 //main registers (V0 -> VF)
    I uint16 //index register
    PC uint16 //program counter
    SP uint16 //stack pointer
    DelayTimer uint8
    SoundTimer uint8
    Display [64 * 32]uint8
    Keypad [16]uint8
}

func CreateCpu() Cpu {
    rv := Cpu{}
    //.text starts at 0x200
    rv.PC = 0x200
    for i := range Fontset {
        rv.Memory[i] = Fontset[i]
    }
    return rv
}
