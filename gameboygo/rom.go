package main

import (
	"io/ioutil"
	"fmt"
)

type rom_header struct{
	title 			string
	cart_type 		uint8
	rom_size 		uint8
	ram_size 		uint8
	header_checksum	uint8
	version			uint8
}

type Cart interface{
	init(data []uint8, filename string) bool
	romWrite(addr uint16, b uint8)
	romRead(addr uint16) uint8
	ramWrite(addr uint16, b uint8)
	ramRead(addr uint16) uint8
	saveCartRam()
}

const(
	ROM_ONLY 			 	uint8 = 0x00
	ROM_MBC1 			 	uint8 = 0x01
	ROM_MBC1_RAM 	 	 	uint8 = 0x02
	ROM_MBC1_RAM_BATT 	 	uint8 = 0x03
	ROM_MBC2 		 	 	uint8 = 0x05
	ROM_MBC2_BATT 		 	uint8 = 0x06
	ROM_RAM 			 	uint8 = 0x08
	ROM_RAM_BATT 		 	uint8 = 0x09
	ROM_MMM01 			 	uint8 = 0x0B
	ROM_MMM01_SRAM  	 	uint8 = 0x0C
	ROM_MMM01_SRAM_BATT  	uint8 = 0x0D
	ROM_MBC3_TIMER_BATT		uint8 = 0x0F
	ROM_MBC3_TIMER_RAM_BATT uint8 = 0x10
	ROM_MBC3 				uint8 = 0x11
	ROM_MBC3_RAM 		 	uint8 = 0x12
	ROM_MBC3_RAM_BATT    	uint8 = 0x13
	ROM_MBC5 			 	uint8 = 0x19
	ROM_MBC5_RAM         	uint8 = 0x1A
	ROM_MBC5_RAM_BATT		uint8 = 0x1B
	ROM_MBC5_RUMBLE			uint8 = 0x1C
	ROM_MBC5_RUMBLE_SRAM	uint8 = 0x1D
	ROM_MBC5_RUMBLE_SRAM_BATT	uint8 = 0x1E
	POCKETCAM				uint8 = 0x1F
	BANDAI_TAMA5			uint8 = 0xFD
	HUDSON_HUC3				uint8 = 0xFE
	HUDSON_HUC1				uint8 = 0xFF
)

var cart Cart

func (h *rom_header) init(data []byte){
	h.title = string(data[0x0134:0x143])
	h.cart_type = data[0x147]
	h.rom_size = data[0x148] //0x8000 << rom_size
	h.ram_size = data[0x149] //0->none, 1->2Kbytes, 2->8Kbytes, 3->32Kbytes
	h.version = data[0x14C]
	h.header_checksum = data[0x14D]
}

func Load_rom(filename string) bool{
	RomData, err := ioutil.ReadFile(filename)
	if err != nil || len(RomData) < 0x8000{
		return false
	}
	switch RomData[0x147]{
	case ROM_MBC1, ROM_MBC1_RAM, ROM_MBC1_RAM_BATT:
		cart = new(MBC1)
	case ROM_MBC2, ROM_MBC2_BATT:
		cart = new(MBC2)
	case ROM_MBC3_TIMER_BATT, ROM_MBC3_TIMER_RAM_BATT, ROM_MBC3, ROM_MBC3_RAM, ROM_MBC3_RAM_BATT:
		cart = new(MBC3)
	default:
		cart = new(genericCart)
	}
	if !cart.init(RomData, filename){
		fmt.Printf("ERROR creating Cart\n")
		return false
	}
	return true
}
func saveCartRam() {
	cart.saveCartRam()
}
