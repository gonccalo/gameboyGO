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
	rom_file		string
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

var Head rom_header
var RomData []uint8
var cartRam []uint8
var bankMode uint8
var romBank uint8
var ramBank uint8
var cartRamEnabled bool

func (h *rom_header) init(data []byte, f string){
	h.title = string(data[0x0134:0x143])
	h.cart_type = data[0x147]
	h.rom_size = data[0x148] //0x8000 << rom_size
	h.ram_size = data[0x149] //0->none, 1->2Kbytes, 2->8Kbytes, 3->32Kbytes
	h.version = data[0x14C]
	h.header_checksum = data[0x14D]
	h.rom_file = f
}
func Load_rom(filename string) bool{
	var err error
	RomData, err = ioutil.ReadFile(filename)
	if err != nil || len(RomData) < 0x8000{
		return false
	}
	Head.init(RomData, filename)
	switch Head.ram_size{
	case 0x00:
		if Head.cart_type == ROM_MBC2 || Head.cart_type == ROM_MBC2_BATT{
			//512*4bit
			cartRam = make([]uint8, 0x1FFF)
		} else{
			//no ram
		}
	case 0x01:
		//2KB
		cartRam = make([]uint8, 0x1FFF)
	case 0x02:
		//8KB  - 1 bank
		cartRam = make([]uint8, 0x1FFF)
	case 0x03:
		//32KB - 4 banks
		cartRam = make([]uint8, 0x1FFF*4)
	case 0x04:
		//128KB - 16 banks
		cartRam = make([]uint8, 0x1FFF*16)
	case 0x05:
		//64KB - 8 banks
		cartRam = make([]uint8, 0x1FFF*8)
	}
	Load_ram(filename + "_save")
	if num := copy(ram[:0x4000], RomData[:0x4000]); num != 0x4000 {
		//bank 0
		return false
	}
	if num := copy(ram[0x4000:0x8000], RomData[0x4000:0x8000]); num != 0x4000{
		//bank 1
		return false
	}
	return true
}
func handleCartRamWrites(addr uint16, b uint8) {
	// Cart ram
	if cartRamEnabled {
		if Head.cart_type == ROM_MBC2 || Head.cart_type == ROM_MBC2_BATT {
			if addr > 0xA1FF {
				return
			}
			ram[addr] = b & 0x0F
		} else{
			ram[addr] = b
		}
	}
}
func handleRomWrites(addr uint16, b uint8) {
	if Head.cart_type == ROM_MBC1 || Head.cart_type == ROM_MBC1_RAM_BATT || Head.cart_type == ROM_MBC1_RAM{
		if addr >= 0x2000 && addr <= 0x3FFF{
			changeLRomBank(b)
		} else if addr >= 0x4000 && addr <= 0x5FFF{
			changeHRomOrRamBank(b)
		} else if addr >= 0x6000 && addr <= 0x7FFF{
			changeRomRamMode(b)
		} else if addr >= 0x0000 && addr <= 0x1FFF{
			ramEnable(b)
		}
	} else if Head.cart_type == ROM_MBC2 || Head.cart_type == ROM_MBC2_BATT {
		if (addr >= 0x0000 && addr <= 0x1FFF) && (addr&0x0100) == 0x0000{
			ramEnable(b)
		} else if (addr >= 0x2000 && addr <= 0x3FFF) && (addr&0x0100) == 0x0100{
			changeRomBank(b & 0x0F)
		}
	}
}

func changeLRomBank(num uint8) {
	var newRomBank uint8
	newRomBank = romBank & 0xE0
	newRomBank = newRomBank | (num & 0x1F)
	switch newRomBank{
	case 0x20:
		newRomBank = 0x21
	case 0x40:
		newRomBank = 0x41
	case 0x60:
		newRomBank = 0x61
	}
	changeRomBank(newRomBank)
}

func changeHRomOrRamBank(num uint8) {
	if bankMode == 0{
		var newRomBank uint8
		newRomBank = romBank & 0x1F
		newRomBank = newRomBank | ((num&0x03)<<5)
		switch newRomBank{
		case 0x20:
			newRomBank = 0x21
		case 0x40:
			newRomBank = 0x41
		case 0x60:
			newRomBank = 0x61
		}
		changeRomBank(newRomBank)
	} else if bankMode == 1 {
		newRamBank := num & 0x03
		changeRamBank(newRamBank)
	}
}
func changeRomBank(newRomBank uint8) {
	if newRomBank == 0x00 {
		newRomBank = 0x01
	}
	if newRomBank == romBank {
		return
	}
	romBank = newRomBank
	if num := copy(ram[0x4000:0x8000], RomData[0x4000 * uint32(romBank):0x4000 * (uint32(romBank) + 1)]); num != 0x4000{
		fmt.Printf("ERROR in rom bank %d, copied %d\n",romBank, num)
		return
	}
}
func changeRamBank(nRamBank uint8) {
	if num := copy(cartRam[0x1FFF * uint32(ramBank):0x1FFF * (uint32(ramBank) + 1)], ram[0xA000:0xBFFF]); num != 0x1FFF{
		fmt.Printf("ERROR in ram bank %d, copied %d\n",ramBank, num)
		return
	}
	ramBank = nRamBank
	if num := copy(ram[0xA000:0xBFFF], cartRam[0x1FFF * uint32(ramBank):0x1FFF * (uint32(ramBank) + 1)]); num != 0x1FFF{
		fmt.Printf("ERROR in ram bank %d, copied %d\n",ramBank, num)
		return
	}
}
func changeRomRamMode(b uint8) {
	if (b & 0x01) == 0 {
		bankMode = 0
		changeRamBank(0)
	} else{
		bankMode = 1
		changeRomBank(romBank & 0x1F)
	}
}
func ramEnable(b uint8) {
	if b & 0x0F == 0x0A{
		//enable
		cartRamEnabled = true
	} else{
		//disable
		cartRamEnabled = false
	}
}
func saveCartRam() {
	if Head.cart_type == ROM_MBC1_RAM_BATT || Head.cart_type == ROM_MBC2_BATT{
		if num := copy(cartRam[0x1FFF * uint32(ramBank):0x1FFF * (uint32(ramBank) + 1)], ram[0xA000:0xBFFF]); num != 0x1FFF{
			fmt.Printf("ERROR in ram bank %d, copied %d\n", ramBank, num)
			return
		}
		if err := ioutil.WriteFile(Head.rom_file + "_save", cartRam, 0666); err != nil{
			panic(err)
		}
	}
}
func Load_ram(filename string) {
	if Head.cart_type == ROM_MBC1_RAM_BATT || Head.cart_type == ROM_MBC2_BATT{
		loadCartRam(filename)
	}
}
func loadCartRam(filename string) {
	newCartRam, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("ERROR reading cart Ram")
		return
	}
	if len(newCartRam) != len(cartRam) {
		fmt.Printf("ERROR reading cart Ram")
		return
	}
	cartRam = newCartRam
	if num := copy(ram[0xA000:0xBFFF], cartRam[0x1FFF * uint32(ramBank):0x1FFF * (uint32(ramBank) + 1)]); num != 0x1FFF{
		fmt.Printf("ERRO in ram bank %d, copiados %d\n", ramBank, num)
		return
	}
}