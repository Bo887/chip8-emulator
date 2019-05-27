package chip8

import (
    "github.com/gdamore/tcell"
)

var Keypad = [...]uint8 {
    '1', '2', '3', '4',
    'q', 'w', 'e', 'r',
    'a', 's', 'd', 'f',
    'z', 'x', 'c', 'v',
}

func (cpu *Cpu) UpdateKeys (ev tcell.Event) bool {
    switch ev := ev.(type) {
    case *tcell.EventKey:
        switch ev.Key() {
        case tcell.KeyEscape, tcell.KeyEnter:
            return true
        case tcell.KeyRune:
            for i, elem := range Keypad {
                if uint8(ev.Rune()) == elem {
                    println(elem)
                    cpu.Keypad[i] = 1
                } else {
                    cpu.Keypad[i] = 0
                }
            }
        }
    }
    return false
}

func (cpu *Cpu) PrintScreen() {
    for i := range cpu.Display {
        if i % (NUM_COLS) == 0 {
            println()
        }
        if cpu.Display[i] == 1 {
            print("*")
        } else {
            print("0")
        }
    }
}
