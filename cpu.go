package chip8emulator

import (
//    "fmt"
)

type Cpu struct {
    //opcode
    Opcode uint16

    //memory
    Memory [4096]uint8

    //stack
    Stack [16]uint16

    //registers (V0 -> VF)
    V [16]uint8

    //index register 
    I uint16

    //program counter
    PC uint16

    //stack pointer
    SP uint16

    //delay timer
    DelayTimer uint8

    //sound timer
    SoundTimer uint8

    //graphics
    Display [64 * 32]uint8

    //keypad
    Keypad [16]uint8
}
