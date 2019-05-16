package chip8

import (
//    "fmt"
    "errors"
)

func (cpu *Cpu) HandleOpcode() error {
    switch (cpu.Opcode & 0xF000) >> 12 {
    case 0x0:
        return cpu.Handle00Opcodes()
    case 0x1:
        return cpu.Handle1NNNOpcode()
    case 0x2:
        return cpu.Handle2NNNOpcode()
    case 0x3:
        return cpu.Handle3XNNOpcode()
    case 0x4:
        return cpu.Handle4XNNOpcode()
    case 0x5:
        return cpu.Handle5XY0Opcode()
    case 0x6:
        return cpu.Handle6XNNOpcode()
    case 0x7:
        return cpu.Handle7xNNOpcode()
    default:
        return errors.New("Unknown opcode in HandleOpcode()!")
    }
}

func (cpu *Cpu) Handle00Opcodes() error {
    switch cpu.Opcode {
    //clear screen
    case 0x00E0:
        for i := range cpu.Display {
            cpu.Display[i] = 0
        }
    //return from function
    case 0x00EE:
        cpu.SP--
        cpu.PC = cpu.Stack[cpu.SP]
    default:
        return errors.New("Invalid 00-- Opcode in Handle00Opcodes()!")
    }
    cpu.PC += 2
    return nil
}

//jump to address NNN
func (cpu *Cpu) Handle1NNNOpcode() error {
    target := cpu.Opcode & 0x0FFF
    cpu.PC = target
    return nil
}

//call function (jump-and-link) at address NNN
func (cpu *Cpu) Handle2NNNOpcode() error {
    if (cpu.SP >= kStackSize) {
        return errors.New("Cannot call function! Stack is full!")
    }
    cpu.Stack[cpu.SP] = cpu.PC      //store current address in stack
    cpu.SP++
    target := cpu.Opcode & 0x0FFF
    cpu.PC = target
    return nil
}

//Skip next instruction if V[X] == NN
func (cpu *Cpu) Handle3XNNOpcode() error {
    nn := uint8(cpu.Opcode & 0x00FF)
    x := (cpu.Opcode & 0x0F00) >> 8
    if cpu.V[x] == nn {
        cpu.PC += 4
    } else {
        cpu.PC += 2
    }
    return nil
}

//Skip next instruction if V[X] != NN
func (cpu *Cpu) Handle4XNNOpcode() error {
    nn := uint8(cpu.Opcode & 0x00FF)
    x := (cpu.Opcode & 0x0F00) >> 8
    if cpu.V[x] != nn {
        cpu.PC += 4
    } else {
        cpu.PC += 2
    }
    return nil
}

//Skip next instruction if V[X] == V[Y]
func (cpu *Cpu) Handle5XY0Opcode() error {
    if cpu.Opcode & 0x000F != 0 {
        return errors.New("Last word in 5XY0 instructions should be 0!")
    }
    x := (cpu.Opcode & 0x0F00) >> 8
    y := (cpu.Opcode & 0x00F0) >> 4
    if cpu.V[x] == cpu.V[y] {
        cpu.PC += 4
    } else {
        cpu.PC += 2
    }
    return nil
}

//V[X] = NN
func (cpu *Cpu) Handle6XNNOpcode() error {
    nn := uint8(cpu.Opcode & 0x00FF)
    x := uint8((cpu.Opcode & 0x0F00) >> 8)
    cpu.V[x] = nn
    cpu.PC += 2
    return nil
}

//V[X] += NN
func (cpu *Cpu) Handle7xNNOpcode() error {
    nn := uint8(cpu.Opcode & 0x00FF)
    x := uint8((cpu.Opcode & 0x0F00) >> 8)
    cpu.V[x] += nn
    cpu.PC += 2
    return nil
}

func (cpu *Cpu) Handle8XYOpcodes() error {
    switch cpu.Opcode & 0x000F {
    case 0x0:
    case 0x1:
    case 0x2:
    case 0x3:
    case 0x4:
    case 0x5:
    case 0x6:
    case 0x7:
    case 0xE:
    default:
        return errors.New("Unknown opcode in 8XY-!")
    }
    cpu.PC += 2
    return nil
}
