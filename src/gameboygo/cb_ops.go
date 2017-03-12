package gameboygo

var cb_op = [0x100]operations{
	rlc_b,		//0x00
	rlc_c,		//0x01
	rlc_d,		//0x02
	rlc_e,		//0x03
	rlc_h,		//0x04
	rlc_l,		//0x05
	rlc_hl,		//0x06
	rlc_a,		//0x07
	rrc_b,		//0x08
	rrc_c,		//0x09
	rrc_d,		//0x0A
	rrc_e,		//0x0B
	rrc_h,		//0x0C
	rrc_l,		//0x0D
	rrc_hl,		//0x0E
	rrc_a,		//0x0F
	rl_b,		//0x10
	rl_c,		//0x11
	rl_d,		//0x12
	rl_e,		//0x13
	rl_h,		//0x14
	rl_l,		//0x15
	rl_hl,		//0x16
	rl_a,		//0x17
	rr_b,		//0x18
	rr_c,		//0x19
	rr_d,		//0x1A
	rr_e,		//0x1B
	rr_h,		//0x1C
	rr_l,		//0x1D
	rr_hl,		//0x1E
	rr_a,		//0x1F
	sla_b,		//0x20
	sla_c,		//0x21
	sla_d,		//0x22
	sla_e,		//0x23
	sla_h,		//0x24
	sla_l,		//0x25
	sla_hl,		//0x26
	sla_a,		//0x27
	sra_b,		//0x28
	sra_c,		//0x29
	sra_d,		//0x2A
	sra_e,		//0x2B
	sra_h,		//0x2C
	sra_l,		//0x2D
	sra_hl,		//0x2E
	sra_a,		//0x2F

}
func sra_a(op uint8) {
	//fmt.Printf("SRA A: %X", op)
	sra(&regs.a)
	CicleCounter += 8
}
func sra_hl(op uint8) {
	//fmt.Printf("SRA HL: %X", op)
	var r uint8 = readByte(regs.hl_read())
	sra(&r)
	writeByte(regs.hl_read(), r)
	CicleCounter += 16
}
func sra_l(op uint8) {
	//fmt.Printf("SRA L: %X", op)
	sra(&regs.l)
	CicleCounter += 8
}
func sra_h(op uint8) {
	//fmt.Printf("SRA H: %X", op)
	sra(&regs.h)
	CicleCounter += 8
}
func sra_e(op uint8) {
	//fmt.Printf("SRA E: %X", op)
	sra(&regs.e)
	CicleCounter += 8
}
func sra_d(op uint8) {
	//fmt.Printf("SRA D: %X", op)
	sra(&regs.d)
	CicleCounter += 8
}
func sra_c(op uint8) {
	//fmt.Printf("SRA C: %X", op)
	sra(&regs.c)
	CicleCounter += 8
}
func sra_b(op uint8) {
	//fmt.Printf("SRA B: %X", op)
	sra(&regs.b)
	CicleCounter += 8
}

func sla_a(op uint8) {
	//fmt.Printf("SLA A: %X", op)
	sla(&regs.a)
	CicleCounter += 8
}
func sla_hl(op uint8) {
	//fmt.Printf("SLA HL: %X", op)
	var r uint8 = readByte(regs.hl_read())
	sla(&r)
	writeByte(regs.hl_read(), r)
	CicleCounter += 16
}
func sla_l(op uint8) {
	//fmt.Printf("SLA L: %X", op)
	sla(&regs.l)
	CicleCounter += 8
}
func sla_h(op uint8) {
	//fmt.Printf("SLA H: %X", op)
	sla(&regs.h)
	CicleCounter += 8
}
func sla_e(op uint8) {
	//fmt.Printf("SLA E: %X", op)
	sla(&regs.e)
	CicleCounter += 8
}
func sla_d(op uint8) {
	//fmt.Printf("SLA D: %X", op)
	sla(&regs.d)
	CicleCounter += 8
}
func sla_c(op uint8) {
	//fmt.Printf("SLA C: %X", op)
	sla(&regs.c)
	CicleCounter += 8
}
func sla_b(op uint8) {
	//fmt.Printf("SLA B: %X", op)
	sla(&regs.b)
	CicleCounter += 8
}

func rr_a(op uint8) {
	//fmt.Printf("RR A: %X", op)
	rr(&regs.a)
	CicleCounter += 8
}
func rr_hl(op uint8) {
	//fmt.Printf("RR HL: %X", op)
	var r uint8 = readByte(regs.hl_read())
	rr(&r)
	writeByte(regs.hl_read(), r)
	CicleCounter += 16
}
func rr_l(op uint8) {
	//fmt.Printf("RR L: %X", op)
	rr(&regs.l)
	CicleCounter += 8
}
func rr_h(op uint8) {
	//fmt.Printf("RR H: %X", op)
	rr(&regs.h)
	CicleCounter += 8
}
func rr_e(op uint8) {
	//fmt.Printf("RR E: %X", op)
	rr(&regs.e)
	CicleCounter += 8
}
func rr_d(op uint8) {
	//fmt.Printf("RR D: %X", op)
	rr(&regs.d)
	CicleCounter += 8
}
func rr_c(op uint8) {
	//fmt.Printf("RR C: %X", op)
	rr(&regs.c)
	CicleCounter += 8
}
func rr_b(op uint8) {
	//fmt.Printf("RR B: %X", op)
	rr(&regs.b)
	CicleCounter += 8
}

func rl_a(op uint8) {
	//fmt.Printf("RL A: %X", op)
	rl(&regs.a)
	CicleCounter += 8
}
func rl_hl(op uint8) {
	//fmt.Printf("RL (HL): %X", op)
	var r uint8 = readByte(regs.hl_read())
	rl(&r)
	writeByte(regs.hl_read(), r)
	CicleCounter += 16
}
func rl_l(op uint8) {
	//fmt.Printf("RL L: %X", op)
	rl(&regs.l)
	CicleCounter += 8
}
func rl_h(op uint8) {
	//fmt.Printf("RL H: %X", op)
	rl(&regs.h)
	CicleCounter += 8
}
func rl_e(op uint8) {
	//fmt.Printf("RL E: %X", op)
	rl(&regs.e)
	CicleCounter += 8
}
func rl_d(op uint8) {
	//fmt.Printf("RL D: %X", op)
	rl(&regs.d)
	CicleCounter += 8
}
func rl_c(op uint8) {
	//fmt.Printf("RL C: %X", op)
	rl(&regs.c)
	CicleCounter += 8
}
func rl_b(op uint8) {
	//fmt.Printf("RL B: %X", op)
	rl(&regs.b)
	CicleCounter += 8
}

func rrc_a(op uint8) {
	//fmt.Printf("RRC A: %X", op)
	rrc(&regs.a)
	CicleCounter += 8
}
func rrc_hl(op uint8) {
	//fmt.Printf("RRC (hl): %X", op)
	var r uint8 = readByte(regs.hl_read())
	rrc(&r)
	writeByte(regs.hl_read(), r)
	CicleCounter += 16
}
func rrc_l(op uint8) {
	//fmt.Printf("RRC L: %X", op)
	rrc(&regs.l)
	CicleCounter += 8
}
func rrc_h(op uint8) {
	//fmt.Printf("RRC H: %X", op)
	rrc(&regs.h)
	CicleCounter += 8
}
func rrc_e(op uint8) {
	//fmt.Printf("RRC E: %X", op)
	rrc(&regs.e)
	CicleCounter += 8
}
func rrc_d(op uint8) {
	//fmt.Printf("RRC D: %X", op)
	rrc(&regs.d)
	CicleCounter += 8
}
func rrc_c(op uint8) {
	//fmt.Printf("RRC C: %X", op)
	rrc(&regs.c)
	CicleCounter += 8
}
func rrc_b(op uint8) {
	//fmt.Printf("RRC B: %X", op)
	rrc(&regs.b)
	CicleCounter += 8
}
func rlc_a(op uint8) {
	//fmt.Printf("RLC A: %X", op)
	rlc(&regs.a)
	CicleCounter += 8
}
func rlc_hl(op uint8) {
	//fmt.Printf("RLC (hl): %X", op)
	var r uint8 = readByte(regs.hl_read())
	rlc(&r)
	writeByte(regs.hl_read(), r)
	CicleCounter += 16
}
func rlc_l(op uint8) {
	//fmt.Printf("RLC L: %X", op)
	rlc(&regs.l)
	CicleCounter += 8
}
func rlc_h(op uint8) {
	//fmt.Printf("RLC H: %X", op)
	rlc(&regs.h)
	CicleCounter += 8
}
func rlc_e(op uint8) {
	//fmt.Printf("RLC E: %X", op)
	rlc(&regs.e)
	CicleCounter += 8
}
func rlc_d(op uint8) {
	//fmt.Printf("RLC D: %X", op)
	rlc(&regs.d)	
	CicleCounter += 8
}
func rlc_c(op uint8) {
	//fmt.Printf("RLC C: %X", op)
	rlc(&regs.c)
	CicleCounter += 8
}
func rlc_b(op uint8) {
	//fmt.Printf("RLC B: %X", op)
	rlc(&regs.b)
	CicleCounter += 8
}



func rlc(r *uint8) {
	regs.clearFlags(SUBTRACT|HALFCARRY)
	if(*r & 0x80) != 0{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}

	*r = (*r << 1) | (*r >> 7)

	if *r == 0{
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
}

func rrc(r *uint8) {
	regs.clearFlags(SUBTRACT|HALFCARRY)
	if(*r & 0x01) != 0{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}

	*r = (*r >> 1) | (*r << 7)

	if *r == 0 {
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
}

func rl(r *uint8) {
	regs.clearFlags(SUBTRACT|HALFCARRY)
	var carry uint8 = 0
	if regs.getFlag(CARRY){
		carry = 1
	}
	if(*r & 0x80) != 0{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}

	*r = *r << 1
	*r += carry

	if *r == 0 {
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
}

func rr(r *uint8) {
	regs.clearFlags(SUBTRACT|HALFCARRY)
	var carry uint8 = 0
	if regs.getFlag(CARRY){
		carry = 0x80
	}
	if(*r & 0x01) != 0{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}

	*r = *r >> 1
	*r |= carry 

	if *r == 0 {
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
}

func sla(r *uint8) {
	regs.clearFlags(SUBTRACT|HALFCARRY)
	if(*r & 0x80) != 0{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}

	*r = *r << 1
	
	if *r == 0 {
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
}

func sra(r *uint8) {
	//"msb doesn't change"?
	regs.clearFlags(SUBTRACT|HALFCARRY)
	if(*r & 0x01) != 0{
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}

	*r = (*r >> 1) | (*r & 0x80)
	
	if *r == 0 {
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
}