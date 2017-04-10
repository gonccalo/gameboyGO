package main

import (
	"fmt"
	"time"
	"os"
	go_flag "flag"
	//"runtime/pprof"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	/*
	if *profiling {
    	var cpuprofile string = "prof"
		f, err := os.Create(cpuprofile)
    	if err != nil {
        	panic(err)
    	}
    	pprof.StartCPUProfile(f)
    	defer pprof.StopCPUProfile()
    }
    */
    //flags
	romFile := go_flag.String("rom", "Bc.gb", "Path to rom file")
	fdebug  := go_flag.Bool("debug", false, "Set to true to enable debugging")
	go_flag.Parse()

	//init sdl stuff
	sdl.Init(sdl.INIT_EVERYTHING)
	window, err := sdl.CreateWindow("gameboyGO", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
        256, 256, sdl.WINDOW_SHOWN)
    if err != nil {
        panic(err)
    }
    defer window.Destroy()
    renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
    if err != nil {
    	panic(err)
    }
    defer renderer.Destroy()
    texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_STREAMING, WIDTH, HEIGHT)
    if err != nil {
    	panic(err)
    }
    
    Load_rom(*romFile)
    fmt.Printf("rom: %+v\n", Head)
    Reset()
    renderer.SetDrawColor(0,0,0,255)
	renderer.Clear()
	if *fdebug {
		for {
			t := time.Now()
			for CicleCounter = 0; (CicleCounter < CPU_FREQ); {
				handleInput()
				ExecuteDebug()
				UpdateGPU(renderer, texture, cicles_op)
			} 
			fmt.Println(time.Second - time.Since(t))
			time.Sleep(time.Second - time.Since(t))
		}
	} else{
		for {
			CicleCounter = 0
			t := time.Now()
			for ;(CicleCounter <= CPU_FREQ); {
				handleInput()
				Execute()
				UpdateGPU(renderer, texture, cicles_op)
			} 
			fmt.Println(time.Second - time.Since(t))
			time.Sleep(time.Second - time.Since(t))
		}
	}
}

func handleInput(){
	var event sdl.Event = sdl.PollEvent()
	switch eventType := event.(type) {
	case *sdl.KeyDownEvent:
		switch eventType.Keysym.Sym{
		case sdl.K_RIGHT :	//right
			KeyPressed(KEY_RIGHT)
		case sdl.K_LEFT:	//left
			KeyPressed(KEY_LEFT)
		case sdl.K_DOWN:	//down
			KeyPressed(KEY_DOWN)
		case sdl.K_UP:		//ip
			KeyPressed(KEY_UP)
		case sdl.K_z:		//a
			KeyPressed(KEY_A)
		case sdl.K_x:		//b
			KeyPressed(KEY_B)
		case sdl.K_q:		//start
			KeyPressed(KEY_START)
		case sdl.K_w:		//select
			KeyPressed(KEY_SELECT)
		}
	case *sdl.KeyUpEvent:
		switch eventType.Keysym.Sym{
		case sdl.K_ESCAPE:
			//printStats()
			//fmt.Printf("%v", ram)
			sdl.Quit()
			os.Exit(0)
			/*
			if *profiling {
				pprof.StopCPUProfile()	
			}
			*/
		case sdl.K_RIGHT :	//right
			KeyReleased(KEY_RIGHT)
		case sdl.K_LEFT:	//left
			KeyReleased(KEY_LEFT)
		case sdl.K_DOWN:	//down
			KeyReleased(KEY_DOWN)
		case sdl.K_UP:		//ip
			KeyReleased(KEY_UP)
		case sdl.K_z:		//a
			KeyReleased(KEY_A)
		case sdl.K_x:		//b
			KeyReleased(KEY_B)
		case sdl.K_q:		//start
			KeyReleased(KEY_START)
		case sdl.K_w:		//select
			KeyReleased(KEY_SELECT)
		}
	}
}