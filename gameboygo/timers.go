package main
import "fmt"
var(
	ciclesThisUpdate int
	tfreq 			 int
	ciclesToDiv		 int
)

func setupTimers() {
	if (ram[0xFF07] & 0x04) == 0{
		tfreq = 0
		return
	}
	lastTfreq := tfreq
	switch ram[0xFF07] & 0x03{
		case 0x00:
			tfreq = CPU_FREQ/4096
		case 0x01:
			tfreq = CPU_FREQ/262144
		case 0x02:
			tfreq = CPU_FREQ/65536
		case 0x03:
			tfreq = CPU_FREQ/16384
		default:
			tfreq = CPU_FREQ/4096
	}
	if lastTfreq != tfreq {
		//reset counter to start over
		ciclesThisUpdate = 0
	}
}
func incTimer(cicles int) {
	if tfreq != 0 {
		ciclesThisUpdate += cicles
		if ciclesThisUpdate >= tfreq {
			ciclesThisUpdate = 0
			ram[0xFF05]++
			if ram[0xFF05] == 255 {
				ram[0xFF05] = ram[0xFF06]
				setInterruptsFlag(TIMER)
				fmt.Println("TIMER INTERRUPT")
			}
		}
	}
	ciclesToDiv += cicles
	if ciclesToDiv >= (CPU_FREQ/16384) {
		ram[0xFF04]++
		ciclesToDiv = 0
	}
}