package main

import (
    "github.com/Bo887/chip8-emulator/chip8"
    "github.com/gdamore/tcell"
)

func main() {
    cpu := chip8.CreateCpu()
    err := cpu.LoadProgram("roms/pong.ch8")
    if err != nil {
        println(err.Error())
        return
    }

    screen, err := tcell.NewScreen()
    if err != nil {
        println(err.Error())
        return
    }

    chip8.InitScreen(screen)

    status := make(chan bool)

    go func () {
        for {
            status <- cpu.UpdateKeys(screen)
        }
    }()

    for {
        var done bool
        select {
        case done = <-status:
        default:
            done = false
        }

        if done {
            break
        }

        err := cpu.Update()
        if err != nil {
            println(err.Error())
            break
        }

        if cpu.ShouldDraw {
            cpu.DrawScreen(screen)
        }
    }
    screen.Fini()
}
