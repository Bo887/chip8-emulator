package chip8

import (
//    "fmt"
    "errors"
    "time"
    "math/rand"
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
    case 0x8:
        return cpu.Handle8XYOpcodes()
    case 0x9:
        return cpu.Handle9XY0Opcode()
    case 0xA:
        return cpu.HandleANNNOpcode()
    case 0xB:
        return cpu.HandleBNNNOpcode()
    case 0xC:
        return cpu.HandleCXNNOpcode(time.Now().UnixNano())
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
    x := (cpu.Opcode & 0x0F00) >> 8
    cpu.V[x] = nn
    cpu.PC += 2
    return nil
}

//V[X] += NN
func (cpu *Cpu) Handle7xNNOpcode() error {
    nn := uint8(cpu.Opcode & 0x00FF)
    x := (cpu.Opcode & 0x0F00) >> 8
    cpu.V[x] += nn
    cpu.PC += 2
    return nil
}

func (cpu *Cpu) Handle8XYOpcodes() error {
    x := (cpu.Opcode & 0x0F00) >> 8
    y := (cpu.Opcode & 0x00F0) >> 4
    switch cpu.Opcode & 0x000F {
    //V[x] = V[y]
    case 0x0:
        cpu.V[x] = cpu.V[y]
    //V[x] = V[x] | V[y]
    case 0x1:
        cpu.V[x] = cpu.V[x] | cpu.V[y]
    //V[x] = V[x] & V[y]
    case 0x2:
        cpu.V[x] = cpu.V[x] & cpu.V[y]
    //V[x] = V[x] ^ V[y]
    case 0x3:
        cpu.V[x] = cpu.V[x] ^ cpu.V[y]
    //V[x] += V[y]
    case 0x4:
        res := uint16(cpu.V[x]) + uint16(cpu.V[y])
        var carry uint8 = 0
        if res > 0xFF {
            carry = 1
        }
        cpu.V[x] = uint8(res)
        cpu.V[0xF] = carry
    //V[x] -= V[y]
    case 0x5:
        res := int16(cpu.V[x]) - int16(cpu.V[y])
        var not_borrow uint8 = 1
        if res < 0x00 {
            not_borrow = 0
        }
        cpu.V[x] = uint8(res)
        cpu.V[0xF] = not_borrow
    //V[x] >>= 1
    case 0x6:
        lsb := cpu.V[x] & 0x01
        cpu.V[0xF] = lsb
        cpu.V[x] >>= 1
    //V[x] = V[y] - V[x]
    case 0x7:
        res := int16(cpu.V[y]) - int16(cpu.V[x])
        var not_borrow uint8 = 1
        if res < 0x00 {
            not_borrow = 0
        }
        cpu.V[x] = uint8(res)
        cpu.V[0xF] = not_borrow
    //V[x] <<= 1
    case 0xE:
        lsb := cpu.V[x] & 0x01
        cpu.V[0xF] = lsb
        cpu.V[x] <<= 1
    default:
        return errors.New("Unknown opcode in 8XY-!")
    }
    cpu.PC += 2
    return nil
}

//Skip next instruction if V[X] != V[Y]
func (cpu *Cpu) Handle9XY0Opcode() error {
    if cpu.Opcode & 0x000F != 0 {
        return errors.New("Last word in 5XY0 instructions should be 0!")
    }
    x := (cpu.Opcode & 0x0F00) >> 8
    y := (cpu.Opcode & 0x00F0) >> 4
    if cpu.V[x] != cpu.V[y] {
        cpu.PC += 4
    } else {
        cpu.PC += 2
    }
    return nil
}

//I = NNN
func (cpu *Cpu) HandleANNNOpcode() error {
    nnn := cpu.Opcode & 0x0FFF
    cpu.I = nnn
    cpu.PC += 2
    return nil
}

//PC = V[0] + NNN
func (cpu *Cpu) HandleBNNNOpcode() error {
    nnn := cpu.Opcode & 0x0FFF
    cpu.PC = uint16(cpu.V[0]) + nnn
    return nil
}

//V[x] = rand() & NN
func (cpu *Cpu) HandleCXNNOpcode(seed int64) error {
    x := (cpu.Opcode & 0x0F00) >> 8
    nn := uint8(cpu.Opcode & 0x00FF)

    source := rand.NewSource(seed)
    rand_generator := rand.New(source)
    rand_num := uint8(rand_generator.Intn(256))
    cpu.V[x] = rand_num & nn
    return nil
}
