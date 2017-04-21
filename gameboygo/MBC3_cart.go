package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

type MBC3 struct {
	rom_file  string
	head      rom_header
	RomData []uint8
	cartRam []uint8
	RTCregs []uint8
	romBank   uint8
	ramBank   uint8
	mapRTC	  bool
	ramRTCEnabled bool
}
func (r *MBC3)init(data []uint8, filename string) bool{
	r.head.init(data)
	r.rom_file = filename
	r.romBank = 1
	r.RTCregs = make([]uint8, 5)
	r.cartRam = make([]uint8, 0x1FFF*4)
	r.loadRam()
	r.RomData = make([]uint8, len(data))
	if num := copy(r.RomData, data); num != len(data){
		return false
	}
	return true
}
func (r *MBC3)romRead(addr uint16) uint8 {
	if addr >= 0x4000 {
		return r.RomData[(0x4000 * uint32(r.romBank)) + (uint32(addr)-0x4000)]
	}
	return r.RomData[addr]
}
func (r *MBC3)romWrite(addr uint16, b uint8) {
	if addr >= 0x2000 && addr <= 0x3FFF{
		switch b & 0x7F{
		case 0x00:
			r.romBank = 0x01
		default:
			r.romBank = (b & 0x7F)
		}
	} else if addr >= 0x4000 && addr <= 0x5FFF{
		if b >= 0x08 && b <= 0x0C{
			r.mapRTC = true
			r.ramBank = b
			return
		}
		if b <= 0x03 {
			r.mapRTC = false
			r.ramBank = b
		}
	} else if addr >= 0x0000 && addr <= 0x1FFF{
		//enable or disable ram
		if b & 0x0F == 0x0A{
			r.ramRTCEnabled = true
		} else{
			r.ramRTCEnabled = false
		}
	} else if addr >= 0x6000 && addr <= 0x7FFF{
		if b == 0x01 {
			//latch rtc
		}
	}
}
func (r *MBC3)ramRead(addr uint16) uint8 {
	if !r.ramRTCEnabled {
		return 0xFF
	}
	if r.mapRTC {
		return r.RTCregs[r.ramBank - 0x08]
	}
	return r.cartRam[(0x1FFF * uint32(r.ramBank)) + (uint32(addr)-0xA000)]
}
func (r *MBC3)ramWrite(addr uint16, b uint8) {
	if !r.ramRTCEnabled {
		return
	}
	if r.mapRTC {
		r.RTCregs[r.ramBank - 0x08] = b
		return 
	}
	r.cartRam[(0x1FFF * uint32(r.ramBank)) + (uint32(addr)-0xA000)] = b
}

func (r *MBC3)loadRam() {
	newCartRam, err := ioutil.ReadFile(r.rom_file + "_save")
	if err != nil {
		if os.IsNotExist(err){
			return	
		}
		fmt.Printf("ERROR reading cart Ram")
		return
	}
	if len(newCartRam) != len(r.cartRam) {
		fmt.Printf("ERROR reading cart Ram")
		return
	}
	r.cartRam = newCartRam
}

func (r *MBC3)saveCartRam() {
	if err := ioutil.WriteFile(r.rom_file + "_save", r.cartRam, 0666); err != nil{
		panic(err)
	}
}