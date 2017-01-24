package gameboygo

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

var regs registers

func reset() {
	regs.a = 0x01
	regs.f = 0xB0
	regs.b = 0x00
	regs.c = 0x13
	regs.d = 0x00
	regs.e = 0xD8
	regs.h = 0x01
	regs.l = 0x4D
	regs.sp = 0xFFFE
}