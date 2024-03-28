package color

import "github.com/charmbracelet/lipgloss"

var (
    Red    = New("1")
    Green  = New("2")
    Yellow = New("3")
    Blue   = New("4")
    Purple = New("5")
    Cyan   = New("6")
    HiBg   = New("7")
    HiFg   = New("10")
    Fg     = New("11")
    Bg     = New("15")
)

func New(color string) lipgloss.Color {
	return lipgloss.Color(color)
}
