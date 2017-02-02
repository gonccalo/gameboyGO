package gameboygo

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
type operations func(uint8)

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
 	ld_a_bhl,	//0x7E
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
	rst_38}		//0xFF
func rst_38(b uint8) {
	
}
func cp_xx(b uint8) {
	
}
func ei(b uint8) {
	
}
func ld_a_aabb(b uint8) {
	
}
func ld_sp_hl(b uint8) {
	
}
func ld_hl_sp(b uint8) {
	
}
func rst_30(b uint8) {
	
}
func or_xx(b uint8) {
	
}
func push_af(b uint8) {
	
}
func di(b uint8) {
	
}
func ld_a_c(b uint8) {
	
}
func pop_af(b uint8) {
	
}
func ldh_a_xx(b uint8) {
	
}
func rst_28(b uint8) {
	
}
func xor_xx(b uint8) {
	
}
func ld_aabb_a(b uint8) {
	
}
func jp_hl(b uint8) {
	
}
func add_sp_xx(b uint8) {
	
}
func rst_20(b uint8) {
	
}
func and_xx(b uint8) {
	
}
func push_hl(b uint8) {
	
}
func ld_c_a(b uint8) {
	
}
func pop_hl(b uint8) {
	
}
func ld_xx_a(b uint8) {
	
}
func rst_18(b uint8) {
	
}
func sbc_a_xx(b uint8) {
	
}
func call_c_aabb(b uint8) {
	
}
func jp_c_aabb(b uint8) {
	
}
func reti(b uint8) {
	
}
func ret_c(b uint8) {
	
}
func rst_10(b uint8) {
	
}
func sub_xx(b uint8) {
	
}
func push_de(b uint8) {
	
}
func call_nc_aabb(b uint8) {
	
}
func jp_nc_aabb(b uint8) {
	
}
func pop_de(b uint8) {
	
}
func ret_nc(b uint8) {
	
}
func rst_08(b uint8) {
	
}
func adc_a_xx(b uint8) {
	
}
func call_aabb(b uint8) {
	
}
func call_z_aabb(b uint8) {
	
}
func many_ops(b uint8) {
	
}
func jp_z_aabb(b uint8) {
	
}
func ret(b uint8) {
	
}
func ret_z(b uint8) {
	
}
func rst_oo(b uint8) {
	
}
func add_a_xx(b uint8) {
	
}
func push_bc(b uint8) {
	
}
func call_nz_aabb(b uint8) {
	
}
func jp_aabb(b uint8) {
	
}
func jp_nz_aabb(b uint8) {
	
}
func pop_bc(b uint8) {
	
}
func ret_nz(b uint8) {
	
}
func cp_x(b uint8) {
	
}
func or_x(b uint8) {
	
}
func xor_x(b uint8) {
	
}
func and_x(b uint8) {
	
}
func sbc_a_x(b uint8) {
	
}
func sub_x(b uint8) {
	
}
func adc_a_x(b uint8) {
	
}
func add_a_x(b uint8) {
	
}
func halt(b uint8) {
	
}
func ld_a_bhl(b uint8) {
	fmt.Printf("LD a, (hl): %X\n", b)
	regs.a = readByte(regs.hl_read())
	//8 cicles
}
func ld_phl_a(b uint8) {
	fmt.Printf("LD (hl), a: %X\n", b)
	writeByte(regs.hl_read(), regs.a)
	//cicles 8
}
func ld_x_y(b uint8) {
	
}
func ld_sp_aabb(b uint8){
	fmt.Printf("LD sp, $aabb: %X\n", b)
	regs.sp = read16bits(regs.pc)
	regs.pc += 2
	//cicles 12
}
func ld_hld_a(b uint8){

}
func inc_sp(b uint8){
	fmt.Printf("INC sp: %X\n", b)
	regs.sp++
}
func inc_bhl(b uint8){
	fmt.Printf("INC hl: %X\n", b)
	var data uint8 = readByte(regs.hl_read())
	if (data & 0x0F) == 0x0F{
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}

	data++
	writeByte(regs.hl_read(), data)

	if data == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	clearFlags(SUBTRACT)
	//12 cicles
}
func dec_bhl(b uint8){
	fmt.Printf("DEC hl: %X\n", b)
	var data uint8 = readByte(regs.hl_read())
	if (data & 0x0F) == 0{
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}

	data--
	writeByte(regs.hl_read(), data)

	if data == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	setFlags(SUBTRACT)
	//12 cicles
}
func ld_hl_xx(b uint8){

}
func scf(b uint8){

}
func jr_c_xx(b uint8){

}
func add_hl_sp(b uint8){
	fmt.Printf("ADD HL, SP: %X\n", b)
	clearFlags(SUBTRACT)
	var result uint32 = uint32(regs.hl_read()) + uint32(regs.sp)

	if (result & 0xFF0000) != 0{
		setFlags(CARRY)
	} else{
		clearFlags(CARRY)
	}
	if ((regs.hl_read() & 0x0F)+(regs.sp & 0x0F)) > 0x0F {
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}
	regs.hl_write(uint16(result))
	//8 cicles
}
func ld_a_hld(b uint8){

}
func dec_sp(b uint8){
	fmt.Printf("DEC sp: %X\n", b)
	regs.sp--
	//8 cicles
}
func inc_a(b uint8){
	fmt.Printf("INC a: %X\n", b)
	if (regs.a & 0x0F) == 0x0F{
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}

	regs.a++

	if regs.a == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	clearFlags(SUBTRACT)
	//4 cicles
}
func dec_a(b uint8){
	fmt.Printf("DEC a: %X\n", b)
	if (regs.a & 0x0F) == 0{
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}

	regs.a--

	if regs.a == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	setFlags(SUBTRACT)
	//4 cicles
}
func ld_a_xx(b uint8){

}
func ccf(b uint8){

} 
func jr_nc_xx(b uint8) {
	
}
func cpl(b uint8) {
	
}
func ld_l_xx(b uint8) {
	fmt.Printf("LD l, xx: %X\n", b)
	regs.l = readByte(regs.pc)
	regs.pc++
	//8 cicles
}
func dec_l(b uint8) {
	fmt.Printf("DEC l: %X\n", b)
	if (regs.l & 0x0F) == 0{
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}

	regs.l--

	if regs.l == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	setFlags(SUBTRACT)
	//4 cicles
}
func inc_l(b uint8) {
	fmt.Printf("INC l: %X\n", b)
	if (regs.l & 0x0F) == 0x0F{
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}

	regs.l++

	if regs.l == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	clearFlags(SUBTRACT)
	//4 cicles
}
func dec_hl(b uint8) {
	fmt.Printf("DEC hl: %X\n", b)
	regs.hl_write(regs.hl_read() - 1)
	//8 cicles
}
func ld_a_hli(b uint8) {
	
}
func add_hl_hl(b uint8) {
	fmt.Printf("ADD HL, HL: %X\n", b)
	clearFlags(SUBTRACT)
	var result uint32 = uint32(regs.hl_read()) + uint32(regs.hl_read())

	if (result & 0xFF0000) != 0{
		setFlags(CARRY)
	} else{
		clearFlags(CARRY)
	}
	if ((regs.hl_read() & 0x0F)+(regs.hl_read() & 0x0F)) > 0x0F {
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}
	regs.hl_write(uint16(result))
}
func jr_z_xx(b uint8) {
	
}
func daa(b uint8) {
	
}
func ld_h_xx(b uint8) {
	fmt.Printf("LD h, xx: %X\n", b)
	regs.h = readByte(regs.pc)
	regs.pc++
	//8 cicles
}
func dec_h(b uint8) {
	fmt.Printf("DEC h: %X\n", b)
	if (regs.h & 0x0F) == 0{
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}

	regs.h--

	if regs.h == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	setFlags(SUBTRACT)
	//4 cicles
}
func inc_h(b uint8) {
	fmt.Printf("INC h: %X\n", b)
	if (regs.h & 0x0F) == 0x0F{
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}

	regs.h++

	if regs.h == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	clearFlags(SUBTRACT)
	//4 cicles
}
func inc_hl(b uint8) {
	fmt.Printf("INC hl: %X\n", b)
	regs.hl_write(regs.hl_read() + 1)
}
func ld_hli_a(b uint8) {
	
}
func ld_hl_aabb(b uint8) {
	fmt.Printf("LD hl, $aabb: %X\n", b)
	regs.hl_write(read16bits(regs.pc))
	regs.pc += 2
	//cicles 12
}
func jr_nz_xx(b uint8) {
	
}
func rra(b uint8) {
	
}
func ld_e_xx(b uint8) {
	fmt.Printf("LD e, xx: %X\n", b)
	regs.e = readByte(regs.pc)
	regs.pc++
	//8 cicles
}
func dec_e(b uint8) {
	fmt.Printf("DEC e: %X\n", b)
	if (regs.e & 0x0F) == 0{
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}

	regs.e--

	if regs.e == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	setFlags(SUBTRACT)
	//4 cicles
}
func inc_e(b uint8) {
	fmt.Printf("INC e: %X\n", b)
	if (regs.e & 0x0F) == 0x0F{
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}

	regs.e++

	if regs.e == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	clearFlags(SUBTRACT)
	//4 cicles
}
func dec_de(b uint8) {
	fmt.Printf("DEC de: %X\n", b)
	regs.de_write(regs.de_read() - 1)
	//8 cicles
}
func ld_a_de(b uint8) {
	fmt.Printf("LD a, (de): %X\n", b)
	regs.a = readByte(regs.de_read())
	//8 cicles
}
func add_hl_de(b uint8) {
	fmt.Printf("ADD HL, DE: %X\n", b)
	clearFlags(SUBTRACT)
	var result uint32 = uint32(regs.hl_read()) + uint32(regs.de_read())

	if (result & 0xFF0000) != 0{
		setFlags(CARRY)
	} else{
		clearFlags(CARRY)
	}
	if ((regs.hl_read() & 0x0F)+(regs.de_read() & 0x0F)) > 0x0F {
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}
	regs.hl_write(uint16(result))
	//8 cicles
}
func jr_xx(b uint8) {
	var dest int16 = int16(readByte(regs.pc))
	regs.pc++
	fmt.Printf("JR %X: %X\n",b,b)
	if dest>0 {
		regs.pc =regs.pc + uint16(dest)
	} else{
		regs.pc = regs.pc - uint16(-dest)
	}
}
func rla(b uint8) {
	
}
func ld_d_xx(b uint8) {
	fmt.Printf("LD d, xx: %X\n", b)
	regs.d = readByte(regs.pc)
	regs.pc++
	//8 cicles
}
func dec_d(b uint8) {
	fmt.Printf("DEC d: %X\n", b)
	if (regs.d & 0x0F) == 0{
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}

	regs.d--

	if regs.d == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	setFlags(SUBTRACT)
	//4 cicles
}
func inc_d(b uint8) {
	fmt.Printf("INC d: %X\n", b)
	if (regs.d & 0x0F) == 0x0F{
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}

	regs.d++

	if regs.d == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	clearFlags(SUBTRACT)
	//4 cicles
}
func inc_de(b uint8) {
	fmt.Printf("INC de: %X\n", b)
	regs.de_write(regs.de_read() + 1)
}
func ld_de_a(b uint8) {
	fmt.Printf("LD (de), a: %X\n", b)
	writeByte(regs.de_read(), regs.a)
	//cicles 8
}
func ld_de_aabb(b uint8) {
	fmt.Printf("LD de, $aabb: %X\n", b)
	regs.de_write(read16bits(regs.pc))
	regs.pc += 2
	//cicles 12
}
func stop(b uint8) {
	fmt.Printf("STOP : %X\n", b)

}
func rrca(b uint8) {
	fmt.Printf("RRCA: %X\n", b)
	if (regs.a & 0x01) == 0x01 {
		setFlags(CARRY)
	} else{
		clearFlags(CARRY)
	}
	regs.a = (regs.a >> 1) | (regs.a << 7)
	if regs.a == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	clearFlags(SUBTRACT|HALFCARRY)
}
func ld_c_xx(b uint8) {
	fmt.Printf("LD c, $xx: %X\n", b)
	regs.c = readByte(regs.pc)
	regs.pc++
	//8 cicles
}
func dec_c(b uint8) {
	fmt.Printf("DEC c: %X\n", b)
	if (regs.c & 0x0F) == 0{
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}

	regs.c--

	if regs.c == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	setFlags(SUBTRACT)
	//4 cicles
}
func inc_c(b uint8) {
	fmt.Printf("INC c: %X\n", b)
	if (regs.c & 0x0F) == 0x0F{
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}

	regs.c++

	if regs.c == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	clearFlags(SUBTRACT)
	//4 cicles
}
//-----------------------
func unknown(b uint8) {
	fmt.Printf("UNKNOWN: %X\n", b)
}

func nop(op uint8) {
	fmt.Printf("NOP: %X\n",op)
	//cicles: 4
}

func ld_bc_aabb(b uint8) {
	fmt.Printf("LD bc, $aabb: %X\n", b)
	regs.bc_write(read16bits(regs.pc))
	regs.pc += 2
	//cicles 12
}

func ld_bc_a(b uint8) {
	fmt.Printf("LD (bc), a: %X\n", b)
	writeByte(regs.bc_read(), regs.a)
	//cicles 8
}

func inc_bc(b uint8) {
	fmt.Printf("INC bc: %X\n", b)
	regs.bc_write(regs.bc_read() + 1)
	//8 cicles
}

func inc_b(b uint8) {
	fmt.Printf("INC b: %X\n", b)
	if (regs.b & 0x0F) == 0x0F{
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}

	regs.b++

	if regs.b == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	clearFlags(SUBTRACT)
	//4 cicles
}

func dec_b(b uint8) {
	fmt.Printf("DEC b: %X\n", b)
	if (regs.b & 0x0F) == 0{
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}

	regs.b--

	if regs.b == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	setFlags(SUBTRACT)
	//4 cicles
}
func ld_b_xx(b uint8) {
	fmt.Printf("LD b, $xx: %X\n", b)
	regs.b = readByte(regs.pc)
	regs.pc++
	//8 cicles
}
func rlca(b uint8) {
	fmt.Printf("RLCA: %X\n", b)
	if (regs.a & 0x80) == 0x80 {
			setFlags(CARRY)
	} else{
		clearFlags(CARRY)
	}
	regs.a = ((regs.a << 1) | (regs.a >> 7))
	if regs.a == 0{
		setFlags(ZERO)
	} else{
		clearFlags(ZERO)
	}
	clearFlags(SUBTRACT|HALFCARRY)
	//4 cicles
}
func ld_aabb_sp(b uint8) {
	var addr uint16 = read16bits(regs.pc)
	fmt.Printf("LD (%X), sp: %X\n", addr, b)
	regs.pc += 2
	write16bits(addr, regs.sp)
	//20 cicles
}
func add_hl_bc(b uint8) {
	fmt.Printf("ADD HL, BC: %X\n", b)
	clearFlags(SUBTRACT)
	var result uint32 = uint32(regs.hl_read()) + uint32(regs.bc_read())

	if (result & 0xFF0000) != 0{
		setFlags(CARRY)
	} else{
		clearFlags(CARRY)
	}
	if ((regs.hl_read() & 0x0F)+(regs.bc_read() & 0x0F)) > 0x0F {
		setFlags(HALFCARRY)
	} else{
		clearFlags(HALFCARRY)
	}
	regs.hl_write(uint16(result))
	//8 cicles
}
func ld_a_bc(b uint8) {
	fmt.Printf("LD a, (bc): %X\n", b)
	regs.a = readByte(regs.bc_read())
	//8 cicles
}
func dec_bc(b uint8) {
	fmt.Printf("DEC bc: %X\n", b)
	regs.bc_write(regs.bc_read() - 1)
	//8 cicles
}
/*
registos 
AF   --> accumulator e flags
BC   --> pode ser acedido como B ou C 
DE   -->
HL   -->
SP   --> stack pointer
PC   --> program counter
*/
type registers struct{
	a  uint8;
	f  uint8;
	
	b  uint8;
	c  uint8;
	
	d  uint8;
	e  uint8;
	
	h  uint8;
	l  uint8;
	
	sp uint16;
	pc uint16;
}
func (r *registers)bc_write(data uint16) {
	r.b = uint8((data & 0xFF00) >> 8)
	r.c = uint8(data & 0x00FF)
}
func (r *registers)bc_read() uint16{
	return uint16((uint16(r.b) << 8) | uint16(r.c)) 
}
func (r *registers)hl_write(data uint16) {
	r.h = uint8((data & 0xFF00) >> 8)
	r.l = uint8(data & 0x00FF)	
}
func (r *registers)hl_read() uint16 {
	return uint16((uint16(r.h) << 8) | uint16(r.l))
}
func (r *registers)de_write(data uint16){
	r.d = uint8((data & 0xFF00) >> 8)
	r.e = uint8(data & 0x00FF)
}
func (r *registers)de_read() uint16 {
	return uint16((uint16(r.d) << 8) | uint16(r.e))
}
var regs registers

func Reset() {
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

func Fetch_decode() {
	var op = readByte(regs.pc)
	regs.pc++
	ops[op](op)

}

