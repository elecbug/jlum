package color

import "fmt"

type Color struct {
	Red   byte `json:"red"`
	Green byte `json:"green"`
	Blue  byte `json:"blue"`
}

func (c Color) ToHex() string {
	return "#" + formatHex(c.Red) + formatHex(c.Green) + formatHex(c.Blue)
}

func formatHex(value byte) string {
	hex := "0123456789ABCDEF"
	return string(hex[value>>4]) + string(hex[value&0x0F])
}

func FromHex(hex string) Color {
	if len(hex) != 7 || hex[0] != '#' {
		return Color{0, 0, 0} // Return black if the format is incorrect
	}
	var r, g, b byte
	_, _ = fmt.Sscanf(hex[1:], "%02x%02x%02x", &r, &g, &b)
	return Color{Red: r, Green: g, Blue: b}
}
