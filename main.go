package main

import (
    "github.com/Bo887/chip8-emulator/chip8"
    "github.com/gdamore/tcell"
)

func main() {
    cpu := chip8.CreateCpu()
    err := cpu.LoadProgram("roms/chip8-picture.ch8")
    if err != nil {
        println(err.Error())
    }

    screen, err := tcell.NewScreen()
    if err != nil {
        panic(err)
    }

    chip8.InitScreen(screen)

    go func () {
        for {
            cpu.UpdateKeys(screen)
        }
    }()

    for {
        err := cpu.Update()
        if err != nil {
            print(err.Error())
            break
        }

        if cpu.ShouldDraw {
            cpu.DrawScreen(screen)
        }
    }
    screen.Fini()
}
