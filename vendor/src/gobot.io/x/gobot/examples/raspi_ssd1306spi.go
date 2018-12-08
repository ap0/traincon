// +build example
//
// Do not build by default.

package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/spi"
	"gobot.io/x/gobot/platforms/raspi"
)

// this example only works for a 128x64 display
var gobotLogo = []byte{0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xf8, 0xe0, 0xb0, 0x90, 0xc8, 0x6e, 0x9a, 0xb6, 0xd, 0x3a, 0x15, 0xf7, 0xd, 0x59, 0x98, 0x94, 0xf4, 0xf4, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xfc, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f, 0x7f, 0xff, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xdf, 0xff, 0x7f, 0xff, 0xff, 0xff, 0x8f, 0x8f, 0xff, 0x97, 0xcf, 0x4f, 0xc3, 0x51, 0xc0, 0x41, 0xd1, 0x40, 0xa4, 0xc4, 0x50, 0xc0, 0x40, 0xd1, 0x87, 0xdf, 0x97, 0x77, 0x3f, 0x8f, 0x4f, 0xff, 0x9f, 0xbf, 0x9f, 0x9f, 0xff, 0xff, 0xff, 0x7f, 0x7f, 0xff, 0x7f, 0xff, 0xff, 0xff, 0xff, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f, 0x3f, 0x37, 0x37, 0xef, 0xda, 0xbf, 0xef, 0x8e, 0x8d, 0xe7, 0xad, 0xfb, 0xaf, 0x1b, 0xdb, 0x29, 0x1f, 0x56, 0xcf, 0x1b, 0xcf, 0x9b, 0x55, 0x8f, 0xdd, 0x1a, 0xd7, 0x1e, 0xb5, 0x9b, 0xad, 0x17, 0xbd, 0xab, 0x15, 0xbf, 0x2d, 0x9a, 0xb7, 0x1d, 0xb7, 0x2d, 0x9a, 0xb7, 0x9e, 0x95, 0x9f, 0x55, 0x9b, 0x5d, 0x97, 0x5a, 0x8f, 0xda, 0xf, 0xda, 0x4f, 0x9a, 0x2e, 0x5a, 0xae, 0x14, 0xff, 0x67, 0xcf, 0x8b, 0x56, 0xdd, 0x74, 0xdd, 0x77, 0x57, 0x5f, 0x5f, 0xdf, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfc, 0xfc, 0xfe, 0xff, 0xf2, 0xff, 0xfc, 0xef, 0xba, 0xef, 0xba, 0xed, 0xbf, 0xea, 0x5f, 0xfa, 0x56, 0xfe, 0xb4, 0xde, 0x74, 0xbe, 0xec, 0x5a, 0xfc, 0x55, 0xfc, 0xd5, 0xbc, 0x74, 0xad, 0x7c, 0xa9, 0xfc, 0xa5, 0xec, 0x9d, 0xc8, 0xcd, 0xac, 0x79, 0xd5, 0x7c, 0xb4, 0xdc, 0xb4, 0xdd, 0x74, 0xdd, 0xf4, 0xac, 0xfd, 0xaa, 0xfc, 0x56, 0xfc, 0xae, 0xfa, 0xad, 0xfe, 0xaa, 0x7f, 0xed, 0x5b, 0xfe, 0x55, 0xff, 0xff, 0xab, 0xaf, 0xf9, 0xf8, 0xfc, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f, 0x7f, 0x7f, 0x7f, 0x3f, 0x3f, 0x3f, 0x3f, 0x3f, 0x1f, 0x1f, 0x6e, 0xfb, 0x6e, 0xdb, 0x7e, 0xd5, 0x7f, 0xd5, 0x7f, 0xd5, 0xff, 0xaa, 0xff, 0x55, 0xff, 0x55, 0xff, 0x55, 0xff, 0xd6, 0xfd, 0xd7, 0x7d, 0xeb, 0xff, 0xaa, 0xff, 0xdb, 0x7f, 0xf5, 0xdf, 0x7b, 0xef, 0xfb, 0xed, 0xbf, 0xea, 0x7f, 0xea, 0xbf, 0xeb, 0xbe, 0xeb, 0x5e, 0xfb, 0x56, 0xff, 0x55, 0xff, 0xaa, 0xff, 0x56, 0xfd, 0x57, 0xfd, 0x57, 0xfd, 0x57, 0xfd, 0xb7, 0xb7, 0xda, 0x1a, 0x3f, 0x3f, 0x3f, 0x3f, 0x7f, 0x7f, 0x7f, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f, 0x7f, 0x7f, 0x3f, 0x3f, 0x3f, 0x3f, 0x3f, 0xff, 0xff, 0xf, 0xf, 0xf, 0x1, 0x20, 0x8, 0x40, 0x10, 0x84, 0x0, 0xa8, 0x0, 0x2, 0xa8, 0x0, 0x2, 0xa8, 0x0, 0x42, 0xfd, 0x9f, 0x75, 0xdf, 0xb5, 0xdf, 0xb5, 0xff, 0xaa, 0x7f, 0x2a, 0x5f, 0x75, 0x3f, 0x6d, 0x7f, 0x5b, 0x7f, 0xf6, 0x3f, 0xed, 0x7f, 0xfb, 0xaf, 0x7e, 0xfb, 0x6f, 0xbd, 0xf7, 0x5f, 0xfd, 0x77, 0xde, 0x7f, 0xf5, 0x5f, 0x7f, 0xf5, 0x5f, 0xff, 0x35, 0xff, 0x57, 0x7d, 0x6f, 0x3d, 0x6b, 0x3f, 0x2a, 0xff, 0x55, 0xff, 0x95, 0x7f, 0x55, 0xff, 0x95, 0x6f, 0xa8, 0x40, 0x0, 0x2a, 0x0, 0x44, 0x10, 0x4, 0x40, 0x10, 0x84, 0x0, 0x28, 0x0, 0x20, 0x1, 0xf, 0x2f, 0x2f, 0x3f, 0x3f, 0x3f, 0x3f, 0x3f, 0x3f, 0x3f, 0x3f, 0x3f, 0x3f, 0x3f, 0x3f, 0x3f, 0x3f, 0x3f, 0xff, 0x7, 0x7, 0x7, 0xa1, 0x0, 0x48, 0xa0, 0xd4, 0x30, 0xe8, 0x1a, 0x4, 0x18, 0x44, 0xc, 0x80, 0xe, 0x20, 0x84, 0x8c, 0x0, 0xcd, 0x8, 0x88, 0x92, 0x8, 0xa0, 0x2, 0x8, 0x42, 0xb0, 0x60, 0x9a, 0xf0, 0xc, 0x10, 0x4d, 0x8, 0x85, 0x24, 0x5, 0x48, 0x4, 0x84, 0x2c, 0x0, 0x8c, 0x19, 0x68, 0xb2, 0x48, 0xf1, 0x0, 0x44, 0x10, 0x0, 0x6c, 0xd8, 0x24, 0xfc, 0x80, 0x4c, 0x84, 0x48, 0x84, 0x4c, 0x80, 0x4c, 0x84, 0x48, 0xd4, 0x38, 0xe4, 0x18, 0x60, 0x4, 0x90, 0x2, 0x20, 0xc4, 0x31, 0xe8, 0x92, 0x78, 0x4, 0x19, 0x44, 0x9, 0x85, 0x25, 0x8, 0x85, 0x24, 0x4, 0x48, 0x5, 0x18, 0xb4, 0x49, 0xb0, 0xd4, 0x21, 0xcc, 0x0, 0x4c, 0x1, 0x8c, 0x24, 0x4, 0x88, 0x76, 0xa8, 0xdc, 0x30, 0x4e, 0x0, 0xc, 0x44, 0xc, 0x0, 0xc, 0x8, 0x62, 0x0, 0xa, 0x0, 0x10, 0x10, 0x80, 0xc0, 0xe0, 0xe2, 0x8, 0x0, 0x13, 0x6, 0x9, 0x27, 0xc, 0xc9, 0x18, 0x2, 0x58, 0x0, 0x9a, 0x10, 0x44, 0x18, 0x81, 0x18, 0x47, 0x18, 0x7, 0x4d, 0x2, 0x10, 0x4, 0x1, 0x13, 0x4, 0xb, 0x26, 0xd, 0xc8, 0x18, 0x82, 0x18, 0x10, 0x54, 0x1, 0x18, 0x40, 0x1a, 0x88, 0x98, 0xc6, 0x9, 0x27, 0xa, 0x85, 0x11, 0x40, 0xa, 0x20, 0x8b, 0x16, 0x49, 0x8e, 0x19, 0x80, 0x58, 0x82, 0x18, 0x81, 0x58, 0x8, 0x90, 0x89, 0xd8, 0x5, 0xe, 0x29, 0x7, 0x12, 0x4, 0x0, 0x68, 0x1, 0x87, 0xa0, 0xf, 0xa, 0xd4, 0x9, 0x18, 0x40, 0x1a, 0x0, 0x58, 0x2, 0x18, 0x40, 0x99, 0x4, 0x58, 0x6, 0x2b, 0x4, 0x13, 0x5, 0x80, 0x8, 0x2, 0x0, 0x28, 0x2, 0x40, 0x8, 0x97, 0xa, 0x9d, 0x53, 0x4, 0x0, 0x51, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}

func main() {
	raspiAdaptor := raspi.NewAdaptor()
	oled := spi.NewSSD1306Driver(raspiAdaptor)
	work := func() {
		oled.Clear()
		oled.SetBufferAndDisplay(gobotLogo)
	}
	robot := gobot.NewRobot("ssd1360",
		[]gobot.Connection{raspiAdaptor},
		[]gobot.Device{oled},
		work,
	)
	robot.Start()
}
