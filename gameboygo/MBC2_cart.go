package main

import (
	"io/ioutil"
	"fmt"
)

type MBC2 struct {
	rom_file  string
	head      rom_header
	RomData []uint8
	cartRam []uint8
	romBank   uint8
	cartRamEnabled bool
}

func (r *MBC2)init(data []uint8, filename string) bool{
	r.head.init(data)
	r.rom_file = filename
	r.romBank = 1
	r.cartRam = make([]uint8, 0x0200) // 512 * 4 bits
	r.loadRam()
	r.RomData = make([]uint8, len(data))
	if num := copy(r.RomData, data); num != len(data){
		return false
	}
	return true
}

func (r *MBC2)romRead(addr uint16) uint8 {
	if addr >= 0x4000 {
		return r.RomData[(0x4000 * uint32(r.romBank)) + (uint32(addr)-0x4000)]
	}
	return r.RomData[addr]
}

func (r *MBC2)romWrite(addr uint16, b uint8) {
	if (addr >= 0x0000 && addr <= 0x1FFF) && (addr&0x0100) == 0x0000{
		if b & 0x0F == 0x0A{
			r.cartRamEnabled = true
		} else{
			r.cartRamEnabled = false
		}
	} else if (addr >= 0x2000 && addr <= 0x3FFF) && (addr&0x0100) == 0x0100{
		if (b & 0x0F) == 0x00 {
			r.romBank = 0x01
		} else{
			r.romBank = (b & 0x0F)
		}
	}
}

func (r *MBC2)ramRead(addr uint16) uint8 {
	if !r.cartRamEnabled {
		return 0xFF
	}
	return r.cartRam[addr-0xA000]
}

func (r *MBC2)ramWrite(addr uint16, b uint8) {
	if addr > 0xA1FF {
		return
	}
	if r.cartRamEnabled {
		r.cartRam[addr-0xA000] = b & 0x0F
	}
}
func (r *MBC2)loadRam() {
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
func (r *MBC2)saveCartRam() {
	if err := ioutil.WriteFile(r.rom_file + "_save", r.cartRam, 0666); err != nil{
		panic(err)
	}
}