package gameboygo

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
		ime = false
		clearInterruptsFlag(V_BLANK)
		write16bits(regs.sp, regs.pc)
		regs.sp -= 2
		regs.pc = 0x40

	}
	if (toExec & uint8(LCD_STAT)) != 0 {
		ime = false
		clearInterruptsFlag(LCD_STAT)
		write16bits(regs.sp, regs.pc)
		regs.sp -= 2
		regs.pc = 0x48
	}
	if (toExec & uint8(TIMER)) != 0 {
		ime = false
		clearInterruptsFlag(TIMER)
		write16bits(regs.sp, regs.pc)
		regs.sp -= 2
		regs.pc = 0x50
	}
	if (toExec & uint8(SERIAL)) != 0 {
		ime = false
		clearInterruptsFlag(SERIAL)
		write16bits(regs.sp, regs.pc)
		regs.sp -= 2
		regs.pc = 0x58
	}
	if (toExec & uint8(JOYPAD)) != 0 {
		ime = false
		clearInterruptsFlag(JOYPAD)
		write16bits(regs.sp, regs.pc)
		regs.sp -= 2
		regs.pc = 0x60
	}
}

func setInterruptsFlag(in interrupt) {
	*InterruptFlag |= uint8(in) 
}
func clearInterruptsFlag(in interrupt) {
	*InterruptFlag = *InterruptFlag &^ uint8(in)
}