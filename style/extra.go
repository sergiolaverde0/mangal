package style

import (
	"github.com/sergiolaverde0/mangal/color"
	"github.com/charmbracelet/lipgloss"
)

var (
	Title      = NewColored(color.Purple, color.HiBg).Padding(0, 1).Render
	ErrorTitle = NewColored(color.Red, color.HiBg).Padding(0, 1).Render
)

func Tag(foreground, background lipgloss.Color) func(...string) string {
	return NewColored(foreground, background).Padding(0, 1).Render
}
