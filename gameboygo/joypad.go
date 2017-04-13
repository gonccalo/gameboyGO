package main
/*
                 P14        P15
                  |          |
        P10-------O-Right----O-A
                  |          |
        P11-------O-Left-----O-B
                  |          |
        P12-------O-Up-------O-Select
                  |          |
        P13-------O-Down-----O-Start
                  |          |
		keys    1111       1111
*/
type key uint8
const(
	KEY_A 		key = 1 << iota 
	KEY_B
	KEY_SELECT
	KEY_START
	KEY_RIGHT
	KEY_LEFT
	KEY_UP	
	KEY_DOWN
)
var keys uint8
func init() {
	keys = 0xFF
}
func getKeys(sel uint8) uint8{
	switch sel & 0x30{
	case 0x30:		//none selected
		return (0xC0|(sel & 0xF0)) | 0x0F
	case 0x10:		//P15 selected (a,b,select,start)
		return (0xC0|(sel & 0xF0)) | (keys & 0x0F)
 	case 0x20:		//P14 selected (direction keys)
		return (0xC0|(sel & 0xF0)) | ((keys & 0xF0)>>4)
	default:
		return (0xC0|(sel & 0xF0)) | 0x0F
	}
}
func KeyPressed(k key) {
	keys = keys &^ uint8(k)
	setInterruptsFlag(JOYPAD)
}
func KeyReleased(k key) {
	keys |= uint8(k)
}