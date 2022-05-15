package reader

import (
	"github.com/fatih/color"
)

var colorMap = map[int64]color.Attribute{
	1:  color.FgRed,
	2:  color.FgGreen,
	3:  color.FgYellow,
	4:  color.FgBlue,
	5:  color.FgMagenta,
	6:  color.FgCyan,
	7:  color.FgHiRed,
	8:  color.FgHiGreen,
	9:  color.FgHiYellow,
	10: color.FgHiBlue,
	11: color.FgHiMagenta,
	12: color.FgHiCyan,
}
