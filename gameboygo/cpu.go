package main

import (
	"fmt"
)
/*
RES
https://github.com/h3nnn4n/gameboy_documentation
http://marc.rawer.de/Gameboy/Docs/GBCPUman.pdf
http://gbdev.gg8.se/files/docs/mirrors/pandocs.html
http://gbdev.gg8.se/files/docs/GBCribSheet000129.pdf
http://goldencrystal.free.fr/GBZ80Opcodes.pdf
https://www.youtube.com/watch?v=CImyDBJSTsQ
https://cturt.github.io/cinoop.html
*/

type operations func(uint8) int
const(
	CPU_FREQ int = 4194304
)
var ops = [0x100]operations{
	nop,  		//0x00
	ld_bc_aabb,	//0x01
	ld_bc_a,	//0x02
	inc_bc,		//0x03
	inc_b, 		//0x04
	dec_b,		//0x05
	ld_b_xx,	//0x06
	rlca,		//0x07
	ld_aabb_sp,	//0x08
	add_hl_bc,	//0x09
	ld_a_bc,	//0x0A
	dec_bc,		//0x0B
	inc_c,		//0x0C
	dec_c,		//0x0D
	ld_c_xx,	//0x0E
	rrca,		//0x0F
	stop, 		//0x10
	ld_de_aabb,	//0x11
	ld_de_a,	//0x12
	inc_de,		//0x13
	inc_d,		//0x14
	dec_d,		//0x15
	ld_d_xx,	//0x16
	rla,		//0x17
	jr_xx,		//0x18
	add_hl_de,	//0x19
	ld_a_de,	//0x1A
	dec_de,		//0x1B
	inc_e,		//0x1C
	dec_e,		//0x1D
	ld_e_xx,	//0x1E
	rra,		//0x1F
	jr_nz_xx,	//0x20
	ld_hl_aabb,	//0x21
	ld_hli_a,	//0x22
	inc_hl,		//0x23
	inc_h,		//0x24
	dec_h,		//0x25
	ld_h_xx,	//0x26
	daa,		//0x27
	jr_z_xx,	//0x28
	add_hl_hl,	//0x29
	ld_a_hli,	//0x2A
	dec_hl,		//0x2B
	inc_l,		//0x2C
	dec_l,		//0x2D
	ld_l_xx,	//0x2E
	cpl,		//0x2F
	jr_nc_xx,	//0x30
	ld_sp_aabb,	//0x31
	ld_hld_a,	//0x32
	inc_sp,		//0x33
	inc_bhl,	//0x34
	dec_bhl,	//0x35
	ld_hl_xx,	//0x36
	scf,		//0x37
	jr_c_xx,	//0x38
	add_hl_sp,	//0x39
	ld_a_hld,	//0x3A
	dec_sp,		//0x3B
	inc_a,		//0x3C
	dec_a,		//0x3D
	ld_a_xx,	//0x3E
	ccf,		//0x3F
	ld_x_y,		//0x40
 	ld_x_y,		//0x41		
 	ld_x_y,		//0x42		
 	ld_x_y,		//0x43		
 	ld_x_y,		//0x44		
 	ld_x_y,		//0x45		
 	ld_x_y,		//0x46
 	ld_x_y,		//0x47		
 	ld_x_y,		//0x48		
 	ld_x_y,		//0x49		
 	ld_x_y,		//0x4A		
 	ld_x_y,		//0x4B		
 	ld_x_y,		//0x4C		
 	ld_x_y,		//0x4D		
 	ld_x_y,		//0x4E		
 	ld_x_y,		//0x4F		
 	ld_x_y,		//0x50		
 	ld_x_y,		//0x51		
 	ld_x_y,		//0x52		
 	ld_x_y,		//0x53		
 	ld_x_y,		//0x54		
 	ld_x_y,		//0x55		
 	ld_x_y,		//0x56		
 	ld_x_y,		//0x57		
 	ld_x_y,		//0x58		
 	ld_x_y,		//0x59		
 	ld_x_y,		//0x5A		
 	ld_x_y,		//0x5B		
 	ld_x_y,		//0x5C		
 	ld_x_y,		//0x5D		
 	ld_x_y,		//0x5E		
 	ld_x_y,		//0x5F		
 	ld_x_y,		//0x60		
 	ld_x_y,		//0x61		
 	ld_x_y,		//0x62		
 	ld_x_y,		//0x63		
 	ld_x_y,		//0x64		
 	ld_x_y,		//0x65		
 	ld_x_y,		//0x66		
 	ld_x_y,		//0x67		
 	ld_x_y,		//0x68		
 	ld_x_y,		//0x69		
 	ld_x_y,		//0x6A		
 	ld_x_y,		//0x6B		
 	ld_x_y,		//0x6C		
 	ld_x_y,		//0x6D		
 	ld_x_y,		//0x6E		
 	ld_x_y,		//0x6F		
 	ld_x_y,		//0x70		
 	ld_x_y,		//0x71		
 	ld_x_y,		//0x72		
 	ld_x_y,		//0x73		
 	ld_x_y,		//0x74		
 	ld_x_y,		//0x75
 	halt,		//0x76
 	ld_phl_a, 	//0x77
	ld_x_y,		//0x78
	ld_x_y,		//0x79
	ld_x_y,		//0x7A
 	ld_x_y,		//0x7B		
 	ld_x_y,		//0x7C		
 	ld_x_y,		//0x7D		
 	ld_x_y,		//0x7E
 	ld_x_y,		//0x7F
	add_a_x,	//0x80 
	add_a_x,	//0x81 
	add_a_x,	//0x82 
	add_a_x,	//0x83 
	add_a_x,	//0x84 
	add_a_x,	//0x85 
	add_a_x,	//0x86
	add_a_x,	//0x87
	adc_a_x,	//0x88 
	adc_a_x,	//0x89 
	adc_a_x,	//0x8A 
	adc_a_x,	//0x8B 
	adc_a_x,	//0x8C 
	adc_a_x,	//0x8D 
	adc_a_x,	//0x8E
	adc_a_x,	//0x8F
	sub_x,		//0x90 
	sub_x,		//0x91 
	sub_x,		//0x92 
	sub_x,		//0x93 
	sub_x,		//0x94 
	sub_x,		//0x95
	sub_x,		//0x96 
	sub_x,		//0x97
	sbc_a_x,	//0x98 		
	sbc_a_x,	//0x99 		
	sbc_a_x,	//0x9A 		
	sbc_a_x,	//0x9B 		
	sbc_a_x,	//0x9C 		
	sbc_a_x,	//0x9D 		
	sbc_a_x,	//0x9E 
	sbc_a_x,	//0x9F
	and_x,		//0xA0 		
	and_x,		//0xA1 		
	and_x,		//0xA2 		
	and_x,		//0xA3 		
	and_x,		//0xA4 		
	and_x,		//0xA5 		
	and_x,		//0xA6
	and_x,		//0xA7
	xor_x,		//0xA8 
	xor_x,		//0xA9 
	xor_x,		//0xAA 
	xor_x,		//0xAB 
	xor_x,		//0xAC 
	xor_x,		//0xAD 
	xor_x,		//0xAE
	xor_x,		//0xAF  			   		
	or_x,		//0xB0 	
	or_x,		//0xB1 	
	or_x,		//0xB2 	
	or_x,		//0xB3 	
	or_x,		//0xB4 	
	or_x,		//0xB5 	
	or_x,		//0xB6
	or_x,		//0xB7
	cp_x,		//0xB8 
	cp_x,		//0xB9 
	cp_x,		//0xBA 
	cp_x,		//0xBB 
	cp_x,		//0xBC 
	cp_x,		//0xBD 
	cp_x,		//0xBE
	cp_x,		//0xBF
	ret_nz,		//0xC0
	pop_bc,		//0xC1
	jp_nz_aabb,	//0xC2
	jp_aabb,	//0xC3
	call_nz_aabb,//0xC4
	push_bc,	//0xC5
	add_a_xx,	//0xC6
	rst_oo,		//0xC7
	ret_z,		//0xC8
	ret,		//0xC9
	jp_z_aabb,	//0xCA
	many_ops,	//0xCB
	call_z_aabb,//0xCC
	call_aabb,	//0xCD
	adc_a_xx,	//0xCE
	rst_08,		//0xCF
	ret_nc,		//0xD0
	pop_de,		//0xD1
	jp_nc_aabb,	//0xD2
	unknown,	//0xD3
	call_nc_aabb,//0xD4
	push_de,	//0xD5
	sub_xx,		//0xD6
	rst_10,		//0xD7
	ret_c,		//0xD8
	reti,		//0xD9
	jp_c_aabb,	//0xDA
	unknown,	//0xDB
	call_c_aabb,//0xDC
	unknown,	//0xDD
	sbc_a_xx,	//0xDE
	rst_18,		//0xDF
	ld_xx_a,	//0xE0
	pop_hl,		//0xE1
	ld_c_a,		//0xE2
	unknown,	//0xE3
	unknown,	//0xE4
	push_hl,	//0xE5
	and_xx,		//0xE6
	rst_20,		//0xE7
	add_sp_xx,	//0xE8
	jp_hl,		//0xE9
	ld_aabb_a,	//0xEA
	unknown,	//0xEB
	unknown,	//0xEC
	unknown,	//0xED
	xor_xx,		//0xEE
	rst_28,		//0xEF
	ldh_a_xx,	//0xF0
	pop_af,		//0xF1
	ld_a_c,		//0xF2
	di,			//0xF3
	unknown,	//0xF4
	push_af,	//0xF5
	or_xx,		//0xF6
	rst_30,		//0xF7
	ld_hl_sp,	//0xF8
	ld_sp_hl,	//0xF9
	ld_a_aabb,	//0xFA
	ei,			//0xFB
	unknown,	//0xFC
	unknown,	//0xFD
	cp_xx,		//0xFE
	rst_38,}	//0xFF

var (
	regs 		 registers
	halted 		 bool
	LastTimer 	 int
	LastDivTimer int
	CicleCounter int
	cicles_op	 int
	ime  		 bool
	debugcount   int
)

func Reset() {
	CicleCounter = 0
	cicles_op    = 0
	ime = true
	halted = false
	LastTimer = 0
	regs.a = 0x01
	regs.f = 0xB0
	regs.b = 0x00
	regs.c = 0x13
	regs.d = 0x00
	regs.e = 0xD8
	regs.h = 0x01
	regs.l = 0x4D
	regs.sp = 0xFFFE
	regs.pc = 0x100
	writeByte(0xFF05, 0x00)
	writeByte(0xFF06, 0x00)
	writeByte(0xFF07, 0x00)
	writeByte(0xFF10, 0x80)
	writeByte(0xFF11, 0xBF)
    writeByte(0xFF12, 0xF3)
    writeByte(0xFF14, 0xBF)
    writeByte(0xFF16, 0x3F)
    writeByte(0xFF17, 0x00)
    writeByte(0xFF19, 0xBF)
    writeByte(0xFF1A, 0x7F)
    writeByte(0xFF1B, 0xFF)
    writeByte(0xFF1C, 0x9F)
    writeByte(0xFF1E, 0xBF)
    writeByte(0xFF20, 0xFF)
    writeByte(0xFF21, 0x00)
    writeByte(0xFF22, 0x00)
    writeByte(0xFF23, 0xBF)
    writeByte(0xFF24, 0x77)
    writeByte(0xFF25, 0xF3)
    writeByte(0xFF26, 0xF1)	//0xF0-SGB
    writeByte(0xFF40, 0x91)
    writeByte(0xFF42, 0x00)
    writeByte(0xFF43, 0x00)
    writeByte(0xFF45, 0x00)
    writeByte(0xFF47, 0xFC)
    writeByte(0xFF48, 0xFF)
    writeByte(0xFF49, 0xFF)
    writeByte(0xFF4A, 0x00)
    writeByte(0xFF4B, 0x00)
    writeByte(0xFFFF, 0x00)
}

func Execute() {
	op := readByte(regs.pc)
	regs.pc++
	cicles_op = ops[op](op)
	CicleCounter += cicles_op 
	incTimer(cicles_op)
	InterruptExec()
}

func ExecuteDebug() {
	var input string
	var op = readByte(regs.pc)
	regs.pc++
	cicles_op = ops[op](op)
	//opstats[op]++
	CicleCounter += cicles_op
	incTimer(cicles_op)
	InterruptExec()
	if debugcount == 0 {
		fmt.Printf("%v\n", &regs)
		input = ""
		fmt.Scanln(&input)
		if input == "s" {
			debugcount = 100
		} else if input == "m"{
			debugcount = 1000
		} else if input == "h"{
			debugcount = 400000
		} else if input == "continue"{
			debugcount = -1
		} else if input == "show" {
			fmt.Scanln(&input)
			var addrtoprint int 
			if _, err := fmt.Sscanf(input, "%4X", &addrtoprint); err == nil {
    			fmt.Println(readByte(uint16(addrtoprint)))
			}
			fmt.Scanln(&input)
		}
		fmt.Print("\033[H\033[2J")
	} else if debugcount == -1{
		
	} else{
		debugcount--
	}
}

func rst_38(b uint8) int{
	//fmt.Printf("RST 38: %X\n", b)
	regs.sp -= 2
	write16bits(regs.sp, regs.pc)
	regs.pc = 0x0038
	return 16
}
func cp_xx(b uint8) int{
	var val uint8 = readByte(regs.pc)
	regs.pc++
	//fmt.Printf("CP %X: %X\n", val, b)
	regs.setFlags(SUBTRACT)
	if regs.a == val{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	if (val & 0x0F) > (regs.a & 0x0F) {
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}
	if val > regs.a{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}
	return 8
}
func ei(b uint8) int{
	//fmt.Printf("EI: %X\n", b)
	ime = true
	return 4
	//interrupts are enabled after instruction after EI is executed
}
func ld_a_aabb(b uint8) int{
	var address uint16 = read16bits(regs.pc)
	regs.pc += 2  
	var val uint8 = readByte(address)
	//fmt.Printf("LD a, %X: %X\n", address, b)
	regs.a = val
	return 16
}
func ld_sp_hl(b uint8) int{
	//fmt.Printf("LD sp, hl: %X\n", b)
	regs.sp = regs.hl_read()
	return 8
}
func ld_hl_sp(b uint8) int{
	//fmt.Printf("LDHL sp, n: %X\n", b)
	regs.clearFlags(ZERO|SUBTRACT)
	var val int8 = int8(readByte(regs.pc))
	regs.pc++
	var res uint32 = uint32(int16(regs.sp) + int16(val))
	if ((regs.sp ^ uint16(val) ^ uint16(res&0xFFFF)) & 0x010) == 0x010{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}
	if ((regs.sp ^ uint16(val) ^ uint16(res&0xFFFF)) & 0x100) == 0x100 {
		regs.setFlags(CARRY)
	} else {
		regs.clearFlags(CARRY)
	}
	regs.hl_write(uint16(0x00FFFF&res))
	return 12
}
func rst_30(b uint8) int{
	//fmt.Printf("RST 30: %X\n", b)
	regs.sp -= 2
	write16bits(regs.sp, regs.pc)
	regs.pc =0x0030
	return 16
}
func or_xx(b uint8) int{
	regs.clearFlags(SUBTRACT|HALFCARRY|CARRY)
	var val uint8 = readByte(regs.pc)
	regs.pc++
	//fmt.Printf("OR %X: %X\n", val, b)
	regs.a |= val
	if regs.a != 0{
		regs.clearFlags(ZERO)
	} else{
		regs.setFlags(ZERO)
	}
	return 8
}
func push_af(b uint8) int{
	//fmt.Printf("PUSH af: %X\n", b)
	regs.sp -= 2
	write16bits(regs.sp, regs.af_read())
	return 16
}
func di(b uint8) int{
	//fmt.Printf("DI: %X\n", b)
	ime = false
	return 4
	//interrupts are disabled after instruction after DI is executed.
}
func ld_a_c(b uint8) int{
	//fmt.Printf("LD a, (0xFF00 + c): %X\n", b)
	regs.a = readByte(0xFF00 + uint16(regs.c))
	return 8
}
func pop_af(b uint8) int{
	//fmt.Printf("POP af: %X\n", b)
	regs.af_write(read16bits(regs.sp))
	regs.sp += 2
	return 12
}
func ldh_a_xx(b uint8) int{
	var val uint8 = readByte(regs.pc)
	regs.pc++
	//fmt.Printf("LDH a, %X: %X\n", val, b)
	regs.a = readByte(0xFF00 + uint16(0x00FF&val))
	return 12
}
func rst_28(b uint8) int{
	//fmt.Printf("RST 28: %X\n", b)
	regs.sp -= 2
	write16bits(regs.sp, regs.pc)
	regs.pc =0x0028
	return 16
}
func xor_xx(b uint8) int{
	regs.clearFlags(SUBTRACT|HALFCARRY|CARRY)
	var val uint8 = readByte(regs.pc)
	regs.pc++
	//fmt.Printf("XOR %X: %X\n", val, b)
	regs.a ^= val
	if regs.a != 0{
		regs.clearFlags(ZERO)
	} else{
		regs.setFlags(ZERO)
	}
	return 8
}
func ld_aabb_a(b uint8) int{
	var dest uint16 = read16bits(regs.pc)
	regs.pc += 2
	//fmt.Printf("LD %X, a: %X\n", dest, b)
	writeByte(dest, regs.a)
	return 16
}
func jp_hl(b uint8) int{
	//fmt.Printf("JP hl: %X\n", b)
	regs.pc = regs.hl_read()
	return 4
}
func add_sp_xx(b uint8) int{
	var val int8 = int8(readByte(regs.pc))
	regs.pc++
	//fmt.Printf("ADD sp, %X: %X\n", val, b)
	regs.clearFlags(ZERO|SUBTRACT)
	var result uint32 = uint32(int16(regs.sp) + int16(val))
	if ((regs.sp ^ uint16(val) ^ (uint16(result& 0xFFFF))) & 0x100) == 0x100{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}
	if ((regs.sp ^ uint16(val) ^ (uint16(result& 0xFFFF))) & 0x010) == 0x010{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}
	regs.sp = uint16(result & 0x0000FFFF)
	return 16
}
func rst_20(b uint8) int{
	//fmt.Printf("RST 20: %X\n", b)
	regs.sp -= 2
	write16bits(regs.sp, regs.pc)
	regs.pc =0x0020
	return 16
}
func and_xx(b uint8) int{
	var val uint8 = readByte(regs.pc)
	regs.pc++
	//fmt.Printf("AND %X: %X\n", val, b)
	regs.clearFlags(SUBTRACT|CARRY)
	regs.setFlags(HALFCARRY)
	regs.a &= val
	if regs.a != 0{
		regs.clearFlags(ZERO)
	} else{
		regs.setFlags(ZERO)
	}
	return 8
}
func push_hl(b uint8) int{
	//fmt.Printf("PUSH hl: %X\n", b)
	regs.sp -= 2
	write16bits(regs.sp, regs.hl_read())
	return 16
}
func ld_c_a(b uint8) int{
	//fmt.Printf("LD (c), a: %X\n", b)
	writeByte(0xFF00+uint16(regs.c), regs.a)
	return 8
}
func pop_hl(b uint8) int{
	//fmt.Printf("POP hl: %X\n", b)
	regs.hl_write(read16bits(regs.sp))
	regs.sp += 2
	return 12
}
func ld_xx_a(b uint8) int{
	var val uint16 = uint16(readByte(regs.pc))
	regs.pc++
	//fmt.Printf("LDH (0xFF00 + %X), a: %X\n", val, b)
	writeByte(0xFF00+val, regs.a)
	return 12
}
func rst_18(b uint8) int{
	//fmt.Printf("RST 18: %X\n", b)
	regs.sp -= 2
	write16bits(regs.sp, regs.pc)
	regs.pc =0x0018
	return 16
}
func sbc_a_xx(b uint8) int{
	var val uint16 = uint16(readByte(regs.pc))
	regs.pc++
	//fmt.Printf("SBC a, %X: %X\n", val, b)
	regs.setFlags(SUBTRACT)
	var carry uint8 = 0
	if regs.getFlag(CARRY){
		carry = 1
	}
	if val + uint16(carry) > uint16(regs.a) {
		regs.setFlags(CARRY)
	} else {
		regs.clearFlags(CARRY)
	}
	if (val & 0x0F) + uint16(carry) > (uint16(regs.a) & 0x0F) {
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.a = uint8(uint16(regs.a) - val - uint16(carry))

	if regs.a == 0 {
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	return 8
}
func call_c_aabb(b uint8) int{
	var dest uint16 = read16bits(regs.pc)
	regs.pc += 2
	//fmt.Printf("CALL c, %X: %X\n", dest, b)
	if regs.getFlag(CARRY){
		regs.sp -= 2
		write16bits(regs.sp, regs.pc)
		regs.pc = dest
		return 24
	}
	return 12
}
func jp_c_aabb(b uint8) int{
	var dest uint16 = read16bits(regs.pc)
	//fmt.Printf("JP c, %X: %X\n", dest, b)
	regs.pc += 2
	if regs.getFlag(CARRY){
		regs.pc = dest
		return 16
	}
	return 12
}
func reti(b uint8) int{
	//fmt.Printf("RETI: %X\n", b)
	regs.pc = read16bits(regs.sp)
	regs.sp += 2
	ime = true
	return 16
}
func ret_c(b uint8) int{
	//fmt.Printf("RET c: %X\n", b)
	if regs.getFlag(CARRY) {
		dest := read16bits(regs.sp)
		regs.sp += 2
		regs.pc = dest
		return 20
	}
	return 8
}
func rst_10(b uint8) int{
	//fmt.Printf("RST 10: %X\n", b)
	regs.sp -= 2
	write16bits(regs.sp, regs.pc)
	regs.pc =0x0010
	return 16
}
func sub_xx(b uint8) int{
	var val uint8 = readByte(regs.pc)
	regs.pc++
	//fmt.Printf("SUB %X: %X\n", val, b)
	regs.setFlags(SUBTRACT)
	if val > regs.a {
		regs.setFlags(CARRY)
	} else {
		regs.clearFlags(CARRY)
	}
	if (val & 0x0F) > (regs.a & 0x0F) {
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.a -= val

	if regs.a != 0 {
		regs.clearFlags(ZERO)
	} else{
		regs.setFlags(ZERO)
	}
	return 8
}
func push_de(b uint8) int{
	//fmt.Printf("PUSH de: %X\n", b)
	regs.sp -= 2
	write16bits(regs.sp, regs.de_read())
	return 16
}
func call_nc_aabb(b uint8) int{
	var dest uint16 = read16bits(regs.pc)
	regs.pc += 2
	//fmt.Printf("CALL nc, %X: %X\n", dest, b)
	if !regs.getFlag(CARRY){
		regs.sp -= 2
		write16bits(regs.sp, regs.pc)
		regs.pc = dest
		return 24
	}
	return 12
}
func jp_nc_aabb(b uint8) int{
	var dest uint16 = read16bits(regs.pc)
	//fmt.Printf("JP nc, %X: %X\n", dest, b)
	regs.pc += 2
	if !regs.getFlag(CARRY){
		regs.pc = dest
		return 16
	}
	return 12
}
func pop_de(b uint8) int{
	//fmt.Printf("POP de: %X\n", b)
	regs.de_write(read16bits(regs.sp))
	regs.sp += 2
	return 12
}
func ret_nc(b uint8) int{
	//fmt.Printf("RET nc: %X\n", b)
	if !regs.getFlag(CARRY) {
		dest := read16bits(regs.sp)
		regs.sp += 2
		regs.pc = dest
		return 20
	}
	return 8
}
func rst_08(b uint8) int{
	//fmt.Printf("RST 08: %X\n", b)
	regs.sp -= 2
	write16bits(regs.sp, regs.pc)
	regs.pc = 0x0008
	return 16
}
func adc_a_xx(b uint8) int{
	var val uint8 = readByte(regs.pc)
	regs.pc++
	//fmt.Printf("ADC a, %X: %X\n", val, b)
	regs.clearFlags(SUBTRACT)
	var result uint16 = 0
	var carry uint8  = 0
	if regs.getFlag(CARRY){
		carry = 1
	}
	result = uint16(regs.a) + uint16(val) + uint16(carry)
	if result > 0xFF{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}
	if ((regs.a & 0x0F) + (val & 0x0F) + carry) > 0x0F{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.a = uint8(result & 0x00FF)
	
	if regs.a != 0{
		regs.clearFlags(ZERO)
	} else{
		regs.setFlags(ZERO)
	}
	return 8
}
func call_aabb(b uint8) int{
	var dest uint16 = read16bits(regs.pc)
	regs.pc += 2
	//fmt.Printf("CALL %X: %X\n", dest, b)
	regs.sp -= 2
	write16bits(regs.sp, regs.pc)
	regs.pc = dest
	return 24
}
func call_z_aabb(b uint8) int{
	var dest uint16 = read16bits(regs.pc)
	regs.pc += 2
	//fmt.Printf("CALL z, %X: %X\n", dest, b)
	if regs.getFlag(ZERO){
		regs.sp -= 2
		write16bits(regs.sp, regs.pc)
		regs.pc = dest
		return 24
	}
	return 12
}
func many_ops(b uint8) int{
	op := readByte(regs.pc)
	//fmt.Printf("CB %X\n", op)
	regs.pc++
	return cb_op[op](op)
}
func jp_z_aabb(b uint8) int{
	var dest uint16 = read16bits(regs.pc)
	//fmt.Printf("JP z, %X: %X\n",dest, b)
	regs.pc += 2
	if regs.getFlag(ZERO){
		regs.pc = dest
		return 16
	}
	return 12
}
func ret(b uint8) int{
	//fmt.Printf("RET: %X\n", b)
	dest := read16bits(regs.sp)
	regs.sp += 2
	regs.pc = dest
	return 16
}
func ret_z(b uint8) int{
	//fmt.Printf("RET z: %X\n", b)
	if regs.getFlag(ZERO) {
		dest := read16bits(regs.sp)
		regs.sp += 2
		regs.pc = dest
		return 20
	}
	return 8
}
func rst_oo(b uint8) int{
	//fmt.Printf("RST 00: %X\n", b)
	regs.sp -= 2
	write16bits(regs.sp, regs.pc)
	regs.pc =0x0000
	return 16
}
func add_a_xx(b uint8) int{
	var val uint8 = readByte(regs.pc)
	regs.pc++
	//fmt.Printf("ADD a, %X: %X\n", val, b)
	regs.clearFlags(SUBTRACT)
	var result uint16
	result = uint16(regs.a) + uint16(val)
	if ((regs.a & 0x0F) + (val & 0x0F)) > 0x0F{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}
	if (result & 0xFF00) > 0{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}

	regs.a = uint8(result & 0x00FF)

	if regs.a != 0{
		regs.clearFlags(ZERO)
	} else{
		regs.setFlags(ZERO)
	}
	return 8
}
func push_bc(b uint8) int{
	//fmt.Printf("PUSH bc: %X\n", b)
	regs.sp -= 2
	write16bits(regs.sp, regs.bc_read())
	return 16
}
func call_nz_aabb(b uint8) int{
	var dest uint16 = read16bits(regs.pc)
	regs.pc += 2
	//fmt.Printf("CALL nz, %X: %X\n", dest, b)
	if !regs.getFlag(ZERO){
		regs.sp -= 2
		write16bits(regs.sp, regs.pc)
		regs.pc = dest
		return 24
	}
	return 12
}
func jp_aabb(b uint8) int{
	var dest uint16 = read16bits(regs.pc)
	//fmt.Printf("JP %X: %X\n", dest, b)
	regs.pc = dest
	return 16
}
func jp_nz_aabb(b uint8) int{
	var dest uint16 = read16bits(regs.pc)
	//fmt.Printf("JP nz, %X: %X\n",dest, b)
	regs.pc += 2
	if !regs.getFlag(ZERO){
		regs.pc = dest
		return 16
	}
	return 12
}
func pop_bc(b uint8) int{
	//fmt.Printf("POP bc: %X\n", b)
	regs.bc_write(read16bits(regs.sp))
	regs.sp += 2
	return 12
}
func ret_nz(b uint8) int{
	//fmt.Printf("RET nz: %X\n", b)
	if !regs.getFlag(ZERO) {
		dest := read16bits(regs.sp)
		regs.sp += 2
		regs.pc = dest
		return 20
	}
	return 8
}
func cp_x(b uint8) int{
	regs.setFlags(SUBTRACT)
	var val uint8
	var cicles int = 0
	switch b{
	case 0xB8:
		//fmt.Printf("CP b: %X\n", b)
		val = regs.b
	case 0xB9:
		//fmt.Printf("CP c: %X\n", b)
		val = regs.c
	case 0xBA:
		//fmt.Printf("CP d: %X\n", b)
		val = regs.d
	case 0xBB:
		//fmt.Printf("CP e: %X\n", b)
		val = regs.e
	case 0xBC:
		//fmt.Printf("CP h: %X\n", b)
		val = regs.h
	case 0xBD:
		//fmt.Printf("CP l: %X\n", b)
		val = regs.l
	case 0xBE:
		//fmt.Printf("CP (hl): %X\n", b)
		val = readByte(regs.hl_read())
		cicles += 4
	case 0xBF:
		//fmt.Printf("CP a: %X\n", b)
		val = regs.a 
	}
	if regs.a == val{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	if (val & 0x0F) > (regs.a & 0x0F) {
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}
	if val > regs.a{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}
	return cicles + 4
}
func or_x(b uint8) int{
	var cicles int = 0
	regs.clearFlags(SUBTRACT|HALFCARRY|CARRY)
	switch b{
	case 0xB0:
		//fmt.Printf("OR b: %X\n", b)
		regs.a |= regs.b
	case 0xB1:
		//fmt.Printf("OR c: %X\n", b)
		regs.a |= regs.c
	case 0xB2:
		//fmt.Printf("OR d: %X\n", b)
		regs.a |= regs.d
	case 0xB3:
		//fmt.Printf("OR e: %X\n", b)
		regs.a |= regs.e
	case 0xB4:
		//fmt.Printf("OR h: %X\n", b)
		regs.a |= regs.h
	case 0xB5:
		//fmt.Printf("OR l: %X\n", b)
		regs.a |= regs.l
	case 0xB6:
		//fmt.Printf("OR (hl): %X\n", b)
		regs.a |= readByte(regs.hl_read())
		cicles += 4
	case 0xB7:
		//fmt.Printf("OR a: %X\n", b)
		regs.a |= regs.a
	}
	if regs.a != 0{
		regs.clearFlags(ZERO)
	} else{
		regs.setFlags(ZERO)
	}
	return cicles + 4
}
func xor_x(b uint8) int{
	var cicles int = 0
	regs.clearFlags(SUBTRACT|HALFCARRY|CARRY)
	switch b{
	case 0xA8:
		//fmt.Printf("XOR b: %X\n", b)
		regs.a ^= regs.b
	case 0xA9:
		//fmt.Printf("XOR c: %X\n", b)
		regs.a ^= regs.c
	case 0xAA:
		//fmt.Printf("XOR d: %X\n", b)
		regs.a ^= regs.d
	case 0xAB:
		//fmt.Printf("XOR e: %X\n", b)
		regs.a ^= regs.e
	case 0xAC:
		//fmt.Printf("XOR h: %X\n", b)
		regs.a ^= regs.h
	case 0xAD:
		//fmt.Printf("XOR l: %X\n", b)
		regs.a ^= regs.l
	case 0xAE:
		//fmt.Printf("XOR (hl): %X\n", b)
		regs.a ^= readByte(regs.hl_read())
		cicles += 4
	case 0xAF:
		//fmt.Printf("XOR a: %X\n", b)
		regs.a ^= regs.a
	}
	if regs.a != 0{
		regs.clearFlags(ZERO)
	} else{
		regs.setFlags(ZERO)
	}
	return cicles + 4
}
func and_x(b uint8) int{
	var cicles int = 0
	regs.clearFlags(SUBTRACT|CARRY)
	regs.setFlags(HALFCARRY)
	switch b{
	case 0xA0:
		//fmt.Printf("AND b: %X\n", b)
		regs.a &= regs.b
	case 0xA1:
		//fmt.Printf("AND c: %X\n", b)
		regs.a &= regs.c
	case 0xA2:
		//fmt.Printf("AND d: %X\n", b)
		regs.a &= regs.d
	case 0xA3:
		//fmt.Printf("AND e: %X\n", b)
		regs.a &= regs.e
	case 0xA4:
		//fmt.Printf("AND h: %X\n", b)
		regs.a &= regs.h
	case 0xA5:
		//fmt.Printf("AND l: %X\n", b)
		regs.a &= regs.l
	case 0xA6:
		//fmt.Printf("AND (hl): %X\n", b)
		regs.a &= readByte(regs.hl_read())
		cicles += 4
	case 0xA7:
		//fmt.Printf("AND a: %X\n", b)
		regs.a &= regs.a
	}
	if regs.a != 0{
		regs.clearFlags(ZERO)
	} else{
		regs.setFlags(ZERO)
	}
	return cicles + 4
}
func sbc_a_x(b uint8) int{
	var cicles int = 0
	var val uint16
	var carry uint16
	regs.setFlags(SUBTRACT)
	if regs.getFlag(CARRY){
		carry = 1
	}
	switch b{
	case 0x98:
		//fmt.Printf("SBC a, b: %X\n", b)
		val = uint16(regs.b)
	case 0x99:
		//fmt.Printf("SBC a, c: %X\n", b)
		val = uint16(regs.c)
	case 0x9A:
		//fmt.Printf("SBC a, d: %X\n", b)
		val = uint16(regs.d)
	case 0x9B:
		//fmt.Printf("SBC a, e: %X\n", b)
		val = uint16(regs.e)
	case 0x9C:
		//fmt.Printf("SBC a, h: %X\n", b)
		val = uint16(regs.h)
	case 0x9D:
		//fmt.Printf("SBC a, l: %X\n", b)
		val = uint16(regs.l)
	case 0x9E:
		//fmt.Printf("SBC a, (hl): %X\n", b)
		val = uint16(readByte(regs.hl_read()))
		cicles += 4
	case 0x9F:
		//fmt.Printf("SBC a, a: %X\n", b)
		val = uint16(regs.a) 
	}
	if (val + carry) > uint16(regs.a) {
		regs.setFlags(CARRY)
	} else {
		regs.clearFlags(CARRY)
	}
	if (val & 0x0F) + carry > (uint16(regs.a) & 0x0F) {
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.a = uint8(uint16(regs.a) - val - carry)

	if regs.a != 0 {
		regs.clearFlags(ZERO)
	} else{
		regs.setFlags(ZERO)
	}
	return cicles + 4
}
func sub_x(b uint8) int{
	var cicles int = 0
	var val uint8
	regs.setFlags(SUBTRACT)
	switch b{
	case 0x90:
		//fmt.Printf("SUB b: %X\n", b)
		val = regs.b
	case 0x91:
		//fmt.Printf("SUB c: %X\n", b)
		val = regs.c
	case 0x92:
		//fmt.Printf("SUB d: %X\n", b)
		val = regs.d
	case 0x93:
		//fmt.Printf("SUB e: %X\n", b)
		val = regs.e
	case 0x94:
		//fmt.Printf("SUB h: %X\n", b)
		val = regs.h
	case 0x95:
		//fmt.Printf("SUB l: %X\n", b)
		val = regs.l
	case 0x96:
		//fmt.Printf("SUB (hl): %X\n", b)
		val = readByte(regs.hl_read())
		cicles += 4
	case 0x97:
		//fmt.Printf("SUB a: %X\n", b)
		val = regs.a 
	}
	if val > regs.a {
		regs.setFlags(CARRY)
	} else {
		regs.clearFlags(CARRY)
	}
	if (val & 0x0F) > (regs.a & 0x0F) {
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.a -= val

	if regs.a != 0 {
		regs.clearFlags(ZERO)
	} else{
		regs.setFlags(ZERO)
	}
	return cicles + 4
}
func adc_a_x(b uint8) int{
	regs.clearFlags(SUBTRACT)
	var cicles int = 0
	var result uint16 = 0
	var tmp uint8
	var carry uint16 = 0
	if regs.getFlag(CARRY){
		carry = 1
	}
	switch b{
	case 0x88:
		//fmt.Printf("ADC a, b: %X\n", b)
		tmp = regs.b
		result = uint16(regs.a) + uint16(regs.b) + carry
	case 0x89:
		//fmt.Printf("ADC a, c: %X\n", b)
		tmp = regs.c
		result = uint16(regs.a) + uint16(regs.c) + carry
	case 0x8A:
		//fmt.Printf("ADC a, d: %X\n", b)
		tmp = regs.d
		result = uint16(regs.a) + uint16(regs.d) + carry
	case 0x8B:
		//fmt.Printf("ADC a, e: %X\n", b)
		tmp = regs.e
		result = uint16(regs.a) + uint16(regs.e) + carry
	case 0x8C:
		//fmt.Printf("ADC a, h: %X\n", b)
		tmp = regs.h
		result = uint16(regs.a) + uint16(regs.h) + carry
	case 0x8D:
		//fmt.Printf("ADC a, l: %X\n", b)
		tmp = regs.l
		result = uint16(regs.a) + uint16(regs.l) + carry
	case 0x8E:
		//fmt.Printf("ADC a, (hl): %X\n", b)
		tmp = readByte(regs.hl_read())
		result = uint16(regs.a) + uint16(tmp) 	 + carry
		cicles += 4
	case 0x8F:
		//fmt.Printf("ADC a, a: %X\n", b)
		tmp = regs.a
		result = uint16(regs.a) + uint16(regs.a) + carry
	}
	if (result & 0xFF00) > 0{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}
	if ((uint16(regs.a) & 0x0F) + (uint16(tmp) & 0x0F) + carry) > 0x0F{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.a = uint8(result & 0x00FF)
	
	if regs.a != 0{
		regs.clearFlags(ZERO)
	} else{
		regs.setFlags(ZERO)
	}
	return cicles + 4
}
func add_a_x(b uint8) int{
	regs.clearFlags(SUBTRACT)
	var cicles int = 0
	var result uint32
	var tmp uint8
	switch b{
	case 0x80:
		//fmt.Printf("ADD a, b: %X\n", b)
		tmp = regs.b
		result = uint32(regs.a) + uint32(regs.b)
	case 0x81:
		//fmt.Printf("ADD a, c: %X\n", b)
		tmp = regs.c
		result = uint32(regs.a) + uint32(regs.c)
	case 0x82:
		//fmt.Printf("ADD a, d: %X\n", b)
		tmp = regs.d
		result = uint32(regs.a) + uint32(regs.d)
	case 0x83:
		//fmt.Printf("ADD a, e: %X\n", b)
		tmp = regs.e
		result = uint32(regs.a) + uint32(regs.e)
	case 0x84:
		//fmt.Printf("ADD a, h: %X\n", b)
		tmp = regs.h
		result = uint32(regs.a) + uint32(regs.h)
	case 0x85:
		//fmt.Printf("ADD a, l: %X\n", b)
		tmp = regs.l
		result = uint32(regs.a) + uint32(regs.l)
	case 0x86:
		//fmt.Printf("ADD a, (hl): %X\n", b)
		tmp = readByte(regs.hl_read())
		result = uint32(regs.a) + uint32(tmp)
		cicles += 4
	case 0x87:
		//fmt.Printf("ADD a, a: %X\n", b)
		tmp = regs.a
		result = uint32(regs.a) + uint32(regs.a)
	}
	if ((regs.a & 0x0F) + (tmp & 0x0F)) > 0x0F{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}
	if (result & 0xFF00) > 0{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}

	regs.a = uint8(result & 0x00FF)

	if regs.a != 0{
		regs.clearFlags(ZERO)
	} else{
		regs.setFlags(ZERO)
	}
	return cicles + 4
}
func halt(b uint8) int{
	//fmt.Printf("HALT: %X\n", b)
	halted = true
	return 4
}
func ld_phl_a(b uint8) int{
	//fmt.Printf("LD (hl), a: %X\n", b)
	writeByte(regs.hl_read(), regs.a)
	return 8
}
func ld_x_y(b uint8) int{
	//fmt.Printf("ld x, y: %X\n", b)
	var cicles int = 0
	switch b{
	case 0x40:
		regs.b = regs.b
	case 0x41:
		regs.b = regs.c
		//4 cicles
	case 0x42:
		regs.b = regs.d
		//4 cicles
	case 0x43:
		regs.b = regs.e
		//4cicles
	case 0x44:
		regs.b = regs.h
		//4 cicles
	case 0x45:
		regs.b = regs.l
		//4 cicles
	case 0x46:
		regs.b = readByte(regs.hl_read())
		cicles += 4
	case 0x47:
		regs.b = regs.a
		//4 cicles
	case 0x48:
		regs.c = regs.b
		//4 cicles
	case 0x49:
		regs.c = regs.c
	case 0x4A:
		regs.c = regs.d
		//4 cicles
	case 0x4B:
		regs.c = regs.e
		//4 cicles
	case 0x4C:
		regs.c = regs.h
		//4 cicles
	case 0x4D:
		regs.c= regs.l
		//4 cicles
	case 0x4E:
		regs.c = readByte(regs.hl_read())
		cicles += 4
	case 0x4F:
		regs.c = regs.a
		//4 cicles
	case 0x50:
		regs.d = regs.b
		//4 cicles
	case 0x51:
		regs.d = regs.c
		//4 cicles
	case 0x52:
		regs.d = regs.d
		//4 cicles
	case 0x53:
		regs.d = regs.e
		//4 cicles
	case 0x54:
		regs.d = regs.h
		//4 cicles
	case 0x55:
		regs.d = regs.l
		//4 cicles
	case 0x56:
		regs.d = readByte(regs.hl_read())
		cicles += 4
	case 0x57:
		regs.d = regs.a
		//4 cicles
	case 0x58:
		regs.e = regs.b
		//4 cicles
	case 0x59:
		regs.e = regs.c
		//4 cicles
	case 0x5A:
		regs.e = regs.d
		//4 cicles
	case 0x5B:
		regs.e = regs.e
		//4 cicles
	case 0x5C:
		regs.e = regs.h
		//4 cicles
	case 0x5D:
		regs.e = regs.l
		//4 cicles
	case 0x5E:
		regs.e = readByte(regs.hl_read())
		cicles += 4
	case 0x5F:
		regs.e = regs.a
		//4 cicles
	case 0x60:
		regs.h = regs.b
		//4 cicles
	case 0x61:
		regs.h = regs.c
		//4 cicles
	case 0x62:
		regs.h = regs.d
		//4 cicles
	case 0x63:
		regs.h = regs.e
		//4 cicles
	case 0x64:
		regs.h = regs.h
		//4 cicles
	case 0x65:
		regs.h = regs.l
		//4 cicles
	case 0x66:
		regs.h = readByte(regs.hl_read())
		cicles += 4
	case 0x67:
		regs.h = regs.a
		//4 cicles
	case 0x68:
		regs.l = regs.b
		//4 cicles
	case 0x69:
		regs.l = regs.c
		//4 cicles
	case 0x6A:
		regs.l = regs.d
		//4 cicles
	case 0x6B:
		regs.l = regs.e
		//4 cicles
	case 0x6C:
		regs.l = regs.h
		//4 cicles
	case 0x6D:
		regs.l = regs.l
		//4 cicles
	case 0x6E:
		regs.l = readByte(regs.hl_read())
		cicles += 4
	case 0x6F:
		regs.l = regs.a
		//4 cicles
	case 0x70:
		writeByte(regs.hl_read(), regs.b)
		cicles += 4
	case 0x71:
		writeByte(regs.hl_read(), regs.c)
		cicles += 4
	case 0x72:
		writeByte(regs.hl_read(), regs.d)
		cicles += 4
	case 0x73:
		writeByte(regs.hl_read(), regs.e)
		cicles += 4
	case 0x74:
		writeByte(regs.hl_read(), regs.h)
		cicles += 4
	case 0x75:
		writeByte(regs.hl_read(), regs.l)
		cicles += 4
	case 0x78:
		regs.a = regs.b
		//4 cicles
	case 0x79:
		regs.a = regs.c
		//4 cicles
	case 0x7A:
		regs.a = regs.d
		//4 cicles
	case 0x7B:
		regs.a = regs.e
		//4 cicles
	case 0x7C:
		regs.a = regs.h
		//4 cicles
	case 0x7D:
		regs.a = regs.l
		//4 cicles
	case 0x7E:
		regs.a = readByte(regs.hl_read())
		cicles += 4
	case 0x7F:
		regs.a = regs.a
		//4 cicles
	}
	return cicles + 4
}
func ld_sp_aabb(b uint8) int{
	regs.sp = read16bits(regs.pc)
	regs.pc += 2
	//fmt.Printf("ADD a, %X: %X\n", regs.sp, b)
	return 12
}
func ld_hld_a(b uint8) int{
	//fmt.Printf("LDD (hl), a: %X\n", b)
	writeByte(regs.hl_read(), regs.a)
	regs.hl_write(regs.hl_read()-1)
	return 8
}
func inc_sp(b uint8) int{
	//fmt.Printf("INC sp: %X\n", b)
	regs.sp++
	return 8
}
func inc_bhl(b uint8) int{
	//fmt.Printf("INC hl: %X\n", b)
	var data uint8 = readByte(regs.hl_read())
	if (data & 0x0F) == 0x0F{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	data++
	writeByte(regs.hl_read(), data)

	if data == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.clearFlags(SUBTRACT)
	return 12
}
func dec_bhl(b uint8) int{
	//fmt.Printf("DEC hl: %X\n", b)
	var data uint8 = readByte(regs.hl_read())
	if (data & 0x0F) == 0{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	data--
	writeByte(regs.hl_read(), data)

	if data == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.setFlags(SUBTRACT)
	return 12
}
func ld_hl_xx(b uint8) int{
	val := readByte(regs.pc)
	writeByte(regs.hl_read(), val)
	regs.pc++
	//fmt.Printf("LD hl, %X: %X\n", val, b)
	return 12
}
func scf(b uint8) int{
	//fmt.Printf("SCF: %X\n",b)
	regs.setFlags(CARRY)
	regs.clearFlags(SUBTRACT|HALFCARRY)
	return 4
}
func jr_c_xx(b uint8) int{
	var dest int8 = int8(readByte(regs.pc))
	regs.pc++
	//fmt.Printf("JR c, %X: %X\n", dest, b)
	if regs.getFlag(CARRY) {
		regs.pc = uint16(int16(regs.pc) + int16(dest))
		return 12
	}
	return 8
}
func add_hl_sp(b uint8) int{
	//fmt.Printf("ADD HL, SP: %X\n", b)
	regs.clearFlags(SUBTRACT)
	var result uint32 = uint32(regs.hl_read()) + uint32(regs.sp)

	if (result & 0xFF0000) != 0{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}
	if ((regs.hl_read()^regs.sp^uint16(result&0xFFFF)) & 0x1000) != 0 {
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}
	regs.hl_write(uint16(0x00FFFF&result))
	return 8
}
func ld_a_hld(b uint8) int{
	//fmt.Printf("LDD a, hl: %X\n", b)
	regs.a = readByte(regs.hl_read())
	regs.hl_write(regs.hl_read()-1)
	return 8
}
func dec_sp(b uint8) int{
	//fmt.Printf("DEC sp: %X\n", b)
	regs.sp--
	return 8
}
func inc_a(b uint8) int{
	//fmt.Printf("INC a: %X\n", b)
	if (regs.a & 0x0F) == 0x0F{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.a++

	if regs.a == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.clearFlags(SUBTRACT)
	return 4
}
func dec_a(b uint8) int{
	//fmt.Printf("DEC a: %X\n", b)
	if (regs.a & 0x0F) == 0{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.a--

	if regs.a == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.setFlags(SUBTRACT)
	return 4
}
func ld_a_xx(b uint8) int{
	regs.a = readByte(regs.pc)
	regs.pc++
	//fmt.Printf("LD a, %X: %X\n", regs.a, b)
	return 8
}
func ccf(b uint8) int{
	//fmt.Printf("CCF: %X\n",b)
	regs.clearFlags(SUBTRACT|HALFCARRY)
	if regs.getFlag(CARRY) {
		regs.clearFlags(CARRY)
	} else{
		regs.setFlags(CARRY)
	}
	return 4
} 
func jr_nc_xx(b uint8) int{
	var dest int8 = int8(readByte(regs.pc))
	regs.pc++
	//fmt.Printf("JR nc, %X: %X\n", dest, b)
	if !regs.getFlag(CARRY) {
		regs.pc = uint16(int16(regs.pc) + int16(dest))
		return 12
	}
	return 8
}
func cpl(b uint8) int{
	//fmt.Printf("CPL: %X\n", b)
	regs.setFlags(SUBTRACT|HALFCARRY)
	regs.a = ^regs.a 
	return 4
}
func ld_l_xx(b uint8) int{
	regs.l = readByte(regs.pc)
	regs.pc++
	//fmt.Printf("LD l, %X: %X\n",regs.l, b)
	return 8
}
func dec_l(b uint8) int{
	//fmt.Printf("DEC l: %X\n", b)
	if (regs.l & 0x0F) == 0{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.l--

	if regs.l == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.setFlags(SUBTRACT)
	return 4
}
func inc_l(b uint8) int{
	//fmt.Printf("INC l: %X\n", b)
	if (regs.l & 0x0F) == 0x0F{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.l++

	if regs.l == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.clearFlags(SUBTRACT)
	return 4
}
func dec_hl(b uint8) int{
	//fmt.Printf("DEC hl: %X\n", b)
	regs.hl_write(regs.hl_read() - 1)
	return 8
}
func ld_a_hli(b uint8) int{
	//fmt.Printf("LDI a, (hl): %X\n", b)
	regs.a = readByte(regs.hl_read())
	regs.hl_write(regs.hl_read()+1)
	return 8
}
func add_hl_hl(b uint8) int{
	//fmt.Printf("ADD HL, HL: %X\n", b)
	regs.clearFlags(SUBTRACT)
	var result uint32 = uint32(regs.hl_read()) + uint32(regs.hl_read())

	if (result & 0xFF0000) != 0{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}
	if ((regs.hl_read() ^ regs.hl_read() ^ uint16(result&0xFFFF)) & 0x1000) != 0{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}
	regs.hl_write(uint16(result))
	return 8
}
func jr_z_xx(b uint8) int{
	var dest int8 = int8(readByte(regs.pc))
	regs.pc++
	//fmt.Printf("JR z, %X: %X\n", dest, b)
	if regs.getFlag(ZERO) {
		regs.pc = uint16(int16(regs.pc) + int16(dest))
		return 12
	}
	return 8
}
func daa(b uint8) int{
	//fmt.Printf("DAA: %X\n", b)
	var tmp uint16 = uint16(regs.a)
	if regs.getFlag(SUBTRACT) {
		if regs.getFlag(HALFCARRY) {
			tmp = (tmp-0x06) & 0xFF
		}
		if regs.getFlag(CARRY) {
			tmp = tmp - 0x60
		}
	} else{
		if regs.getFlag(HALFCARRY) || ((tmp & 0x0F) > 9) {
			tmp = tmp + 0x06
		}
		if regs.getFlag(CARRY) || tmp > 0x9F{
			tmp = tmp + 0x60
		}
	}
	if tmp >= 0x100 {
		regs.setFlags(CARRY)
	}
	regs.a = uint8(tmp & 0xFF)
	if regs.a != 0 {
		regs.clearFlags(ZERO)
	} else{
		regs.setFlags(ZERO)
	}
	regs.clearFlags(HALFCARRY)
	return 4
}
func ld_h_xx(b uint8) int{
	regs.h = readByte(regs.pc)
	regs.pc++
	//fmt.Printf("LD h, %X: %X\n",regs.h, b)
	return 8
}
func dec_h(b uint8) int{
	//fmt.Printf("DEC h: %X\n", b)
	if (regs.h & 0x0F) == 0{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.h--

	if regs.h == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.setFlags(SUBTRACT)
	return 4
}
func inc_h(b uint8) int{
	//fmt.Printf("INC h: %X\n", b)
	if (regs.h & 0x0F) == 0x0F{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.h++

	if regs.h == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.clearFlags(SUBTRACT)
	return 4
}
func inc_hl(b uint8) int{
	//fmt.Printf("INC hl: %X\n", b)
	regs.hl_write(regs.hl_read() + 1)
	return 8
}
func ld_hli_a(b uint8) int{
	//fmt.Printf("LDI hl, a: %X\n", b)
	writeByte(regs.hl_read(), regs.a)
	regs.hl_write(regs.hl_read()+1)
	return 8
}
func ld_hl_aabb(b uint8) int{
	regs.hl_write(read16bits(regs.pc))
	regs.pc += 2
	//fmt.Printf("LD hl, %X: %X\n", regs.hl_read(), b)
	return 12
}
func jr_nz_xx(b uint8) int{
	var dest int8 = int8(readByte(regs.pc))
	regs.pc++
	//fmt.Printf("JR nz, %d: %X\n",dest, b)
	if !regs.getFlag(ZERO) {
		regs.pc = uint16(int16(regs.pc) + int16(dest))
		return 12
	}
	return 8
}
func rra(b uint8) int{
	//fmt.Printf("RRA: %X\n", b)
	regs.clearFlags(SUBTRACT|HALFCARRY|ZERO)
	var carry uint8 = 0
	if regs.getFlag(CARRY) {
		carry = 0x80
	}
	if (regs.a & 0x01) != 0{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}

	regs.a = (regs.a >> 1) | carry

	return 4
}
func ld_e_xx(b uint8) int{
	//fmt.Printf("LD e, xx: %X\n", b)
	regs.e = readByte(regs.pc)
	regs.pc++
	return 8
}
func dec_e(b uint8) int{
	//fmt.Printf("DEC e: %X\n", b)
	if (regs.e & 0x0F) == 0{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.e--

	if regs.e == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.setFlags(SUBTRACT)
	return 4
}
func inc_e(b uint8) int{
	//fmt.Printf("INC e: %X\n", b)
	if (regs.e & 0x0F) == 0x0F{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.e++

	if regs.e == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.clearFlags(SUBTRACT)
	return 4
}
func dec_de(b uint8) int{
	//fmt.Printf("DEC de: %X\n", b)
	regs.de_write(regs.de_read() - 1)
	return 8
}
func ld_a_de(b uint8) int{
	//fmt.Printf("LD a, (de): %X\n", b)
	regs.a = readByte(regs.de_read())
	return 8
}
func add_hl_de(b uint8) int{
	//fmt.Printf("ADD HL, DE: %X\n", b)
	regs.clearFlags(SUBTRACT)
	var result uint32 = uint32(regs.hl_read()) + uint32(regs.de_read())

	if (result & 0xFF0000) != 0{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}
	if ((regs.hl_read() ^ regs.de_read() ^ uint16(result&0xFFFF)) & 0x1000) != 0 {
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}
	regs.hl_write(uint16(result & 0x00FFFF))
	return 8
}
func jr_xx(b uint8) int{
	var dest int8 = int8(readByte(regs.pc))
	regs.pc++
	//fmt.Printf("JR %X: %X\n",b,b)
	regs.pc = uint16(int16(regs.pc) + int16(dest))
	return 12
}
func rla(b uint8) int{
	//fmt.Printf("RLA: %X\n", b)
	regs.clearFlags(SUBTRACT|HALFCARRY|ZERO)
	var carry uint8 = 0
	if regs.getFlag(CARRY) {
		carry = 0x01
	}
	if (regs.a & 0x80) != 0{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}
	regs.a = (regs.a << 1) | carry
	return 4
}
func ld_d_xx(b uint8) int{
	//fmt.Printf("LD d, xx: %X\n", b)
	regs.d = readByte(regs.pc)
	regs.pc++
	return 8
}
func dec_d(b uint8) int{
	//fmt.Printf("DEC d: %X\n", b)
	if (regs.d & 0x0F) == 0{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.d--

	if regs.d == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.setFlags(SUBTRACT)
	return 4
}
func inc_d(b uint8) int{
	//fmt.Printf("INC d: %X\n", b)
	if (regs.d & 0x0F) == 0x0F{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.d++

	if regs.d == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.clearFlags(SUBTRACT)
	return 4
}
func inc_de(b uint8) int{
	//fmt.Printf("INC de: %X\n", b)
	regs.de_write(regs.de_read() + 1)
	return 8
}
func ld_de_a(b uint8) int{
	//fmt.Printf("LD (de), a: %X\n", b)
	writeByte(regs.de_read(), regs.a)
	return 8
}
func ld_de_aabb(b uint8) int{
	//fmt.Printf("LD de, $aabb: %X\n", b)
	regs.de_write(read16bits(regs.pc))
	regs.pc += 2
	return 12
}
func stop(b uint8) int{
	//fmt.Printf("TODO-STOP : %X\n", b)
	regs.pc++
	return 4
}
func rrca(b uint8) int{
	//fmt.Printf("RRCA: %X\n", b)
	if (regs.a & 0x01) == 0x01 {
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}
	regs.a = (regs.a >> 1) | (regs.a << 7)
	
	regs.clearFlags(SUBTRACT|HALFCARRY|ZERO)
	return 4
}
func ld_c_xx(b uint8) int{
	//fmt.Printf("LD c, $xx: %X\n", b)
	regs.c = readByte(regs.pc)
	regs.pc++
	return 8
}
func dec_c(b uint8) int{
	//fmt.Printf("DEC c: %X\n", b)
	if (regs.c & 0x0F) == 0{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.c--

	if regs.c == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.setFlags(SUBTRACT)
	return 4
}
func inc_c(b uint8) int{
	//fmt.Printf("INC c: %X\n", b)
	if (regs.c & 0x0F) == 0x0F{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.c++

	if regs.c == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.clearFlags(SUBTRACT)
	return 4
}
//-----------------------
func unknown(b uint8) int{
	//fmt.Printf("UNKNOWN: %X\n", b)
	return 0
}
//0x00
func nop(op uint8) int{
	//fmt.Printf("NOP: %X\n",op)
	return 4
}
//0x01
func ld_bc_aabb(b uint8) int{
	//fmt.Printf("LD bc, %X: %X\n",read16bits(regs.pc), b)
	regs.bc_write(read16bits(regs.pc))
	regs.pc += 2
	return 12
}
//0x02
func ld_bc_a(b uint8) int{
	//fmt.Printf("LD (bc), a: %X\n", b)
	writeByte(regs.bc_read(), regs.a)
	return 8
}
//0x03
func inc_bc(b uint8) int{
	//fmt.Printf("INC bc: %X\n", b)
	regs.bc_write(regs.bc_read() + 1)
	return 8
}
//0x04
func inc_b(b uint8) int{
	//fmt.Printf("INC b: %X\n", b)
	if (regs.b & 0x0F) == 0x0F{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.b++

	if regs.b == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.clearFlags(SUBTRACT)
	return 4
}

func dec_b(b uint8) int{
	//fmt.Printf("DEC b: %X\n", b)
	if (regs.b & 0x0F) == 0{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}

	regs.b--

	if regs.b == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.setFlags(SUBTRACT)
	return 4
}
func ld_b_xx(b uint8) int{
	//fmt.Printf("LD b, $xx: %X\n", b)
	regs.b = readByte(regs.pc)
	regs.pc++
	return 8
}
func rlca(b uint8) int{
	//fmt.Printf("RLCA: %X\n", b)
	if (regs.a & 0x80) == 0x80 {
			regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}
	regs.a = ((regs.a << 1) | (regs.a >> 7))

	regs.clearFlags(SUBTRACT|HALFCARRY|ZERO)
	return 4
}
func ld_aabb_sp(b uint8) int{
	var addr uint16 = read16bits(regs.pc)
	//fmt.Printf("LD (%X), sp: %X\n", addr, b)
	regs.pc += 2
	write16bits(addr, regs.sp)
	return 20
}
func add_hl_bc(b uint8) int{
	//fmt.Printf("ADD HL, BC: %X\n", b)
	regs.clearFlags(SUBTRACT)
	var result uint32 = uint32(regs.hl_read()) + uint32(regs.bc_read())

	if (result & 0xFF0000) != 0{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}
	if ((regs.hl_read() ^ regs.bc_read() ^ uint16(result&0xFFFF)) & 0x1000) != 0{
		regs.setFlags(HALFCARRY)
	} else{
		regs.clearFlags(HALFCARRY)
	}
	regs.hl_write(uint16(result))
	return 8
}
func ld_a_bc(b uint8) int{
	//fmt.Printf("LD a, (bc): %X\n", b)
	regs.a = readByte(regs.bc_read())
	return 8
}
func dec_bc(b uint8) int{
	//fmt.Printf("DEC bc: %X\n", b)
	regs.bc_write(regs.bc_read() - 1)
	return 8
}

