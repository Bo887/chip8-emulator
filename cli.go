package main

import (
	"errors"
	"github.com/Bo887/chip8-emulator/chip8"
	"github.com/gdamore/tcell"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Chip8 Emulator"
	app.Description = "An emulator for Chip8 programs. Pass the path to the ROM as the first argument."
	app.Action = RunEmulator

	err := app.Run(os.Args)
	if err != nil {
		println(err.Error())
	}
}

func RunEmulator(c *cli.Context) error {
	cpu := chip8.CreateCpu()
	args := c.Args()

	if len(args) != 1 {
		return errors.New("Not enough arguments present! Please specify the filename of the ROM!")
	}

	err := cpu.LoadProgram(args.Get(0))
	if err != nil {
		return err
	}

	screen, err := tcell.NewScreen()
	if err != nil {
		return err
	}

	chip8.InitScreen(screen)

	status := make(chan bool)

	go func() {
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
			return err
		}

		if cpu.ShouldDraw {
			cpu.DrawScreen(screen)
		}
	}
	screen.Fini()
	return nil
}
