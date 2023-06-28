package ui

import (
	"bytes"
	"log"

	"github.com/ahsar/cli-chat/internal/ui/components/contacts"
	"github.com/ahsar/cli-chat/internal/ui/constant"

	//"github.com/ahsar/cli-chat/internal/ui/components/dialog"
	"github.com/ahsar/cli-chat/internal/ui/components/message"
	"github.com/ahsar/cli-chat/internal/ui/components/rencent"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keymap struct {
	next, quit, enter key.Binding
}

type model struct {
	rencent  rencent.Model
	contacts contacts.Model
	message  *message.Model
	width    int
	height   int
	keymap   keymap
	help     help.Model
}

func NewModel() (m model) {
	log.Println("new ui model")

	m = model{
		help: help.New(),
		keymap: keymap{
			next: key.NewBinding(
				key.WithKeys("tab"),
				key.WithHelp("tab", "next"),
			),
			quit: key.NewBinding(
				key.WithKeys("esc", "ctrl+c"),
				key.WithHelp("esc", "quit"),
			),
			enter: key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "send or select"),
			),
		},
		rencent:  rencent.NewModel(),
		message:  message.NewModel(),
		contacts: contacts.NewModel(),
	}

	m.contacts.Focus()
	m.SetContacts()
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
			m.blur()
			return m, tea.Quit
		case key.Matches(msg, m.keymap.next):
			m.focusInTurn()
		case key.Matches(msg, m.keymap.enter):
			//log.Println("ui layout enter")
			//todo
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}
	m.sizeInput()

	var cmd tea.Cmd
	i := m.getCurrent()
	switch i {
	case i & constant.ContactPanel:
		_, cmd = m.contacts.Update(msg)
		cmds = append(cmds, cmd)
		m.keymap.enter = key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "enter 选择联系人"))

	case i & constant.RencentPanel:
		_, cmd = m.rencent.Update(msg)
		cmds = append(cmds, cmd)
		m.keymap.enter = key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "enter 发送消息"),
		)

	case i & constant.DialogPanel:
		_, cmd = m.message.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	helpinfo := m.help.ShortHelpView([]key.Binding{
		m.keymap.next,
		m.keymap.quit,
		m.keymap.enter,
	})

	// Panels
	var buff bytes.Buffer
	buff.WriteString(
		lipgloss.JoinHorizontal(
			lipgloss.Right,
			m.rencent.View(),
			m.message.View(),
			m.contacts.View(),
		))
	buff.WriteString("\n\n" + helpinfo)

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
