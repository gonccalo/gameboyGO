package main
//import "os"
//import "fmt"
//import "io/ioutil"
var ram [0x10000]uint8

func writeByte(addr uint16, b uint8) bool{
	if addr < 0x8000{
		//handleRomWrites(addr, b)
		cart.romWrite(addr, b)
		return true
	}
	if addr >= 0xE000 && addr < 0xFE00 {	//echo zone
		ram[addr] = b
		ram[0xC000+(addr-0xE000)] = b
		return true
	} else if addr == 0xFF07{
		ram[addr] = b
		setupTimers()
		return true
	} else if (addr == 0xFF04) || (addr == 0xFF44){				//Divider register || LCDC y coordinate
		ram[addr] = 0x00
		return true
	} else if addr == 0xFF46 {
		/*
		Writing to this register launches a DMA transfer from ROM or RAM to OAM memory (sprite attribute table). The written value specifies the transfer source address divided by 100h, ie. source & destination are:
  		Source:      XX00-XX9F   ;XX in range from 00-F1h
  		Destination: FE00-FE9F
		*/
		var src uint16 = uint16(b) << 8
		var i uint16
		for i = 0; i < 160; i++ {
			writeByte(0xFe00+i, readByte(src+i))
		}
		return true
	} else if addr == 0xFF01{
		//serial?
	} else if addr >= 0xA000 && addr <= 0xBFFF{
		// Cart ram
		//handleCartRamWrites(addr, b)
		cart.ramWrite(addr, b)
		return true
	}
	ram[addr] = b
	return true
}
func readByte(addr uint16) uint8{
	if addr < 0x8000{
		return cart.romRead(addr)
	}
	if addr == 0xFF00{
		return getKeys(ram[addr])
	}
	if addr >= 0xA000 && addr <= 0xBFFF{
		return cart.ramRead(addr)
	}
	return ram[addr]
}
func read16bits(addr uint16) uint16{
	return uint16(uint16(readByte(addr)) | (uint16(readByte(addr+1)) << 8))
}
func write16bits(addr uint16, data uint16) {
	writeByte(addr, uint8(data & 0x00FF))
	writeByte(addr+1,uint8((data & 0xFF00) >> 8))
}