package chip8

import (
    "os"
)

type Cpu struct {
    Opcode uint16
    Memory [MEMORY_SIZE]uint8
    Stack [STACK_SIZE]uint16
    V [NUM_REGISTERS]uint8 //main registers (V0 -> VF)
    I uint16 //index register
    PC uint16 //program counter
    SP uint16 //stack pointer
    DelayTimer uint8
    SoundTimer uint8
    Display [NUM_COLS * NUM_ROWS]uint8
    Keypad [KEYPAD_SIZE]uint8
    ShouldDraw bool
}

func CreateCpu() Cpu {
    rv := Cpu{}
    for i := range Fontset {
        rv.Memory[i] = Fontset[i]
    }
    rv.PC = PC_START
    return rv
}

func (cpu *Cpu) LoadProgram(path string) {
    file, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    //TODO: actually load program
}
