package colors

import (
	"fmt"
	"strconv"
	"strings"
)

const (
  AnsiReset  = "\033[0m"
  Black = "#000000"
  White = "#FFFFFF"
  Gray = "#808080"
  Silver = "#C0C0C0"
  DimGray = "#696969"
  DarkGray = "#A9A9A9"
  LightGray = "#D3D3D3"
  Gainsboro = "#DCDCDC"
  WhiteSmoke = "#F5F5F5"
  Red = "#FF0000"
  Crimson = "#DC143C"
  FireBrick = "#B22222"
  DarkRed = "#8B0000"
  Maroon = "#800000"
  Pink = "#FFC0CB"
  LightPink = "#FFB6C1"
  HotPink = "#FF69B4"
  DeepPink = "#FF1493"
  Magenta = "#FF00FF"
  Orange = "#FFA500"
  DarkOrange = "#FF8C00"
  Coral = "#FF7F50"
  Tomato = "#FF6347"
  OrangeRed = "#FF4500"
  Yellow = "#FFFF00"
  Gold = "#FFD700"
  LightYellow = "#FFFFE0"
  LemonChiffon = "#FFFACD"
  Moccasin = "#FFE4B5"
  Green = "#00FF00"
  Lime = "#00FF00"
  ForestGreen = "#228B22"
  DarkGreen = "#006400"
  Olive = "#808000"
  OliveDrab = "#6B8E23"
  YellowGreen = "#9ACD32"
  LimeGreen = "#32CD32"
  SpringGreen = "#00FF7F"
  SeaGreen = "#2E8B57"
  Blue = "#0000FF"
  MediumBlue = "#0000CD"
  DarkBlue = "#00008B"
  Navy = "#000080"
  MidnightBlue = "#191970"
  Cyan = "#00FFFF"
  Aqua = "#00FFFF"
  Teal = "#008080"
  Turquoise = "#40E0D0"
  DeepSkyBlue = "#00BFFF"
  SkyBlue = "#87CEEB"
  LightSteelBlue = "#B0C4DE"
  Purple = "#800080"
  Violet = "#EE82EE"
  Indigo = "#4B0082"
  DarkMagenta = "#8B008B"
  DarkViolet = "#9400D3"
  Amethyst = "#9966CC"
  Lavender = "#E6E6FA"
  Plum = "#DDA0DD"
  Orchid = "#DA70D6"
  Brown = "#A52A2A"
  SaddleBrown = "#8B4513"
  Sienna = "#A0522D"
  Chocolate = "#D2691E"
  Peru = "#CD853F"
  SandyBrown = "#F4A460"
  BurlyWood = "#DEB887"
  Tan = "#D2B48C"
  Wheat = "#F5DEB3"
  Beige = "#F5F5DC"
)

func hexToAnsi(hex string) string {
	hex = strings.TrimPrefix(strings.TrimSpace(hex), "#")

	if len(hex) != 6 {
		return Reset
	}

	r, _ := strconv.ParseInt(hex[0:2], 16, 64)
	g, _ := strconv.ParseInt(hex[2:4], 16, 64)
	b, _ := strconv.ParseInt(hex[4:6], 16, 64)

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func ColorPrint(text string, color string) {
	if !strings.Contains(strings.TrimSpace(color), "AnsiReset") {
		color = hexToAnsi(color)
	}

	fmt.Printf("%s%s%s", color, text, AnsiReset)
}

func ColorPrintln(text string, color string) {
	if strings.HasPrefix(strings.TrimSpace(color), "#") {
		color = hexToAnsi(color)
	}

	fmt.Printf("%s%s%s\n", color, text, Reset)
}

