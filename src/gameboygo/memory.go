package gameboygo
//import "fmt"
var ram [0x10000]uint8
var Rom1 		= ram[0x0000:0x4000]
var RomN 		= ram[0x4000:0x8000]

func writeByte(addr uint16, b uint8) bool{
	if addr >= 0xE000 && addr < 0xFE00 {	//echo zone
		ram[addr] = b
		ram[0xC000+(addr-0xE000)] = b
		return true
	} else if (addr == 0xFF04) || (addr == 0xFF44){				//Divider register || LCDC y coordinate
		ram[addr] = 0x00
		return true
	} else if addr == 0xFF46 {				// OAM DMA
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
	}
	ram[addr] = b
	return true
}
func readByte(addr uint16) uint8{
	if addr == 0xFF00{
		return getKeys(ram[addr])
	}
	return ram[addr]
}
func read16bits(addr uint16) uint16{
	return uint16(uint16(ram[addr]) | (uint16(ram[addr+1]) << 8))
}
func write16bits(addr uint16, data uint16) {
	writeByte(addr, uint8(data & 0x00FF))
	writeByte(addr+1,uint8((data & 0xFF00) >> 8))
}
func getTimerFreq() int{ //0 if clock is stoped or frequency if started
	if (ram[0xFF07] & 0x04) == 0{
		return 0
	}
	switch ram[0xFF07] & 0x03{
		case 0x00:
			return CPU_FREQ/4096
		case 0x01:
			return CPU_FREQ/262144
		case 0x02:
			return CPU_FREQ/65536
		case 0x03:
			return CPU_FREQ/16384
		default:
			return 0
	}
}
func incTimer() {
	if tfreq := getTimerFreq(); (tfreq != 0) && ((int(CicleCounter/tfreq) - LastTimer) >= 1) {
		LastTimer = int(CicleCounter/tfreq)
		ram[0xFF05]++
		if ram[0xFF05] == 0 {
			setInterruptsFlag(TIMER)
			ram[0xFF05] = ram[0xFF06]
			//fmt.Println("TIMER INTERRUPT")
		}
	}
	if int(CicleCounter/(CPU_FREQ/16384)) - LastDivTimer >= 1 {
		LastDivTimer = int(CicleCounter/(CPU_FREQ/16384))
		ram[0xFF04]++
	}
}