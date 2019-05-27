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

func (cpu *Cpu) UpdateKeys (screen tcell.Screen) bool {
    for i := range Keypad {
        cpu.Keypad[i] = 0
    }
    ev := screen.PollEvent()
    switch ev := ev.(type) {
    case *tcell.EventKey:
        switch ev.Key() {
        case tcell.KeyEscape, tcell.KeyEnter:
            return true
        case tcell.KeyRune:
            for i, elem := range Keypad {
                if uint8(ev.Rune()) == elem {
                    cpu.Keypad[i] = 1
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

func InitScreen(screen tcell.Screen) {
    screen.Init()
    screen.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))
    screen.Show()
}

func DisplayToWindow(x int, y int, width int, height int) (int, int) {
    rescaled_x := float32(x)*float32(width)/float32(NUM_COLS)
    rescaled_y := float32(y)*float32(height)/float32(NUM_ROWS)
    return int(rescaled_x), int(rescaled_y)
}

func (cpu *Cpu) DrawScreen(screen tcell.Screen) {
    width, height := screen.Size()

    const filled = '▄'
    const empty = ' '

    for x := 0; x < NUM_COLS; x++ {
        for y := 0; y < NUM_ROWS; y++ {
            new_x, new_y := DisplayToWindow(x, y, width, height)
            if cpu.Display[y * NUM_COLS + x] == uint8(1.0) {
                screen.SetCell(new_x, new_y, tcell.StyleDefault, filled)
            } else {
                screen.SetCell(new_x, new_y, tcell.StyleDefault, empty)
            }
        }
    }
    screen.Show()
}
