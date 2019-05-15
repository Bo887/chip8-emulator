package chip8

import (
//    "fmt"
)

func (cpu *Cpu) Handle00Opcodes() {
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
        panic("Unknown opcode in 00--!")
    }
    cpu.PC += 2
}

//jump to address NNN
func (cpu *Cpu) Handle1NNNOpcode() {
    target := cpu.Opcode & 0x0FFF
    cpu.PC = target
}

//call function (jump-and-link) at address NNN
func (cpu *Cpu) Handle2NNNOpcode() {
    cpu.Stack[cpu.SP] = cpu.PC      //store current address in stack
    cpu.SP++
    target := cpu.Opcode & 0x0FFF
    cpu.PC = target
}

//Skip next instruction if V[X] == NN
func (cpu *Cpu) Handle3XNNOpcode() {
    nn := uint8(cpu.Opcode & 0x00FF)
    x := cpu.Opcode & 0x0F00
    if cpu.V[x] == nn {
        cpu.PC += 4
    } else {
        cpu.PC += 2
    }
}

//Skip next instruction if V[X] != NN
func (cpu *Cpu) Handle4XNNOpcode() {
    nn := uint8(cpu.Opcode & 0x00FF)
    x := cpu.Opcode & 0x0F00
    if cpu.V[x] != nn {
        cpu.PC += 4
    } else {
        cpu.PC += 2
    }
}

//Skip next instruction if V[X] != V[Y]
func (cpu *Cpu) Handle5XY0Opcode() {
    if cpu.Opcode & 0x000F != 0 {
        panic("Last word in 5XY0 instructions should be 0!")
    }
    x := cpu.Opcode & 0x0F00
    y := cpu.Opcode & 0x00F0
    if cpu.V[x] == cpu.V[y] {
        cpu.PC += 4
    } else {
        cpu.PC += 2
    }
}

//V[X] = NN
func (cpu *Cpu) Handle6XNNOpcode() {
    nn := uint8(cpu.Opcode & 0x00FF)
    x := uint8(cpu.Opcode & 0x0F00)
    cpu.V[nn] = x
    cpu.PC += 2
}

//V[X] += NN
func (cpu *Cpu) Handle7xNNOpcode() {
    nn := uint8(cpu.Opcode & 0x00FF)
    x := uint8(cpu.Opcode & 0x0F00)
    cpu.V[nn] += x
    cpu.PC += 2
}

func (cpu *Cpu) Handle8XYOpcodes() {
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
        panic("Unknown opcode in 8XY-!")
    }
    cpu.PC += 2
}
