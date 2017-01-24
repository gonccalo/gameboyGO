package main

import (
	"fmt"
	"gameboygo"
)
func main() {
    fmt.Println("wram: %d", len(gameboygo.Wram))
    fmt.Println("vram: %d", len(gameboygo.Vram))
    gameboygo.Load_rom("Tetris.gb")
    fmt.Printf("rom: %+v\n", gameboygo.Head)
}