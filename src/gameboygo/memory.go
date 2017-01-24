package gameboygo

var ram [0x10000]uint8
var Rom1 		= ram[0x0000:0x4000]
var RomN 		= ram[0x4000:0x8000]
var Vram 		= ram[0x8000:0xA000]
var CartRam 	= ram[0xA000:0xC000]
var Wram 		= ram[0xC000:0xE000]
var EchoRam 	= ram[0xC000:0xDE00]
var Oram 		= ram[0xFE00:0xFEA0]
var Io 			= ram[0xFF00:0xFF80]
var Hram 		= ram[0xFF80:0xFFFF]
var InterruptReg = &ram[0xFFFF]





/*
func (r *mem) read(addr uint16) uint8{

}
func (r * mem)write(addr uint16, word uint8) {
	
}
*/