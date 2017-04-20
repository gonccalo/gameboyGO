package main

type genericCart struct {
	head      rom_header
	RomData []uint8
	cartRam []uint8
}

func (r *genericCart)init(data []uint8, filename string) bool{
	r.head.init(data)
	r.cartRam = make([]uint8, 0x1FFF)
	r.RomData = make([]uint8, len(data))
	if num := copy(r.RomData, data); num != len(data){
		return false
	}
	return true
}
func (r *genericCart)romRead(addr uint16) uint8 {
	return r.RomData[addr]
}
func (r *genericCart)romWrite(addr uint16, b uint8) {
	return
}
func (r *genericCart)ramRead(addr uint16) uint8 {
	return r.cartRam[addr]
}
func (r *genericCart)ramWrite(addr uint16, b uint8) {
	r.cartRam[addr] = b
}
func (r *genericCart)saveCartRam() {
	return
}