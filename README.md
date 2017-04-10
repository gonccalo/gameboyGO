[![Build Status](https://travis-ci.org/gonccalo/gameboyGO.svg?branch=master)](https://travis-ci.org/gonccalo/gameboyGO)

# gameboyGO
Gameboy emulator in go

## Current state:
* Interrupts
* Joypad
* Video
* Timers
* Cpu test roms passed
* Tennis and Tetris working

![out2](https://cloud.githubusercontent.com/assets/5223817/24856667/7309c2f0-1ddd-11e7-82f3-ff527d51b0d8.gif)
![out](https://cloud.githubusercontent.com/assets/5223817/23906938/74379006-08c7-11e7-9f99-e7e6121e1a64.gif)

## TODO:
* Fix framerate
* STOP and HALT instructions
* Support roms bigger than 32K
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
