# gameboyGO [![Build Status](https://travis-ci.org/gonccalo/gameboyGO.svg?branch=master)](https://travis-ci.org/gonccalo/gameboyGO)
Gameboy emulator in go

## Current state:
* Interrupts
* Joypad
* Video
* Timers
* Support MBC1 & MBC2 roms
* Load saved games
* Cpu test roms passed
### Tested Games
* Tennis
* Tetris
* Dr. Mario
* Kirby's Dream Land
* Super Mario Land
* Pocket Monsters Ao
* The Legend of Zelda - Link's Awakening
* The Final Fantasy Legend
* Seiken Densetsu

![out3](https://cloud.githubusercontent.com/assets/5223817/25020097/09d047b4-2085-11e7-87af-c88fafc6a51a.gif)
![out2](https://cloud.githubusercontent.com/assets/5223817/24856667/7309c2f0-1ddd-11e7-82f3-ff527d51b0d8.gif)
![out](https://cloud.githubusercontent.com/assets/5223817/23906938/74379006-08c7-11e7-9f99-e7e6121e1a64.gif)

## TODO:
* STOP and HALT instructions
* Support MBC3 & MBC5
* Sound

## Compile and run

```bash
go get -v github.com/veandco/go-sdl2/sdl
go get -v github.com/gonccalo/gameboyGO/gameboygo
go build github.com/gonccalo/gameboyGO/gameboygo
export GODEBUG=cgocheck=0
./gameboygo -rom <file>
```

## Dependencies
[SDL2](https://github.com/veandco/go-sdl2)
