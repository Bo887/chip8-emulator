package main

import (
    "github.com/Bo887/chip8-emulator/chip8"
    //"github.com/gdamore/tcell"
)

func main() {
    cpu := chip8.CreateCpu()
    cpu.LoadProgram("roms/pong.ch8")

    for {
        err := cpu.Update()
        if err != nil {
            break
        }
    }

    /*
    screen, err := tcell.NewScreen()
    if err != nil {
        panic(err)
    }
    screen.Init()
    screen.Show()
    for{
        ev := screen.PollEvent()
        switch ev := ev.(type) {
        case *tcell.EventKey:
            switch ev.Key() {
            case tcell.KeyEscape, tcell.KeyEnter:
                    return
            case tcell.KeyRune:
                print(ev.Rune())
            }
        }
    }
    screen.Fini()
    */
}
