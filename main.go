package main

import (
	"fmt"
	"gameboygo"
	"time"
	"os"
	"github.com/veandco/go-sdl2/sdl"
)
func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
        800, 600, sdl.WINDOW_SHOWN)
    if err != nil {
        panic(err)
    }
    defer window.Destroy()
    renderer, err := sdl.CreateRenderer(window, -1, 0)
    if err != nil {
    	panic(err)
    }
    defer renderer.Destroy()
    gameboygo.Load_rom("Tetris.gb")
    fmt.Printf("rom: %+v\n", gameboygo.Head)
    gameboygo.Reset()
	for{
		gameboygo.LastTimer = 0
		gameboygo.LastScanLine = 0
		t := time.Now()
		for gameboygo.CicleCounter = 0; gameboygo.CicleCounter < gameboygo.CPU_FREQ; {
			handleInput()
			gameboygo.Execute()
			gameboygo.UpdateGPU(renderer)
		}
		renderer.Present() 
		fmt.Println(time.Second - time.Since(t))
		time.Sleep(time.Second - time.Since(t))
	}
}

func handleInput() {
	var event sdl.Event = sdl.PollEvent()
	switch eventType := event.(type) {
	case *sdl.KeyDownEvent:
		switch eventType.Keysym.Sym{
		case sdl.K_RIGHT :	//right
			gameboygo.KeyPressed(gameboygo.KEY_RIGHT)
		case sdl.K_LEFT:	//left
			gameboygo.KeyPressed(gameboygo.KEY_LEFT)
		case sdl.K_DOWN:	//down
			gameboygo.KeyPressed(gameboygo.KEY_DOWN)
		case sdl.K_UP:		//ip
			gameboygo.KeyPressed(gameboygo.KEY_UP)
		case sdl.K_z:		//a
			gameboygo.KeyPressed(gameboygo.KEY_A)
		case sdl.K_x:		//b
			gameboygo.KeyPressed(gameboygo.KEY_B)
		case sdl.K_q:		//start
			gameboygo.KeyPressed(gameboygo.KEY_START)
		case sdl.K_w:		//select
			gameboygo.KeyPressed(gameboygo.KEY_SELECT)
		}
	case *sdl.KeyUpEvent:
		switch eventType.Keysym.Sym{
		case sdl.K_ESCAPE:
			fmt.Printf("ESC\n\n\n\n\n")
			sdl.Quit()
			gameboygo.PrintStats()
			os.Exit(0)
		case sdl.K_RIGHT :	//right
			gameboygo.KeyReleased(gameboygo.KEY_RIGHT)
		case sdl.K_LEFT:	//left
			gameboygo.KeyReleased(gameboygo.KEY_LEFT)
		case sdl.K_DOWN:	//down
			gameboygo.KeyReleased(gameboygo.KEY_DOWN)
		case sdl.K_UP:		//ip
			gameboygo.KeyReleased(gameboygo.KEY_UP)
		case sdl.K_z:		//a
			gameboygo.KeyReleased(gameboygo.KEY_A)
		case sdl.K_x:		//b
			gameboygo.KeyReleased(gameboygo.KEY_B)
		case sdl.K_q:		//start
			gameboygo.KeyReleased(gameboygo.KEY_START)
		case sdl.K_w:		//select
			gameboygo.KeyReleased(gameboygo.KEY_SELECT)
		}
	}
}