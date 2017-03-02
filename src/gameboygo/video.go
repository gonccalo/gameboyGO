package gameboygo

import "github.com/veandco/go-sdl2/sdl"
const(
	WHITE		uint8 = 255
	LIGHT_GRAY	uint8 = 170
	DARK_GRAY	uint8 = 85
	BLACK		uint8 = 0
)
var colors [4] uint8 = {WHITE,		//00
					   LIGHT_GRAY,	//01
					   DARK_GRAY,	//10
					   BLACK		//11
					  }

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
This register assigns gray shades for sprite palette 0. It works exactly as BGP (FF47), except that the lower two bits aren't used because sprite data 00 is transparent.
*/
var Obp0 = &ram[0xFF48]

/*
This register assigns gray shades for sprite palette 1. It works exactly as BGP (FF47), except that the lower two bits aren't used because sprite data 00 is transparent.
*/
var Obp1 = &ram[0xFF49]

func DrawLine(renderer *sdl.Renderer) {
	if (*LcdControl & 0x80) == 0 {  //display off
		return
	}
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
		
		if ((*LcdControl & 0x20) == 1) && (*Ly > *Wy) {      //Bit 5 - Window Display Enable(0=Off, 1=On)
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
		TODO: 8*16 pix tiles
		Background Tile Map contains the numbers of tiles to be displayed. It is organized
 		as 32 rows of 32 bytes each. Each byte contains a number of a tile to be displayed
 		*/
		var tileLine uint16 = (uint16(currentTile)/8)*32
		var x uint8
		var i uint8
		for i = 0; i < 160; i++ {
			if drawWindow && (i >= *Wx) { // window
				x = i - *Wx
			} else{
				x = i + *ScX
			}
			var tileNumAddress uint16 = tileMap + tileLine + (uint16(x)/8)
			var tileAddr uint16
			if signed {
				tileAddr = tileData + (uint16( int16(int8(ram[tileNumAddr])) + 128 ) * 16) //maybe this works
			} else{
				tileAddr = tileData + (uint16(ram[tileNumAddress]) * 16)
			}
			var l uint16 = (currentTile % 8) * 2 //line of the tile that we are drawing: each tile has 8 lines and each line is 2 bytes
			var data0 = ram[tileAddr + l + 0]
			var data1 = ram[tileAddr + l + 1]
			var c = getPixelColor(data0, data1, 7-(x%8))
			renderer.SetDrawColor(c,c,c,255)
			renderer.DrawPoint(i,*Ly)
		}
	}
	if (*LcdControl & 0x02) != 0 { //draw sprites
		
	}
}

func getPixelColor(lower, upper, bitNum uint8) uint{
	var mask uint8 = (1 << bitNum)
	if bitNum < 2 {
		return getColor(((upper & mask) << (1-bitNum)) | ((lower & mask) >> bitNum))
	}
	return getColor((upper & mask) >> (bitNum - 1) | ((lower & mask) >> bitNum))
}

func getColor(code uint8) uint{
	return colors[((*Bgp & (0x03 << 2*code)) >> 2*code)]
}