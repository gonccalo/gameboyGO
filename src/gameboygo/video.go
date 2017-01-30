package gameboygo

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

/*
TODO
NOT IMPLEMENTED
LCD Color Palettes (CGB only)
FF68 - BCPS/BGPI - CGB Mode Only - Background Palette Index
FF69 - BCPD/BGPD - CGB Mode Only - Background Palette Data
FF6A - OCPS/OBPI - CGB Mode Only - Sprite Palette Index
FF6B - OCPD/OBPD - CGB Mode Only - Sprite Palette Data

LCD VRAM Bank (CGB only)
FF4F - VBK - CGB Mode Only - VRAM Bank

LCD VRAM DMA Transfers (CGB only)

FF51 - HDMA1 - CGB Mode Only - New DMA Source, High
FF52 - HDMA2 - CGB Mode Only - New DMA Source, Low
FF53 - HDMA3 - CGB Mode Only - New DMA Destination, High
FF54 - HDMA4 - CGB Mode Only - New DMA Destination, Low
FF55 - HDMA5 - CGB Mode Only - New DMA Length/Mode/Start
*/
