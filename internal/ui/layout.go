package ui

import (
	"bytes"
	"log"

	"github.com/ahsar/cli-chat/internal/ui/components/contacts"
	"github.com/ahsar/cli-chat/internal/ui/components/dialog"
	"github.com/ahsar/cli-chat/internal/ui/components/message"
	"github.com/ahsar/cli-chat/internal/ui/components/rencent"
	"github.com/ahsar/cli-chat/internal/ui/constant"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keymap struct {
	next, prev, quit, enter key.Binding
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
	current  int
}

func NewModel() (m model) {
	log.Println("new ui model")

	cts := contacts.NewModel()
	cts.SetRow([]table.Row{
		{"1", "xj"},
		{"2", "wll"},
		{"3", "tlt"},
	})

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
			enter: key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "send or select"),
			),
		},
		rencent:  rencent.NewModel(),
		message:  message.NewModel(),
		contacts: cts,
		current:  1, // 默认通讯录高亮
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
			m.blur()
			return m, tea.Quit
		case key.Matches(msg, m.keymap.next):
			m.focus()
		case key.Matches(msg, m.keymap.prev):
			//todo
		case key.Matches(msg, m.keymap.enter):
			log.Println("ui layout enter")
			//todo
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}
	m.sizeInput()

	//m.updateKeybindings()

	var cmd tea.Cmd
	if m.current == constant.ContactPanel {
		_, cmd := m.contacts.Update(msg)
		cmds = append(cmds, cmd)

		// 重新定义enter键
		m.keymap.enter = key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "enter 选择联系人"),
		)
	} else {
		_, cmd = m.rencent.Update(msg)
		cmds = append(cmds, cmd)

		m.keymap.enter = key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "enter 发送消息"),
		)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	help := m.help.ShortHelpView([]key.Binding{
		m.keymap.next,
		m.keymap.prev,
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
