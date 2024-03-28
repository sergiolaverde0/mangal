package color

import "github.com/charmbracelet/lipgloss"

var (
    Red    = lipgloss.Color("1")
    Green  = lipgloss.Color("2")
    Yellow = lipgloss.Color("3")
    Blue   = lipgloss.Color("4")
    Purple = lipgloss.Color("5")
    Cyan   = lipgloss.Color("6")
)

var (
    HiBg   = lipgloss.AdaptiveColor{Light: "7", Dark: "0"}
    HiFg   = lipgloss.AdaptiveColor{Light: "10", Dark: "14"}
    Fg     = lipgloss.AdaptiveColor{Light: "11", Dark: "12"}
    Bg     = lipgloss.AdaptiveColor{Light: "15", Dark: "8"}
)
