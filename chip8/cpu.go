package chip8

import (
    "os"
    "bufio"
    "time"
    "errors"
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
    TimeController <-chan time.Time
}

func CreateCpu() Cpu {
    rv := Cpu{}
    for i := range Fontset {
        rv.Memory[i] = Fontset[i]
    }
    rv.PC = PC_START
    rv.TimeController = time.Tick(time.Second/time.Duration(600))
    return rv
}

func (cpu *Cpu) LoadProgram(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    stats, err := file.Stat()
    if err != nil {
        return err
    }

    size := stats.Size()
    if size > MEMORY_SIZE - PC_START {
        return errors.New("Program will not fit in memory!")
    }

    bytes := make([]byte, size)
    reader := bufio.NewReader(file)
    _, err = reader.Read(bytes)
    if err != nil {
        return err
    }

    copy(cpu.Memory[PC_START:], bytes)
    return nil
}

func (cpu *Cpu) Update() error {
    cpu.Opcode = (uint16(cpu.Memory[cpu.PC]) << 8) | uint16(cpu.Memory[cpu.PC+1])
    cpu.ShouldDraw = false
    select {
    case <-cpu.TimeController:
        err := cpu.HandleOpcode()
        if cpu.DelayTimer > 0 {
            cpu.DelayTimer--
        }
        if cpu.SoundTimer > 0 {
            cpu.SoundTimer--
        }
        return err
    }
    return nil
}
