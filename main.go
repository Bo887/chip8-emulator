package main

import (
    "github.com/Bo887/chip8-emulator/chip8"
    //"github.com/gdamore/tcell"
)

func main() {
    cpu := chip8.CreateCpu()
    err := cpu.LoadProgram("roms/pong.ch8")
    if err != nil {
        println(err.Error())
    }

    /*
    screen, err := tcell.NewScreen()
    if err != nil {
        panic(err)
    }
    screen.Init()
    screen.Show()
    */

    for{
        /*
        ev := screen.PollEvent()
        done := cpu.UpdateKeys(ev)
        if done {
            break
        }
        */

        err := cpu.Update()
        if err != nil {
            print(err.Error())
            break
        }

        if cpu.ShouldDraw {
            cpu.PrintScreen()
            println()
        }
    }
    //screen.Fini()
}
