package colors

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	AnsiReset      = "\033[0m"
	Black          = "#000000"
	White          = "#FFFFFF"
	Gray           = "#808080"
	Silver         = "#C0C0C0"
	DimGray        = "#696969"
	DarkGray       = "#A9A9A9"
	LightGray      = "#D3D3D3"
	Gainsboro      = "#DCDCDC"
	WhiteSmoke     = "#F5F5F5"
	Red            = "#FF0000"
	Crimson        = "#DC143C"
	FireBrick      = "#B22222"
	DarkRed        = "#8B0000"
	Maroon         = "#800000"
	Pink           = "#FFC0CB"
	LightPink      = "#FFB6C1"
	HotPink        = "#FF69B4"
	DeepPink       = "#FF1493"
	Magenta        = "#FF00FF"
	Orange         = "#FFA500"
	DarkOrange     = "#FF8C00"
	Coral          = "#FF7F50"
	Tomato         = "#FF6347"
	OrangeRed      = "#FF4500"
	Yellow         = "#FFFF00"
	Gold           = "#FFD700"
	LightYellow    = "#FFFFE0"
	LemonChiffon   = "#FFFACD"
	Moccasin       = "#FFE4B5"
	Green          = "#00FF00"
	Lime           = "#00FF00"
	ForestGreen    = "#228B22"
	DarkGreen      = "#006400"
	Olive          = "#808000"
	OliveDrab      = "#6B8E23"
	YellowGreen    = "#9ACD32"
	LimeGreen      = "#32CD32"
	SpringGreen    = "#00FF7F"
	SeaGreen       = "#2E8B57"
	Blue           = "#0000FF"
	MediumBlue     = "#0000CD"
	DarkBlue       = "#00008B"
	Navy           = "#000080"
	MidnightBlue   = "#191970"
	Cyan           = "#00FFFF"
	Aqua           = "#00FFFF"
	Teal           = "#008080"
	Turquoise      = "#40E0D0"
	DeepSkyBlue    = "#00BFFF"
	SkyBlue        = "#87CEEB"
	LightSteelBlue = "#B0C4DE"
	Purple         = "#800080"
	Violet         = "#EE82EE"
	Indigo         = "#4B0082"
	DarkMagenta    = "#8B008B"
	DarkViolet     = "#9400D3"
	Amethyst       = "#9966CC"
	Lavender       = "#E6E6FA"
	Plum           = "#DDA0DD"
	Orchid         = "#DA70D6"
	Brown          = "#A52A2A"
	SaddleBrown    = "#8B4513"
	Sienna         = "#A0522D"
	Chocolate      = "#D2691E"
	Peru           = "#CD853F"
	SandyBrown     = "#F4A460"
	BurlyWood      = "#DEB887"
	Tan            = "#D2B48C"
	Wheat          = "#F5DEB3"
	Beige          = "#F5F5DC"
)

var basicAnsiMap = map[string]int{
	"#000000": 30, "#FF0000": 31, "#00FF00": 32, "#FFFF00": 33,
	"#0000FF": 34, "#FF00FF": 35, "#00FFFF": 36, "#FFFFFF": 37,
	"#808080": 90, "#DC143C": 31, "#B22222": 31, "#8B0000": 31,
	"#800000": 31, "#FFC0CB": 91, "#FFB6C1": 91, "#FF69B4": 91,
	"#FF1493": 91, "#FFA500": 33, "#FF8C00": 33, "#FF7F50": 91,
	"#FF6347": 31, "#FF4500": 31, "#FFD700": 33, "#FFFFE0": 93,
	"#FFFACD": 93, "#FFE4B5": 93, "#228B22": 32, "#006400": 32,
	"#808000": 33, "#6B8E23": 32, "#9ACD32": 92, "#32CD32": 92,
	"#00FF7F": 92, "#2E8B57": 32, "#0000CD": 34, "#00008B": 34,
	"#000080": 34, "#191970": 34, "#008080": 36, "#40E0D0": 96,
	"#00BFFF": 96, "#87CEEB": 96, "#B0C4DE": 37, "#800080": 35,
	"#EE82EE": 95, "#4B0082": 35, "#8B008B": 35, "#9400D3": 35,
	"#9966CC": 95, "#E6E6FA": 37, "#DDA0DD": 95, "#DA70D6": 95,
	"#A52A2A": 31, "#8B4513": 31, "#A0522D": 31, "#D2691E": 33,
	"#CD853F": 33, "#F4A460": 93, "#DEB887": 37, "#D2B48C": 37,
	"#F5DEB3": 37, "#F5F5DC": 37, "#C0C0C0": 37, "#696969": 90,
	"#A9A9A9": 37, "#D3D3D3": 37, "#DCDCDC": 37, "#F5F5F5": 37,
}

func SupportsTrueColor() bool {
	term := os.Getenv("TERM")
	colorterm := os.Getenv("COLORTERM")
	if colorterm == "truecolor" || colorterm == "24bit" {
		return true
	}
	if strings.Contains(term, "truecolor") || strings.Contains(term, "24bit") {
		return true
	}
	return false
}

func hexToAnsi8(hex string) string {
	hex = "#" + strings.ToUpper(strings.TrimPrefix(strings.TrimSpace(hex), "#"))
	if code, found := basicAnsiMap[hex]; found {
		return fmt.Sprintf("\033[%dm", code)
	}

	r, g, b := ToRGB(hex)
	r8 := int(math.Round(float64(r) / 255.0 * 5.0))
	g8 := int(math.Round(float64(g) / 255.0 * 5.0))
	b8 := int(math.Round(float64(b) / 255.0 * 5.0))
	ansiCode := 16 + (36 * r8) + (6 * g8) + b8

	return fmt.Sprintf("\033[38;5;%dm", ansiCode)
}

func ToRGB(color string) (r, g, b int) {
	hex := strings.TrimPrefix(strings.TrimSpace(color), "#")
	if len(hex) != 6 {
		return 0, 0, 0
	}
	rVal, _ := strconv.ParseInt(hex[0:2], 16, 32)
	gVal, _ := strconv.ParseInt(hex[2:4], 16, 32)
	bVal, _ := strconv.ParseInt(hex[4:6], 16, 32)
	return int(rVal), int(gVal), int(bVal)
}

func ToRGBA(color string, alpha float64) (r, g, b, a uint8) {
	rInt, gInt, bInt := ToRGB(color)
	if alpha < 0 {
		alpha = 0
	} else if alpha > 1 {
		alpha = 1
	}
	alphaUint8 := uint8(alpha * 255)
	return uint8(rInt), uint8(gInt), uint8(bInt), alphaUint8
}

func HexToAnsi(hex string) string {
	hex = strings.TrimPrefix(strings.TrimSpace(hex), "#")
	if len(hex) != 6 {
		return AnsiReset
	}

	if !SupportsTrueColor() {
		return hexToAnsi8(hex)
	}

	r, _ := strconv.ParseInt(hex[0:2], 16, 64)
	g, _ := strconv.ParseInt(hex[2:4], 16, 64)
	b, _ := strconv.ParseInt(hex[4:6], 16, 64)
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func ColorPrint(text string, color string) {
	if !strings.Contains(strings.TrimSpace(color), "AnsiReset") {
		color = HexToAnsi(color)
	}
	fmt.Printf("%s%s%s", color, text, AnsiReset)
}

func ColorPrintln(text string, color string) {
	if strings.HasPrefix(strings.TrimSpace(color), "#") {
		color = HexToAnsi(color)
	}
	fmt.Printf("%s%s%s\n", color, text, AnsiReset)
}
