package main

import "github.com/veandco/go-sdl2/sdl"
import "unsafe"
//import "fmt"
const(
	WHITE		uint8 = 255
	LIGHT_GRAY	uint8 = 170
	DARK_GRAY	uint8 = 85
	BLACK		uint8 = 0

	SCAN_CICLES int   = 456

	LCD_STAT_HBLANK 	 uint8 = 0x00   
	LCD_STAT_VBLANK 	 uint8 = 0x01
	LCD_STAT_OAM_RAM 	 uint8 = 0x02
	LCD_STAT_DATA2DRIVER uint8 = 0x03

	WIDTH		int   = 160
	HEIGHT		int   = 144
	PITCH   	int   = WIDTH * 4
	BUFFER_SIZE int   = WIDTH*HEIGHT*4
)
var colors = [4]uint8 {WHITE,		//00
					   LIGHT_GRAY,	//01
					   DARK_GRAY,	//10
					   BLACK,}		//11

var pitch int = PITCH
var p_pixels unsafe.Pointer
/*
  Bit 7 - LCD Display Enable             (0=Off, 1=On)
  Bit 6 - Window Tile Map Display Select (0=9800-9BFF, 1=9C00-9FFF)
  Bit 5 - Window Display Enable          (0=Off, 1=On)
  Bit 4 - BG & Window Tile Data Select   (0=8800-97FF, 1=8000-8FFF)
  Bit 3 - BG Tile Map Display Select     (0=9800-9BFF, 1=9C00-9FFF)
  Bit 2 - OBJ (Sprite) Size              (0=8x8, 1=8x16)
  Bit 1 - OBJ (Sprite) Display Enable    (0=Off, 1=On)
  Bit 0 - BG Display (for CGB see below) (0=Off, 1=On)
*/
var LcdControl = &ram[0xFF40]

/* 
  Bit 6 - LYC=LY Coincidence Interrupt (1=Enable) (Read/Write)
  Bit 5 - Mode 2 OAM Interrupt         (1=Enable) (Read/Write)
  Bit 4 - Mode 1 V-Blank Interrupt     (1=Enable) (Read/Write)
  Bit 3 - Mode 0 H-Blank Interrupt     (1=Enable) (Read/Write)
  Bit 2 - Coincidence Flag  (0:LYC<>LY, 1:LYC=LY) (Read Only)
  Bit 1-0 - Mode Flag       (Mode 0-3, see below) (Read Only)
            0: During H-Blank
            1: During V-Blank
            2: During Searching OAM-RAM
            3: During Transfering Data to LCD Driver
*/
var LcdStatus = &ram[0xFF41]

/*
Specifies the position in the 256x256 pixels BG map (32x32 tiles) which is to be displayed at the upper/left LCD display position.
Values in range from 0-255 may be used for X/Y each, the video controller automatically wraps back to the upper (left) position in BG map when drawing exceeds the lower (right) border of the BG map area.
*/
var ScY = &ram[0xFF42]
var ScX = &ram[0xFF43]

/*
The LY indicates the vertical line to which the present data is transferred to the LCD Driver. The LY can take on any value between 0 through 153. The values between 144 and 153 indicate the V-Blank period. Writing will reset the counter.
*/
var Ly  = &ram[0xFF44]

/*
The gameboy permanently compares the value of the LYC and LY registers. When both values are identical, the coincident bit in the STAT register becomes set, and (if enabled) a STAT interrupt is requested.
*/
var LyC = &ram[0xFF45]

/*
Specifies the upper/left positions of the Window area. (The window is an alternate background area which can be displayed above of the normal background. OBJs (sprites) may be still displayed above or behinf the window, just as for normal BG.)
The window becomes visible (if enabled) when positions are set in range WX=0..166, WY=0..143. A postion of WX=7, WY=0 locates the window at upper left, it is then completly covering normal background.
*/
var Wy  = &ram[0xFF4A]
var Wx  = &ram[0xFF4B]

/*
This register assigns gray shades to the color numbers of the BG and Window tiles.
  Bit 7-6 - Shade for Color Number 3
  Bit 5-4 - Shade for Color Number 2
  Bit 3-2 - Shade for Color Number 1
  Bit 1-0 - Shade for Color Number 0
*/
var Bgp = &ram[0xFF47]

/*
ram[0xFF48]
This register assigns gray shades for sprite palette 0. It works exactly as BGP (FF47), except that the lower two bits aren't used because sprite data 00 is transparent.

ram[0xFF49]
This register assigns gray shades for sprite palette 1. It works exactly as BGP (FF47), except that the lower two bits aren't used because sprite data 00 is transparent.
*/

var LastScanLine int = 0
func UpdateGPU(renderer *sdl.Renderer, tex *sdl.Texture) {
	if (*LcdControl & 0x80) == 0{
		//display off
		//must clean Ly and set mode 1
		LastScanLine = 0
		*Ly = 0
		setLcdStatMode(LCD_STAT_VBLANK)
		return
	}
	if *Ly >= 144 {
		if ((*LcdStatus & 0x10) != 0) && ((*LcdStatus & 0x03) != LCD_STAT_VBLANK) {
			//mode 1: During V-Blank
			setInterruptsFlag(LCD_STAT)
		}
		setLcdStatMode(LCD_STAT_VBLANK)
	} else{
		//maybe i should redo all this CicleCounter thing
		var modeCicles = CicleCounter % SCAN_CICLES
		if modeCicles < 80{
			//mode 2: 80 cicles of the 456
			if ((*LcdStatus & 0x20) != 0) && ((*LcdStatus & 0x03) != LCD_STAT_OAM_RAM) {
				//if this interrupt is enabled and just changed mode
				setInterruptsFlag(LCD_STAT)
			}
			setLcdStatMode(LCD_STAT_OAM_RAM)
		} else if modeCicles < (80 + 172) {
			//mode 3: 172 cicles of the 456
			setLcdStatMode(LCD_STAT_DATA2DRIVER)
		} else{
			//mode 0: remaining cicles
			if ((*LcdStatus & 0x08) != 0) && ((*LcdStatus & 0x03) != LCD_STAT_HBLANK) {
				//if this interrupt is enabled and just changed mode
				setInterruptsFlag(LCD_STAT)
			}
			setLcdStatMode(LCD_STAT_HBLANK)
		}
	}

	if (*Ly == *LyC) {
		//Bit 2 - Coincidence Flag  (0:LYC<>LY, 1:LYC=LY) (Read Only)
		if (*LcdStatus & 0x40) != 0 && (*LcdStatus & 0x04) == 0{
			setInterruptsFlag(LCD_STAT)
		}
		*LcdStatus |= 0x04
	} else{
		*LcdStatus &= 0xFB
	}
	if (int(CicleCounter/SCAN_CICLES) - LastScanLine) >= 1{ 
		//next scanline
		LastScanLine = int(CicleCounter/SCAN_CICLES)
		*Ly += 1
		if *Ly == 144 { 
			//VBLANK 
			setInterruptsFlag(V_BLANK)
		} else if *Ly > 153 {  
			//end VBLANK
			*Ly = 255
			tex.Unlock()
			renderer.Copy(tex,nil,nil)
			renderer.Present()
		} else if *Ly < 144{
			tex.Lock(nil, &p_pixels, &pitch)
			DrawLine(renderer)
		}
	}
}
func DrawLine(renderer *sdl.Renderer) {
	if (*LcdControl & 0x01) != 0 { //draw background
		var signed bool
		var tileMap uint16
		var tileData uint16
		var currentTile uint8
		var drawWindow bool
		if (*LcdControl & 0x10) == 0{  //Bit 4 - BG & Window Tile Data Select   (0=8800-97FF, 1=8000-8FFF)
			tileData = 0x8800
			signed = true
		} else{
			tileData = 0x8000
			signed = false
		}
		if ((*LcdControl & 0x20) != 0) && (*Ly > *Wy) {      //Bit 5 - Window Display Enable(0=Off, 1=On)
			//drawing window
			drawWindow = true
			if (*LcdControl & 0x40) == 0 { //Bit 6 - Window Tile Map Display Select (0=9800-9BFF, 1=9C00-9FFF)
				tileMap = 0x9800
			} else {
				tileMap = 0x9C00
			}
			currentTile = *Ly - *Wy
		} else{
			drawWindow = false
			if (*LcdControl & 0x08) == 0 { //Bit 3 - BG Tile Map Display Select (0=9800-9BFF, 1=9C00-9FFF)
				tileMap = 0x9800
			} else{
				tileMap = 0x9C00
			}
			currentTile = *ScY + *Ly
		}
		/*
		tiles are 8*8 pixels
		Background Tile Map contains the numbers of tiles to be displayed. It is organized
 		as 32 rows of 32 bytes each. Each byte contains a number of a tile to be displayed
 		*/
		var tileLine uint16 = (uint16(currentTile)/8)*32
		var x uint8
		var i uint8
		for i = 0; i < 160; i++ {
			if drawWindow && (i >= (*Wx-7)) { // window
				x = i - (*Wx-7)
			} else{
				x = i + *ScX
			}
			var tileNumAddress uint16 = tileMap + tileLine + (uint16(x)/8)
			var tileAddr uint16
			if signed {
				tileAddr = tileData + (uint16( int16(int8(ram[tileNumAddress])) + 128 ) * 16) //maybe this works
			} else{
				tileAddr = tileData + (uint16(ram[tileNumAddress]) * 16)
			}
			var l uint16 = (uint16(currentTile) % 8) * 2 //line of the tile that we are drawing: each tile has 8 lines and each line is 2 bytes
			var data0 = ram[tileAddr + l + 0]
			var data1 = ram[tileAddr + l + 1]
			var c = getPixelColor(data0, data1, 7-(x%8), *Bgp)  // ^x & 0x07
			pos := ((int(*Ly)*WIDTH)*4) + (int(i)*4)
			(*[BUFFER_SIZE]uint8)(p_pixels)[pos + 0] = c
			(*[BUFFER_SIZE]uint8)(p_pixels)[pos + 1] = c
			(*[BUFFER_SIZE]uint8)(p_pixels)[pos + 2] = c
			(*[BUFFER_SIZE]uint8)(p_pixels)[pos + 3] = 255
		}
	}
	if (*LcdControl & 0x02) != 0 { //draw sprites
		var sy uint8 = 8 + ((*LcdControl & 0x04) << 1)
		/*
		Byte0  Y position on the screen
  		Byte1  X position on the screen
  		Byte2  Pattern number 0-255 (Unlike some tile
         numbers, sprite pattern numbers are unsigned.
         LSB is ignored (treated as 0) in 8x16 mode.)
  		Byte3  Flags
		*/
		var i uint16
		for i = 0; i < 40; i++ { //all 40 sprites at 0xFE00 (4 bytes per sprite)
			var spriteY     uint8 = ram[0xFE00 + (i*4) + 0] - 16
			var spriteX     uint8 = ram[0xFE00 + (i*4) + 1] -  8
			var spriteNum   uint8 = ram[0xFE00 + (i*4) + 2] //LSB = 0 when 8x16
			var spriteFlags uint8 = ram[0xFE00 + (i*4) + 3]

			if !((*Ly >= spriteY) && (*Ly < (spriteY+sy))) {
				//sprite not in current scanline
				continue
			}
			var line uint8 = *Ly - spriteY
			if (spriteFlags & 0x40) != 0 {
				//yFlip: read from bottom
				line = (^line) & 0x07
			}

			//palette --> (0=OBP0(0xFF48), 1=OBP1(0xFF49))
			var palette uint8 = ram[0xFF48 + ((uint16(spriteFlags) & 0x10) >> 4) ]

			var spriteDataAddr uint16 = 0x8000 + (uint16(spriteNum) * 16) + (uint16(line)*2)
			var data0 = ram[spriteDataAddr + 0]
			var data1 = ram[spriteDataAddr + 1]
			var Xpixel uint8
			for Xpixel = 7; Xpixel != 255 ; Xpixel-- {
				//draw all pixels from this line
				var bit uint8 = Xpixel
				if (spriteFlags & 0x20) != 0 {
					//X flip (0=Normal, 1=Horizontally mirrored)
					bit = (^Xpixel) & 0x07
				}
				color := getPixelColor(data0, data1, bit, palette)
				if color == WHITE{
					//the white color in sprites is transparent
					continue
				}
				pos :=  (int(*Ly)*WIDTH*4) + (int(spriteX) + int((^Xpixel)&0x07)) * 4
				(*[BUFFER_SIZE]uint8)(p_pixels)[pos + 0] = color
				(*[BUFFER_SIZE]uint8)(p_pixels)[pos + 1] = color
				(*[BUFFER_SIZE]uint8)(p_pixels)[pos + 2] = color
				(*[BUFFER_SIZE]uint8)(p_pixels)[pos + 3] = 255
			}
		}
	}
}

func getPixelColor(lower, upper, bitNum uint8, palette uint8) uint8{
	var mask uint8 = (1 << bitNum)
	if bitNum < 2 {
		return getColor(((upper & mask) << (1-bitNum)) | ((lower & mask) >> bitNum), palette)
	}
	return getColor((upper & mask) >> (bitNum - 1) | ((lower & mask) >> bitNum), palette)
}

func getColor(code uint8, palette uint8) uint8{
	return colors[((palette & (0x03 << (2*code))) >> (2*code))]
}

func setLcdStatMode(mode uint8) {
	*LcdStatus &= 0xFC //clean last mode
	*LcdStatus |= mode //set
}
