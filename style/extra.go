package style

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/metafates/mangal/color"
)

var (
	Title      = NewColored(color.Purple, color.HiBg).Padding(0, 1).Render
	ErrorTitle = NewColored(color.Red, color.HiBg).Padding(0, 1).Render
)

func Tag(foreground, background lipgloss.Color) func(...string) string {
	return NewColored(foreground, background).Padding(0, 1).Render
}
