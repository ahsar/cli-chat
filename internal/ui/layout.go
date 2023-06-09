package ui

import (
	"bytes"

	"github.com/ahsar/cli-chat/internal/ui/components/contacts"
	"github.com/ahsar/cli-chat/internal/ui/components/dialog"
	"github.com/ahsar/cli-chat/internal/ui/components/message"
	"github.com/ahsar/cli-chat/internal/ui/components/rencent"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	// 帮助栏高度
	helpHeight = 1
)

type keymap struct {
	next, prev, quit key.Binding
}

type model struct {
	rencent  rencent.Model
	contacts contacts.Model
	dialog   dialog.Model
	message  message.Model
	width    int
	height   int
	keymap   keymap
	help     help.Model
	focus    int
}

func NewModel() (m model) {
	m = model{
		help: help.New(),
		keymap: keymap{
			next: key.NewBinding(
				key.WithKeys("tab"),
				key.WithHelp("tab", "next"),
			),
			prev: key.NewBinding(
				key.WithKeys("shift+tab"),
				key.WithHelp("shift+tab", "prev"),
			),
			quit: key.NewBinding(
				key.WithKeys("esc", "ctrl+c"),
				key.WithHelp("esc", "quit"),
			),
		},
		focus: 3, // 默认通讯录高亮
	}

	//m.updateKeybindings()
	return
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			//todo blur
			return m, tea.Quit
		case key.Matches(msg, m.keymap.next):
			// todo focus
			m.contacts.Focus()
		case key.Matches(msg, m.keymap.prev):
			//m.inputs[m.focus].Blur()
			//m.focus--
			//if m.focus < 0 {
			//m.focus = len(m.inputs) - 1
			//}
			//cmd := m.inputs[m.focus].Focus()
			//cmds = append(cmds, cmd)
			//case key.Matches(msg, m.keymap.add):
			//m.inputs = append(m.inputs, newTextarea())

			//case key.Matches(msg, m.keymap.remove):
			//m.inputs = m.inputs[:len(m.inputs)-1]
			//if m.focus > len(m.inputs)-1 {
			//m.focus = len(m.inputs) - 1
			//}
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}

	//m.updateKeybindings()

	// Update all textareas
	//for i := range m.inputs {
	//newModel, cmd := m.inputs[i].Update(msg)
	//m.inputs[i] = newModel
	//cmds = append(cmds, cmd)
	//}

	// TODO update helpHeight
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	help := m.help.ShortHelpView([]key.Binding{
		m.keymap.next,
		m.keymap.prev,
		m.keymap.quit,
	})

	// Page
	var buff bytes.Buffer
	buff.WriteString(
		lipgloss.JoinHorizontal(
			lipgloss.Right,
			m.rencent.View(),
			m.message.View(),
			m.contacts.View(),
		))
	buff.WriteString("\n\n" + help)

	// Render
	lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Align(lipgloss.Center).
		Background(lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}).
		Height(20).
		//Border(lipgloss.NormalBorder()).
		//AlignHorizontal(lipgloss.Center).
		Render(buff.String())

	return buff.String()
}
