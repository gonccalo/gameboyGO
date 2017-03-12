# gameboyGO
Gameboy emulator in go
##Current state:
* Almost all instructions implemented
* Interrupts
* Joypad
* Background almost complete

![deepinscreenshot20170311174946](https://cloud.githubusercontent.com/assets/5223817/23828040/fdea1502-06bb-11e7-878d-41d2599a5f08.png)
![deepinscreenshot20170311174919](https://cloud.githubusercontent.com/assets/5223817/23828042/11940824-06bc-11e7-8d8e-8faa3f5fa198.png)


##TODO:
* Video (in progress)
* Missing instructions
* Support roms bigger than 32K
* Sound

##Compile and run

```bash
export GOPATH=/full/path/to/gameboyGO/
export GODEBUG=cgocheck=0
go build
./gameboyGO
```

##Dependencies
[SDL2](https://github.com/veandco/go-sdl2)
