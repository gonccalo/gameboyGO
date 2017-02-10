package gameboygo
import (
	"fmt"
)
/*
registos 
AF   --> accumulator e flags
BC   --> pode ser acedido como B ou C 
DE   -->
HL   -->
SP   --> stack pointer
PC   --> program counter
*/

type flag uint8
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
const(
	ZERO 		flag = 0x80
	SUBTRACT 	flag = 0x40
	HALFCARRY	flag = 0x20
	CARRY		flag = 0X10
)
func (r *registers)af_write(data uint16) {
	r.a = uint8((data & 0xFF00) >> 8)
	r.f = uint8(data & 0x00FF)
}
func (r *registers)af_read() uint16{
	return uint16((uint16(r.a) << 8) | uint16(r.f)) 
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
func (r *registers)setFlags(f flag){
	r.f |= uint8(f)
}
func (r *registers)clearFlags(f flag) {
	r.f = r.f &^ uint8(f)
}
func (r *registers)getFlag(f flag) bool{
	return (r.f & uint8(f)) > 0
}
func (r *registers)String() string {
	return fmt.Sprintf("Flags:\nZERO: %t\nSUBTRACT: %t\nHALFCARRY: %t\nCARRY: %t\nRegistos:\na: %X\nf: %X\nb: %X\nc: %X\nd: %X\ne: %X\nh: %X\nl: %X\n\nsp: %X\npc: %X\n", r.getFlag(ZERO), r.getFlag(SUBTRACT),r.getFlag(HALFCARRY), r.getFlag(CARRY), r.a, r.f, r.b, r.c, r.d, r.e, r.h, r.l, r.sp, r.pc)
}