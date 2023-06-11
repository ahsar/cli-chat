package constant

import "github.com/charmbracelet/lipgloss"

const (
	HelpHeight = 1
)

var (
	CursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))

	CursorLineStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("57")).
			Foreground(lipgloss.Color("230"))

	PlaceholderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("238"))

	EndOfBufferStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("235"))

	FocusedPlaceholderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("99"))

	FocusedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.ThickBorder()).
				BorderForeground(lipgloss.Color("238"))

	BlurredBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder())
)
