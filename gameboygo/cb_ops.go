package main

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
	swap_b,		//0x30
	swap_c,		//0x31
	swap_d,		//0x32
	swap_e,		//0x33
	swap_h,		//0x34
	swap_l,		//0x35
	swap_hl,	//0x36
	swap_a,		//0x37
	srl_b,		//0x38
	srl_c,		//0x39
	srl_d,		//0x3A
	srl_e,		//0x3B
	srl_h,		//0x3C
	srl_l,		//0x3D
	srl_hl,		//0x3E
	srl_a,		//0x3F
	bit_0_r,	//0x40
	bit_0_r,	//0x41
	bit_0_r,	//0x42
	bit_0_r,	//0x43
	bit_0_r,	//0x44
	bit_0_r,	//0x45
	bit_0_r,	//0x46
	bit_0_r,	//0x47
	bit_1_r,	//0x48
	bit_1_r,	//0x49
	bit_1_r,	//0x4A
	bit_1_r,	//0x4B
	bit_1_r,	//0x4C
	bit_1_r,	//0x4D
	bit_1_r,	//0x4E
	bit_1_r,	//0x4F
	bit_2_r,	//0x50
	bit_2_r,	//0x51
	bit_2_r,	//0x52
	bit_2_r,	//0x53
	bit_2_r,	//0x54
	bit_2_r,	//0x55
	bit_2_r,	//0x56
	bit_2_r,	//0x57
	bit_3_r,	//0x58
	bit_3_r,	//0x59
	bit_3_r,	//0x5A
	bit_3_r,	//0x5B
	bit_3_r,	//0x5C
	bit_3_r,	//0x5D
	bit_3_r,	//0x5E
	bit_3_r,	//0x5F
	bit_4_r,	//0x60
	bit_4_r,	//0x61
	bit_4_r,	//0x62
	bit_4_r,	//0x63
	bit_4_r,	//0x64
	bit_4_r,	//0x65
	bit_4_r,	//0x66
	bit_4_r,	//0x67
	bit_5_r,	//0x68
	bit_5_r,	//0x69
	bit_5_r,	//0x6A
	bit_5_r,	//0x6B
	bit_5_r,	//0x6C
	bit_5_r,	//0x6D
	bit_5_r,	//0x6E
	bit_5_r,	//0x6F
	bit_6_r,	//0x70
	bit_6_r,	//0x71
	bit_6_r,	//0x72
	bit_6_r,	//0x73
	bit_6_r,	//0x74
	bit_6_r,	//0x75
	bit_6_r,	//0x76
	bit_6_r,	//0x77
	bit_7_r,	//0x78
	bit_7_r,	//0x79
	bit_7_r,	//0x7A
	bit_7_r,	//0x7B
	bit_7_r,	//0x7C
	bit_7_r,	//0x7D
	bit_7_r,	//0x7E
	bit_7_r,	//0x7F
	res_0_r,	//0x80
	res_0_r,	//0x81
	res_0_r,	//0x82
	res_0_r,	//0x83
	res_0_r,	//0x84
	res_0_r,	//0x85
	res_0_r,	//0x86
	res_0_r,	//0x87
	res_1_r,	//0x88
	res_1_r,	//0x89
	res_1_r,	//0x8A
	res_1_r,	//0x8B
	res_1_r,	//0x8C
	res_1_r,	//0x8D
	res_1_r,	//0x8E
	res_1_r,	//0x8F
	res_2_r,	//0x90
	res_2_r,	//0x91
	res_2_r,	//0x92
	res_2_r,	//0x93
	res_2_r,	//0x94
	res_2_r,	//0x95
	res_2_r,	//0x96
	res_2_r,	//0x97
	res_3_r,	//0x98
	res_3_r,	//0x99
	res_3_r,	//0x9A
	res_3_r,	//0x9B
	res_3_r,	//0x9C
	res_3_r,	//0x9D
	res_3_r,	//0x9E
	res_3_r,	//0x9F
	res_4_r,	//0xA0
	res_4_r,	//0xA1
	res_4_r,	//0xA2
	res_4_r,	//0xA3
	res_4_r,	//0xA4
	res_4_r,	//0xA5
	res_4_r,	//0xA6
	res_4_r,	//0xA7
	res_5_r,	//0xA8
	res_5_r,	//0xA9
	res_5_r,	//0xAA
	res_5_r,	//0xAB
	res_5_r,	//0xAC
	res_5_r,	//0xAD
	res_5_r,	//0xAE
	res_5_r,	//0xAF
	res_6_r,	//0xB0
	res_6_r,	//0xB1
	res_6_r,	//0xB2
	res_6_r,	//0xB3
	res_6_r,	//0xB4
	res_6_r,	//0xB5
	res_6_r,	//0xB6
	res_6_r,	//0xB7
	res_7_r,	//0xB8
	res_7_r,	//0xB9
	res_7_r,	//0xBA
	res_7_r,	//0xBB
	res_7_r,	//0xBC
	res_7_r,	//0xBD
	res_7_r,	//0xBE
	res_7_r,	//0xBF
	set_0_r,	//0xC0
	set_0_r,	//0xC1
	set_0_r,	//0xC2
	set_0_r,	//0xC3
	set_0_r,	//0xC4
	set_0_r,	//0xC5
	set_0_r,	//0xC6
	set_0_r,	//0xC7
	set_1_r,	//0xC8
	set_1_r,	//0xC9
	set_1_r,	//0xCA
	set_1_r,	//0xCB
	set_1_r,	//0xCC
	set_1_r,	//0xCD
	set_1_r,	//0xCE
	set_1_r,	//0xCF
	set_2_r,	//0xD0
	set_2_r,	//0xD1
	set_2_r,	//0xD2
	set_2_r,	//0xD3
	set_2_r,	//0xD4
	set_2_r,	//0xD5
	set_2_r,	//0xD6
	set_2_r,	//0xD7
	set_3_r,	//0xD8
	set_3_r,	//0xD9
	set_3_r,	//0xDA
	set_3_r,	//0xDB
	set_3_r,	//0xDC
	set_3_r,	//0xDD
	set_3_r,	//0xDE
	set_3_r,	//0xDF
	set_4_r,	//0xE0
	set_4_r,	//0xE1
	set_4_r,	//0xE2
	set_4_r,	//0xE3
	set_4_r,	//0xE4
	set_4_r,	//0xE5
	set_4_r,	//0xE6
	set_4_r,	//0xE7
	set_5_r,	//0xE8
	set_5_r,	//0xE9
	set_5_r,	//0xEA
	set_5_r,	//0xEB
	set_5_r,	//0xEC
	set_5_r,	//0xED
	set_5_r,	//0xEE
	set_5_r,	//0xEF
	set_6_r,	//0xF0
	set_6_r,	//0xF1
	set_6_r,	//0xF2
	set_6_r,	//0xF3
	set_6_r,	//0xF4
	set_6_r,	//0xF5
	set_6_r,	//0xF6
	set_6_r,	//0xF7
	set_7_r,	//0xF8
	set_7_r,	//0xF9
	set_7_r,	//0xFA
	set_7_r,	//0xFB
	set_7_r,	//0xFC
	set_7_r,	//0xFD
	set_7_r,	//0xFE
	set_7_r,	//0xFF
}

func set_7_r(op uint8) {
	var to_set uint8 = 0x80
	switch (op & 0x0F){
	case 0x08:
		//fmt.Printf("SET 1, B: %X", op)
		set(to_set, &regs.b)
	case 0x09:
		//fmt.Printf("SET 1, C: %X", op)
		set(to_set, &regs.c)
	case 0x0A:
		//fmt.Printf("SET 1, D: %X", op)
		set(to_set, &regs.d)
	case 0x0B:
		//fmt.Printf("SET 1, E: %X", op)
		set(to_set, &regs.e)
	case 0x0C:
		//fmt.Printf("SET 1, H: %X", op)
		set(to_set, &regs.h)
	case 0x0D:
		//fmt.Printf("SET 1, L: %X", op)
		set(to_set, &regs.l)
	case 0x0E:
		//fmt.Printf("SET 1, (HL): %X", op)
		var r uint8 = readByte(regs.hl_read())
		set(to_set, &r)
		writeByte(regs.hl_read(), r)
		CicleCounter += 8
	case 0x0F:
		//fmt.Printf("SET 1, A: %X", op)
		set(to_set, &regs.a)
	}
	CicleCounter += 8
}
func set_6_r(op uint8) {
	var to_set uint8 = 0x40
	switch (op & 0x0F){
	case 0x00:
		//fmt.Printf("SET 0, B: %X", op)
		set(to_set, &regs.b)
	case 0x01:
		//fmt.Printf("SET 0, C: %X", op)
		set(to_set, &regs.c)
	case 0x02:
		//fmt.Printf("SET 0, D: %X", op)
		set(to_set, &regs.d)
	case 0x03:
		//fmt.Printf("SET 0, E: %X", op)
		set(to_set, &regs.e)
	case 0x04:
		//fmt.Printf("SET 0, H: %X", op)
		set(to_set, &regs.h)
	case 0x05:
		//fmt.Printf("SET 0, L: %X", op)
		set(to_set, &regs.l)
	case 0x06:
		//fmt.Printf("SET 0, (HL): %X", op)
		var r uint8 = readByte(regs.hl_read())
		set(to_set, &r)
		writeByte(regs.hl_read(), r)
		CicleCounter += 8
	case 0x07:
		//fmt.Printf("SET 0, A: %X", op)
		set(to_set, &regs.a)
	}
	CicleCounter += 8
}
func set_5_r(op uint8) {
	var to_set uint8 = 0x20
	switch (op & 0x0F){
	case 0x08:
		//fmt.Printf("SET 1, B: %X", op)
		set(to_set, &regs.b)
	case 0x09:
		//fmt.Printf("SET 1, C: %X", op)
		set(to_set, &regs.c)
	case 0x0A:
		//fmt.Printf("SET 1, D: %X", op)
		set(to_set, &regs.d)
	case 0x0B:
		//fmt.Printf("SET 1, E: %X", op)
		set(to_set, &regs.e)
	case 0x0C:
		//fmt.Printf("SET 1, H: %X", op)
		set(to_set, &regs.h)
	case 0x0D:
		//fmt.Printf("SET 1, L: %X", op)
		set(to_set, &regs.l)
	case 0x0E:
		//fmt.Printf("SET 1, (HL): %X", op)
		var r uint8 = readByte(regs.hl_read())
		set(to_set, &r)
		writeByte(regs.hl_read(), r)
		CicleCounter += 8
	case 0x0F:
		//fmt.Printf("SET 1, A: %X", op)
		set(to_set, &regs.a)
	}
	CicleCounter += 8
}
func set_4_r(op uint8) {
	var to_set uint8 = 0x10
	switch (op & 0x0F){
	case 0x00:
		//fmt.Printf("SET 0, B: %X", op)
		set(to_set, &regs.b)
	case 0x01:
		//fmt.Printf("SET 0, C: %X", op)
		set(to_set, &regs.c)
	case 0x02:
		//fmt.Printf("SET 0, D: %X", op)
		set(to_set, &regs.d)
	case 0x03:
		//fmt.Printf("SET 0, E: %X", op)
		set(to_set, &regs.e)
	case 0x04:
		//fmt.Printf("SET 0, H: %X", op)
		set(to_set, &regs.h)
	case 0x05:
		//fmt.Printf("SET 0, L: %X", op)
		set(to_set, &regs.l)
	case 0x06:
		//fmt.Printf("SET 0, (HL): %X", op)
		var r uint8 = readByte(regs.hl_read())
		set(to_set, &r)
		writeByte(regs.hl_read(), r)
		CicleCounter += 8
	case 0x07:
		//fmt.Printf("SET 0, A: %X", op)
		set(to_set, &regs.a)
	}
	CicleCounter += 8
}
func set_3_r(op uint8) {
	var to_set uint8 = 0x08
	switch (op & 0x0F){
	case 0x08:
		//fmt.Printf("SET 1, B: %X", op)
		set(to_set, &regs.b)
	case 0x09:
		//fmt.Printf("SET 1, C: %X", op)
		set(to_set, &regs.c)
	case 0x0A:
		//fmt.Printf("SET 1, D: %X", op)
		set(to_set, &regs.d)
	case 0x0B:
		//fmt.Printf("SET 1, E: %X", op)
		set(to_set, &regs.e)
	case 0x0C:
		//fmt.Printf("SET 1, H: %X", op)
		set(to_set, &regs.h)
	case 0x0D:
		//fmt.Printf("SET 1, L: %X", op)
		set(to_set, &regs.l)
	case 0x0E:
		//fmt.Printf("SET 1, (HL): %X", op)
		var r uint8 = readByte(regs.hl_read())
		set(to_set, &r)
		writeByte(regs.hl_read(), r)
		CicleCounter += 8
	case 0x0F:
		//fmt.Printf("SET 1, A: %X", op)
		set(to_set, &regs.a)
	}
	CicleCounter += 8
}
func set_2_r(op uint8) {
	var to_set uint8 = 0x04
	switch (op & 0x0F){
	case 0x00:
		//fmt.Printf("SET 0, B: %X", op)
		set(to_set, &regs.b)
	case 0x01:
		//fmt.Printf("SET 0, C: %X", op)
		set(to_set, &regs.c)
	case 0x02:
		//fmt.Printf("SET 0, D: %X", op)
		set(to_set, &regs.d)
	case 0x03:
		//fmt.Printf("SET 0, E: %X", op)
		set(to_set, &regs.e)
	case 0x04:
		//fmt.Printf("SET 0, H: %X", op)
		set(to_set, &regs.h)
	case 0x05:
		//fmt.Printf("SET 0, L: %X", op)
		set(to_set, &regs.l)
	case 0x06:
		//fmt.Printf("SET 0, (HL): %X", op)
		var r uint8 = readByte(regs.hl_read())
		set(to_set, &r)
		writeByte(regs.hl_read(), r)
		CicleCounter += 8
	case 0x07:
		//fmt.Printf("SET 0, A: %X", op)
		set(to_set, &regs.a)
	}
	CicleCounter += 8
}
func set_1_r(op uint8) {
	var to_set uint8 = 0x02
	switch (op & 0x0F){
	case 0x08:
		//fmt.Printf("SET 1, B: %X", op)
		set(to_set, &regs.b)
	case 0x09:
		//fmt.Printf("SET 1, C: %X", op)
		set(to_set, &regs.c)
	case 0x0A:
		//fmt.Printf("SET 1, D: %X", op)
		set(to_set, &regs.d)
	case 0x0B:
		//fmt.Printf("SET 1, E: %X", op)
		set(to_set, &regs.e)
	case 0x0C:
		//fmt.Printf("SET 1, H: %X", op)
		set(to_set, &regs.h)
	case 0x0D:
		//fmt.Printf("SET 1, L: %X", op)
		set(to_set, &regs.l)
	case 0x0E:
		//fmt.Printf("SET 1, (HL): %X", op)
		var r uint8 = readByte(regs.hl_read())
		set(to_set, &r)
		writeByte(regs.hl_read(), r)
		CicleCounter += 8
	case 0x0F:
		//fmt.Printf("SET 1, A: %X", op)
		set(to_set, &regs.a)
	}
	CicleCounter += 8
}
func set_0_r(op uint8) {
	var to_set uint8 = 0x01
	switch (op & 0x0F){
	case 0x00:
		//fmt.Printf("SET 0, B: %X", op)
		set(to_set, &regs.b)
	case 0x01:
		//fmt.Printf("SET 0, C: %X", op)
		set(to_set, &regs.c)
	case 0x02:
		//fmt.Printf("SET 0, D: %X", op)
		set(to_set, &regs.d)
	case 0x03:
		//fmt.Printf("SET 0, E: %X", op)
		set(to_set, &regs.e)
	case 0x04:
		//fmt.Printf("SET 0, H: %X", op)
		set(to_set, &regs.h)
	case 0x05:
		//fmt.Printf("SET 0, L: %X", op)
		set(to_set, &regs.l)
	case 0x06:
		//fmt.Printf("SET 0, (HL): %X", op)
		var r uint8 = readByte(regs.hl_read())
		set(to_set, &r)
		writeByte(regs.hl_read(), r)
		CicleCounter += 8
	case 0x07:
		//fmt.Printf("SET 0, A: %X", op)
		set(to_set, &regs.a)
	}
	CicleCounter += 8
}

func res_7_r(op uint8) {
	var to_res uint8 = 0x80
	switch (op & 0x0F){
	case 0x08:
		//fmt.Printf("RES 7, B: %X", op)
		res(to_res, &regs.b)
	case 0x09:
		//fmt.Printf("RES 7, C: %X", op)
		res(to_res, &regs.c)
	case 0x0A:
		//fmt.Printf("RES 7, D: %X", op)
		res(to_res, &regs.d)
	case 0x0B:
		//fmt.Printf("RES 7, E: %X", op)
		res(to_res, &regs.e)
	case 0x0C:
		//fmt.Printf("RES 7, H: %X", op)
		res(to_res, &regs.h)
	case 0x0D:
		//fmt.Printf("RES 7, L: %X", op)
		res(to_res, &regs.l)
	case 0x0E:
		//fmt.Printf("RES 7, (HL): %X", op)
		var r uint8 = readByte(regs.hl_read())
		res(to_res, &r)
		writeByte(regs.hl_read(), r)
		CicleCounter += 8
	case 0x0F:
		//fmt.Printf("RES 7, A: %X", op)
		res(to_res, &regs.a)
	}
	CicleCounter += 8
}
func res_6_r(op uint8) {
	var to_res uint8 = 0x40
	switch (op & 0x0F){
	case 0x00:
		//fmt.Printf("RES 6, B: %X", op)
		res(to_res, &regs.b)
	case 0x01:
		//fmt.Printf("RES 6, C: %X", op)
		res(to_res, &regs.c)
	case 0x02:
		//fmt.Printf("RES 6, D: %X", op)
		res(to_res, &regs.d)
	case 0x03:
		//fmt.Printf("RES 6, E: %X", op)
		res(to_res, &regs.e)
	case 0x04:
		//fmt.Printf("RES 6, H: %X", op)
		res(to_res, &regs.h)
	case 0x05:
		//fmt.Printf("RES 6, L: %X", op)
		res(to_res, &regs.l)
	case 0x06:
		//fmt.Printf("RES 6, (HL): %X", op)
		var r uint8 = readByte(regs.hl_read())
		res(to_res, &r)
		writeByte(regs.hl_read(), r)
		CicleCounter += 8
	case 0x07:
		//fmt.Printf("RES 6, A: %X", op)
		res(to_res, &regs.a)
	}
	CicleCounter += 8
}
func res_5_r(op uint8) {
	var to_res uint8 = 0x20
	switch (op & 0x0F){
	case 0x08:
		//fmt.Printf("RES 5, B: %X", op)
		res(to_res, &regs.b)
	case 0x09:
		//fmt.Printf("RES 5, C: %X", op)
		res(to_res, &regs.c)
	case 0x0A:
		//fmt.Printf("RES 5, D: %X", op)
		res(to_res, &regs.d)
	case 0x0B:
		//fmt.Printf("RES 5, E: %X", op)
		res(to_res, &regs.e)
	case 0x0C:
		//fmt.Printf("RES 5, H: %X", op)
		res(to_res, &regs.h)
	case 0x0D:
		//fmt.Printf("RES 5, L: %X", op)
		res(to_res, &regs.l)
	case 0x0E:
		//fmt.Printf("RES 5, (HL): %X", op)
		var r uint8 = readByte(regs.hl_read())
		res(to_res, &r)
		writeByte(regs.hl_read(), r)
		CicleCounter += 8
	case 0x0F:
		//fmt.Printf("RES 5, A: %X", op)
		res(to_res, &regs.a)
	}
	CicleCounter += 8
}
func res_4_r(op uint8) {
	var to_res uint8 = 0x10
	switch (op & 0x0F){
	case 0x00:
		//fmt.Printf("RES 4, B: %X", op)
		res(to_res, &regs.b)
	case 0x01:
		//fmt.Printf("RES 4, C: %X", op)
		res(to_res, &regs.c)
	case 0x02:
		//fmt.Printf("RES 4, D: %X", op)
		res(to_res, &regs.d)
	case 0x03:
		//fmt.Printf("RES 4, E: %X", op)
		res(to_res, &regs.e)
	case 0x04:
		//fmt.Printf("RES 4, H: %X", op)
		res(to_res, &regs.h)
	case 0x05:
		//fmt.Printf("RES 4, L: %X", op)
		res(to_res, &regs.l)
	case 0x06:
		//fmt.Printf("RES 4, (HL): %X", op)
		var r uint8 = readByte(regs.hl_read())
		res(to_res, &r)
		writeByte(regs.hl_read(), r)
		CicleCounter += 8
	case 0x07:
		//fmt.Printf("RES 4, A: %X", op)
		res(to_res, &regs.a)
	}
	CicleCounter += 8
}
func res_3_r(op uint8) {
	var to_res uint8 = 0x08
	switch (op & 0x0F){
	case 0x08:
		//fmt.Printf("RES 3, B: %X", op)
		res(to_res, &regs.b)
	case 0x09:
		//fmt.Printf("RES 3, C: %X", op)
		res(to_res, &regs.c)
	case 0x0A:
		//fmt.Printf("RES 3, D: %X", op)
		res(to_res, &regs.d)
	case 0x0B:
		//fmt.Printf("RES 3, E: %X", op)
		res(to_res, &regs.e)
	case 0x0C:
		//fmt.Printf("RES 3, H: %X", op)
		res(to_res, &regs.h)
	case 0x0D:
		//fmt.Printf("RES 3, L: %X", op)
		res(to_res, &regs.l)
	case 0x0E:
		//fmt.Printf("RES 3, (HL): %X", op)
		var r uint8 = readByte(regs.hl_read())
		res(to_res, &r)
		writeByte(regs.hl_read(), r)
		CicleCounter += 8
	case 0x0F:
		//fmt.Printf("RES 3, A: %X", op)
		res(to_res, &regs.a)
	}
	CicleCounter += 8
}
func res_2_r(op uint8) {
	var to_res uint8 = 0x04
	switch (op & 0x0F){
	case 0x00:
		//fmt.Printf("RES 2, B: %X", op)
		res(to_res, &regs.b)
	case 0x01:
		//fmt.Printf("RES 2, C: %X", op)
		res(to_res, &regs.c)
	case 0x02:
		//fmt.Printf("RES 2, D: %X", op)
		res(to_res, &regs.d)
	case 0x03:
		//fmt.Printf("RES 2, E: %X", op)
		res(to_res, &regs.e)
	case 0x04:
		//fmt.Printf("RES 2, H: %X", op)
		res(to_res, &regs.h)
	case 0x05:
		//fmt.Printf("RES 2, L: %X", op)
		res(to_res, &regs.l)
	case 0x06:
		//fmt.Printf("RES 2, (HL): %X", op)
		var r uint8 = readByte(regs.hl_read())
		res(to_res, &r)
		writeByte(regs.hl_read(), r)
		CicleCounter += 8
	case 0x07:
		//fmt.Printf("RES 2, A: %X", op)
		res(to_res, &regs.a)
	}
	CicleCounter += 8
}
func res_1_r(op uint8) {
	var to_res uint8 = 0x02
	switch (op & 0x0F){
	case 0x08:
		//fmt.Printf("RES 1, B: %X", op)
		res(to_res, &regs.b)
	case 0x09:
		//fmt.Printf("RES 1, C: %X", op)
		res(to_res, &regs.c)
	case 0x0A:
		//fmt.Printf("RES 1, D: %X", op)
		res(to_res, &regs.d)
	case 0x0B:
		//fmt.Printf("RES 1, E: %X", op)
		res(to_res, &regs.e)
	case 0x0C:
		//fmt.Printf("RES 1, H: %X", op)
		res(to_res, &regs.h)
	case 0x0D:
		//fmt.Printf("RES 1, L: %X", op)
		res(to_res, &regs.l)
	case 0x0E:
		//fmt.Printf("RES 1, (HL): %X", op)
		var r uint8 = readByte(regs.hl_read())
		res(to_res, &r)
		writeByte(regs.hl_read(), r)
		CicleCounter += 8
	case 0x0F:
		//fmt.Printf("RES 1, A: %X", op)
		res(to_res, &regs.a)
	}
	CicleCounter += 8
}
func res_0_r(op uint8) {
	var to_res uint8 = 0x01
	switch (op & 0x0F){
	case 0x00:
		//fmt.Printf("RES 0, B: %X", op)
		res(to_res, &regs.b)
	case 0x01:
		//fmt.Printf("RES 0, C: %X", op)
		res(to_res, &regs.c)
	case 0x02:
		//fmt.Printf("RES 0, D: %X", op)
		res(to_res, &regs.d)
	case 0x03:
		//fmt.Printf("RES 0, E: %X", op)
		res(to_res, &regs.e)
	case 0x04:
		//fmt.Printf("RES 0, H: %X", op)
		res(to_res, &regs.h)
	case 0x05:
		//fmt.Printf("RES 0, L: %X", op)
		res(to_res, &regs.l)
	case 0x06:
		//fmt.Printf("RES 0, (HL): %X", op)
		var r uint8 = readByte(regs.hl_read())
		res(to_res, &r)
		writeByte(regs.hl_read(), r)
		CicleCounter += 8
	case 0x07:
		//fmt.Printf("RES 0, A: %X", op)
		res(to_res, &regs.a)
	}
	CicleCounter += 8
}

func bit_7_r(op uint8) {
	var to_check uint8 = 0x80
	switch (op & 0x0F){
	case 0x08:
		//fmt.Printf("BIT 7, B: %X", op)
		bit(to_check, regs.b)
	case 0x09:
		//fmt.Printf("BIT 7, C: %X", op)
		bit(to_check, regs.c)
	case 0x0A:
		//fmt.Printf("BIT 7, D: %X", op)
		bit(to_check, regs.d)
	case 0x0B:
		//fmt.Printf("BIT 7, E: %X", op)
		bit(to_check, regs.e)
	case 0x0C:
		//fmt.Printf("BIT 7, H: %X", op)
		bit(to_check, regs.h)
	case 0x0D:
		//fmt.Printf("BIT 7, L: %X", op)
		bit(to_check, regs.l)
	case 0x0E:
		//fmt.Printf("BIT 7, HL: %X", op)
		bit(to_check, readByte(regs.hl_read()))
		CicleCounter += 8
	case 0x0F:
		//fmt.Printf("BIT 7, A: %X", op)
		bit(to_check, regs.a)
	}
	CicleCounter += 8
}
func bit_6_r(op uint8) {
	var to_check uint8 = 0x40
	switch (op & 0x0F){
	case 0x00:
		//fmt.Printf("BIT 6, B: %X", op)
		bit(to_check, regs.b)
	case 0x01:
		//fmt.Printf("BIT 6, C: %X", op)
		bit(to_check, regs.c)
	case 0x02:
		//fmt.Printf("BIT 6, D: %X", op)
		bit(to_check, regs.d)
	case 0x03:
		//fmt.Printf("BIT 6, E: %X", op)
		bit(to_check, regs.e)
	case 0x04:
		//fmt.Printf("BIT 6, H: %X", op)
		bit(to_check, regs.h)
	case 0x05:
		//fmt.Printf("BIT 6, L: %X", op)
		bit(to_check, regs.l)
	case 0x06:
		//fmt.Printf("BIT 6, HL: %X", op)
		bit(to_check, readByte(regs.hl_read()))
		CicleCounter += 8
	case 0x07:
		//fmt.Printf("BIT 6, A: %X", op)
		bit(to_check, regs.a)
	}
	CicleCounter += 8
}
func bit_5_r(op uint8) {
	var to_check uint8 = 0x20
	switch (op & 0x0F){
	case 0x08:
		//fmt.Printf("BIT 5, B: %X", op)
		bit(to_check, regs.b)
	case 0x09:
		//fmt.Printf("BIT 5, C: %X", op)
		bit(to_check, regs.c)
	case 0x0A:
		//fmt.Printf("BIT 5, D: %X", op)
		bit(to_check, regs.d)
	case 0x0B:
		//fmt.Printf("BIT 5, E: %X", op)
		bit(to_check, regs.e)
	case 0x0C:
		//fmt.Printf("BIT 5, H: %X", op)
		bit(to_check, regs.h)
	case 0x0D:
		//fmt.Printf("BIT 5, L: %X", op)
		bit(to_check, regs.l)
	case 0x0E:
		//fmt.Printf("BIT 5, HL: %X", op)
		bit(to_check, readByte(regs.hl_read()))
		CicleCounter += 8
	case 0x0F:
		//fmt.Printf("BIT 5, A: %X", op)
		bit(to_check, regs.a)
	}
	CicleCounter += 8
}
func bit_4_r(op uint8) {
	var to_check uint8 = 0x10
	switch (op & 0x0F){
	case 0x00:
		//fmt.Printf("BIT 4, B: %X", op)
		bit(to_check, regs.b)
	case 0x01:
		//fmt.Printf("BIT 4, C: %X", op)
		bit(to_check, regs.c)
	case 0x02:
		//fmt.Printf("BIT 4, D: %X", op)
		bit(to_check, regs.d)
	case 0x03:
		//fmt.Printf("BIT 4, E: %X", op)
		bit(to_check, regs.e)
	case 0x04:
		//fmt.Printf("BIT 4, H: %X", op)
		bit(to_check, regs.h)
	case 0x05:
		//fmt.Printf("BIT 4, L: %X", op)
		bit(to_check, regs.l)
	case 0x06:
		//fmt.Printf("BIT 4, HL: %X", op)
		bit(to_check, readByte(regs.hl_read()))
		CicleCounter += 8
	case 0x07:
		//fmt.Printf("BIT 4, A: %X", op)
		bit(to_check, regs.a)
	}
	CicleCounter += 8
}
func bit_3_r(op uint8) {
	var to_check uint8 = 0x08
	switch (op & 0x0F){
	case 0x08:
		//fmt.Printf("BIT 3, B: %X", op)
		bit(to_check, regs.b)
	case 0x09:
		//fmt.Printf("BIT 3, C: %X", op)
		bit(to_check, regs.c)
	case 0x0A:
		//fmt.Printf("BIT 3, D: %X", op)
		bit(to_check, regs.d)
	case 0x0B:
		//fmt.Printf("BIT 3, E: %X", op)
		bit(to_check, regs.e)
	case 0x0C:
		//fmt.Printf("BIT 3, H: %X", op)
		bit(to_check, regs.h)
	case 0x0D:
		//fmt.Printf("BIT 3, L: %X", op)
		bit(to_check, regs.l)
	case 0x0E:
		//fmt.Printf("BIT 3, HL: %X", op)
		bit(to_check, readByte(regs.hl_read()))
		CicleCounter += 8
	case 0x0F:
		//fmt.Printf("BIT 3, A: %X", op)
		bit(to_check, regs.a)
	}
	CicleCounter += 8
}
func bit_2_r(op uint8) {
	var to_check uint8 = 0x04
	switch (op & 0x0F){
	case 0x00:
		//fmt.Printf("BIT 2, B: %X", op)
		bit(to_check, regs.b)
	case 0x01:
		//fmt.Printf("BIT 2, C: %X", op)
		bit(to_check, regs.c)
	case 0x02:
		//fmt.Printf("BIT 2, D: %X", op)
		bit(to_check, regs.d)
	case 0x03:
		//fmt.Printf("BIT 2, E: %X", op)
		bit(to_check, regs.e)
	case 0x04:
		//fmt.Printf("BIT 2, H: %X", op)
		bit(to_check, regs.h)
	case 0x05:
		//fmt.Printf("BIT 2, L: %X", op)
		bit(to_check, regs.l)
	case 0x06:
		//fmt.Printf("BIT 2, HL: %X", op)
		bit(to_check, readByte(regs.hl_read()))
		CicleCounter += 8
	case 0x07:
		//fmt.Printf("BIT 2, A: %X", op)
		bit(to_check, regs.a)
	}
	CicleCounter += 8
}
func bit_1_r(op uint8) {
	var to_check uint8 = 0x02
	switch (op & 0x0F){
	case 0x08:
		//fmt.Printf("BIT 1, B: %X", op)
		bit(to_check, regs.b)
	case 0x09:
		//fmt.Printf("BIT 1, C: %X", op)
		bit(to_check, regs.c)
	case 0x0A:
		//fmt.Printf("BIT 1, D: %X", op)
		bit(to_check, regs.d)
	case 0x0B:
		//fmt.Printf("BIT 1, E: %X", op)
		bit(to_check, regs.e)
	case 0x0C:
		//fmt.Printf("BIT 1, H: %X", op)
		bit(to_check, regs.h)
	case 0x0D:
		//fmt.Printf("BIT 1, L: %X", op)
		bit(to_check, regs.l)
	case 0x0E:
		//fmt.Printf("BIT 1, HL: %X", op)
		bit(to_check, readByte(regs.hl_read()))
		CicleCounter += 8
	case 0x0F:
		//fmt.Printf("BIT 1, A: %X", op)
		bit(to_check, regs.a)
	}
	CicleCounter += 8
}
func bit_0_r(op uint8) {
	var to_check uint8 = 0x01
	switch (op & 0x0F){
	case 0x00:
		//fmt.Printf("BIT 0, B: %X", op)
		bit(to_check, regs.b)
	case 0x01:
		//fmt.Printf("BIT 0, C: %X", op)
		bit(to_check, regs.c)
	case 0x02:
		//fmt.Printf("BIT 0, D: %X", op)
		bit(to_check, regs.d)
	case 0x03:
		//fmt.Printf("BIT 0, E: %X", op)
		bit(to_check, regs.e)
	case 0x04:
		//fmt.Printf("BIT 0, H: %X", op)
		bit(to_check, regs.h)
	case 0x05:
		//fmt.Printf("BIT 0, L: %X", op)
		bit(to_check, regs.l)
	case 0x06:
		//fmt.Printf("BIT 0, HL: %X", op)
		bit(to_check, readByte(regs.hl_read()))
		CicleCounter += 8
	case 0x07:
		//fmt.Printf("BIT 0, A: %X", op)
		bit(to_check, regs.a)
	}
	CicleCounter += 8
}

func srl_a(op uint8) {
	//fmt.Printf("SRL A: %X", op)
	srl(&regs.a)
	CicleCounter += 8
}
func srl_hl(op uint8) {
	//fmt.Printf("SRL HL: %X", op)
	var r uint8 = readByte(regs.hl_read())
	srl(&r)
	writeByte(regs.hl_read(), r)
	CicleCounter += 8
}
func srl_l(op uint8) {
	//fmt.Printf("SRL L: %X", op)
	srl(&regs.l)
	CicleCounter += 8
}
func srl_h(op uint8) {
	//fmt.Printf("SRL H: %X", op)
	srl(&regs.h)
	CicleCounter += 8
}
func srl_e(op uint8) {
	//fmt.Printf("SRL E: %X", op)
	srl(&regs.e)
	CicleCounter += 8
}
func srl_d(op uint8) {
	//fmt.Printf("SRL D: %X", op)
	srl(&regs.d)
	CicleCounter += 8
}
func srl_c(op uint8) {
	//fmt.Printf("SRL C: %X", op)
	srl(&regs.c)
	CicleCounter += 8
}
func srl_b(op uint8) {
	//fmt.Printf("SRL B: %X", op)
	srl(&regs.b)
	CicleCounter += 8
}

func swap_a(op uint8) {
	//fmt.Printf("SWAP A: %X", op)
	swap(&regs.a)
	CicleCounter += 8
}
func swap_hl(op uint8) {
	//fmt.Printf("SWAP (HL): %X", op)
	var r uint8 = readByte(regs.hl_read())
	swap(&r)
	writeByte(regs.hl_read(), r)
	CicleCounter += 16
}
func swap_l(op uint8) {
	//fmt.Printf("SWAP L: %X", op)
	swap(&regs.l)
	CicleCounter += 8
}
func swap_h(op uint8) {
	//fmt.Printf("SWAP H: %X", op)
	swap(&regs.h)
	CicleCounter += 8
}
func swap_e(op uint8) {
	//fmt.Printf("SWAP E: %X", op)
	swap(&regs.e)
	CicleCounter += 8
}
func swap_d(op uint8) {
	//fmt.Printf("SWAP D: %X", op)
	swap(&regs.d)
	CicleCounter += 8
}
func swap_c(op uint8) {
	//fmt.Printf("SWAP C: %X", op)
	swap(&regs.c)
	CicleCounter += 8
}
func swap_b(op uint8) {
	//fmt.Printf("SWAP B: %X", op)
	swap(&regs.b)
	CicleCounter += 8
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

func swap(r *uint8) {
	regs.clearFlags(SUBTRACT|HALFCARRY|CARRY)

	*r = ((*r & 0x0F) << 4) | ((*r & 0xF0) >> 4) 

	if *r == 0 {
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
}

func srl(r *uint8) {
	regs.clearFlags(SUBTRACT|HALFCARRY)
	if (*r & 0x01) != 0 {
		regs.setFlags(CARRY)
	} else{
		regs.clearFlags(CARRY)
	}

	*r = *r >> 1

	if *r == 0 {
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
}

func bit(b uint8, r uint8) {
	if (r & b) == 0 {
		regs.setFlags(ZERO)
	} else{
		regs.clearFlags(ZERO)
	}
	regs.clearFlags(SUBTRACT)
	regs.setFlags(HALFCARRY)
}

func res(b uint8, r *uint8) {
	*r = *r & (^b)
}

func set(b uint8, r *uint8) {
	*r = *r | b
}
