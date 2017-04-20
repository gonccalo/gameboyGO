package main

import (
	"io/ioutil"
	"fmt"
)

type MBC1 struct {
	rom_file  string
	head      rom_header
	RomData []uint8
	cartRam []uint8
	bankMode  uint8
	romBank   uint8
	ramBank   uint8
	cartRamEnabled bool
}
func (r *MBC1)init(data []uint8, filename string) bool{
	r.head.init(data)
	r.rom_file = filename
	r.romBank = 1
	r.cartRam = make([]uint8, 0x1FFF*4)
	r.loadRam()
	r.RomData = make([]uint8, len(data))
	if num := copy(r.RomData, data); num != len(data){
		return false
	}
	return true
}
func (r *MBC1)romRead(addr uint16) uint8 {
	if addr >= 0x4000 {
		return r.RomData[(0x4000 * uint32(r.romBank)) + (uint32(addr)-0x4000)]
	}
	return r.RomData[addr]
}
func (r *MBC1)romWrite(addr uint16, b uint8) {
	if addr >= 0x2000 && addr <= 0x3FFF{
		switch newRomBank := (r.romBank & 0xE0) | (b & 0x1F); newRomBank{
		case 0x00, 0x20, 0x40, 0x60:
			r.romBank = newRomBank + 1
		default:
			r.romBank = newRomBank
		}
	} else if addr >= 0x4000 && addr <= 0x5FFF{
		if r.bankMode == 0{
			switch newRomBank := (r.romBank & 0x1F) | ((b & 0x03) << 5); newRomBank{
			case 0x00, 0x20, 0x40, 0x60:
				r.romBank = newRomBank + 1
			default:
				r.romBank = newRomBank
			}	
		} else if r.bankMode == 1 {
			r.ramBank = (b & 0x03)
		}
	} else if addr >= 0x6000 && addr <= 0x7FFF{
		if (b & 0x01) == 0 {
			r.bankMode = 0
			r.ramBank = 0
		} else{
			r.bankMode = 1
			switch newRomBank := (r.romBank & 0x1F); newRomBank{
			case 0x00, 0x20, 0x40, 0x60:
				r.romBank = newRomBank + 1
			default:
				r.romBank = newRomBank
			}
		}
	} else if addr >= 0x0000 && addr <= 0x1FFF{
		//enable or disable ram
		if b & 0x0F == 0x0A{
			r.cartRamEnabled = true
		} else{
			r.cartRamEnabled = false
		}
	}
}
func (r *MBC1)ramRead(addr uint16) uint8 {
	if !r.cartRamEnabled {
		return 0xFF
	}
	return r.cartRam[(0x1FFF * uint32(r.ramBank)) + (uint32(addr)-0xA000)]
}
func (r *MBC1)ramWrite(addr uint16, b uint8) {
	if r.cartRamEnabled {
		r.cartRam[(0x1FFF * uint32(r.ramBank)) + (uint32(addr)-0xA000)] = b
	}
}

func (r *MBC1)loadRam() {
	newCartRam, err := ioutil.ReadFile(r.rom_file + "_save")
	if err != nil {
		fmt.Printf("ERROR reading cart Ram")
		return
	}
	if len(newCartRam) != len(r.cartRam) {
		fmt.Printf("ERROR reading cart Ram")
		return
	}
	r.cartRam = newCartRam
}

func (r *MBC1)saveCartRam() {
	if err := ioutil.WriteFile(r.rom_file + "_save", r.cartRam, 0666); err != nil{
		panic(err)
	}
}