package main

//import "fmt"

type interrupt uint8
const(
	V_BLANK interrupt = 1 << iota
	LCD_STAT
	TIMER  
	SERIAL 
	JOYPAD
)
var InterruptFlag 	= 	&ram[0xFF0F]
var InterruptEnable =	&ram[0xFFFF]

func InterruptExec() {
	if !ime{
		return
	}
	var toExec uint8 = *InterruptEnable & *InterruptFlag
	if (toExec & uint8(V_BLANK)) != 0 {
		//fmt.Println("V_BLANK")
		ime = false
		clearInterruptsFlag(V_BLANK)
		regs.sp -= 2
		write16bits(regs.sp, regs.pc)
		regs.pc = 0x40
	} else if (toExec & uint8(LCD_STAT)) != 0 {
		//fmt.Println("LCD_STAT")
		ime = false
		clearInterruptsFlag(LCD_STAT)
		regs.sp -= 2
		write16bits(regs.sp, regs.pc)
		regs.pc = 0x48
	} else if (toExec & uint8(TIMER)) != 0 {
		//fmt.Println("TIMER")
		ime = false
		clearInterruptsFlag(TIMER)
		regs.sp -= 2
		write16bits(regs.sp, regs.pc)
		regs.pc = 0x50
	} else if (toExec & uint8(SERIAL)) != 0 {
		//fmt.Println("SERIAL")
		ime = false
		clearInterruptsFlag(SERIAL)
		regs.sp -= 2
		write16bits(regs.sp, regs.pc)
		regs.pc = 0x58
	} else if (toExec & uint8(JOYPAD)) != 0 {
		//fmt.Println("JOYPAD")
		ime = false
		clearInterruptsFlag(JOYPAD)
		regs.sp -= 2
		write16bits(regs.sp, regs.pc)
		regs.pc = 0x60
	}
}

func setInterruptsFlag(in interrupt) {
	*InterruptFlag |= uint8(in) 
}
func clearInterruptsFlag(in interrupt) {
	*InterruptFlag = *InterruptFlag &^ uint8(in)
}