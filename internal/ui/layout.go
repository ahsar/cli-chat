package ui

import (
	"bytes"

	"log"

	"github.com/ahsar/cli-chat/internal/ui/components/contacts"
	"github.com/ahsar/cli-chat/internal/ui/constant"
	"github.com/eatmoreapple/openwechat"

	"github.com/ahsar/cli-chat/internal/ui/components/message"
	"github.com/ahsar/cli-chat/internal/ui/components/rencent"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keymap struct {
	next, quit, enter, send key.Binding
}

type model struct {
	rencent  rencent.Rencent
	contacts contacts.Model
	message  *message.Model
	width    int
	height   int
	keymap   keymap
	help     help.Model
	msgch    chan *openwechat.Message
}

func NewModel(ch chan *openwechat.Message) (m model) {
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
				key.WithHelp("enter", "选择联系人"),
			),
			send: key.NewBinding(
				key.WithKeys("ctrl+s"),
				key.WithHelp("ctrl+s", "发送消息"),
			),
		},
		rencent:  rencent.NewList(),
		message:  message.NewModel(),
		contacts: contacts.NewModel(),
		msgch:    ch,
	}

	m.contacts.Focus()
	m.SetContacts()
	go m.consumer()

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
			m.exit()
			return m, tea.Quit
		case key.Matches(msg, m.keymap.next):
			m.focusInTurn()
			//case key.Matches(msg, m.keymap.enter):
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
			key.WithHelp("通讯录", "enter 选择联系人"))

	case i & constant.RencentPanel:
		_, cmd = m.rencent.Update(msg)
		cmds = append(cmds, cmd)
		m.keymap.enter = key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("最近联系的人", "enter 选择联系人"))

	case i & constant.DialogPanel:
		_, cmd = m.message.Update(msg)
		cmds = append(cmds, cmd)
		m.keymap.enter = key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("输入框", "ctrl+enter 发送消息"))
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	helpinfo := m.help.ShortHelpView([]key.Binding{
		m.keymap.next,
		m.keymap.quit,
		m.keymap.enter,
		m.keymap.send,
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

func (m *model) consumer() {
	for v := range m.msgch {
		m.onMsg(v)
	}
}
