package ui

import (
	//"bytes"
	//"fmt"

	"fmt"
	//"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	// 面板数量
	panelsNum = 4

	// 帮助栏高度
	helpHeight = 1
)

type keymap struct {
	next, prev, quit key.Binding
}

type model struct {
	width  int
	height int
	keymap keymap
	help   help.Model
	inputs []textarea.Model
	focus  int
}

func NewModel() (m model) {
	m = model{
		inputs: make([]textarea.Model, panelsNum),
		help:   help.New(),
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

	for i := 0; i < panelsNum; i++ {
		m.inputs[i] = newTextarea(i)
	}

	m.inputs[m.focus].Focus()
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
			for i := range m.inputs {
				m.inputs[i].Blur()
			}
			return m, tea.Quit
		case key.Matches(msg, m.keymap.next):
			m.inputs[m.focus].Blur()
			m.focus++
			if m.focus > len(m.inputs)-1 {
				m.focus = 0
			}
			cmd := m.inputs[m.focus].Focus()
			cmds = append(cmds, cmd)
		case key.Matches(msg, m.keymap.prev):
			m.inputs[m.focus].Blur()
			m.focus--
			if m.focus < 0 {
				m.focus = len(m.inputs) - 1
			}
			cmd := m.inputs[m.focus].Focus()
			cmds = append(cmds, cmd)
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
	m.sizeInput()

	// Update all textareas
	for i := range m.inputs {
		newModel, cmd := m.inputs[i].Update(msg)
		m.inputs[i] = newModel
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() (s string) {
	help := m.help.ShortHelpView([]key.Binding{
		m.keymap.next,
		m.keymap.prev,
		m.keymap.quit,
	})

	//var buff bytes.Buffer
	var views = make([]string, 0, 2)
	//for i := range m.inputs {
	//views = append(views, m.inputs[i].View())

	//switch i {
	//case 0:
	//fmt.Println(1)
	//buff.WriteString(lipgloss.JoinHorizontal(lipgloss.Bottom, m.inputs[i].View()))
	//case 1:
	//fmt.Println(2)
	//buff.WriteString(lipgloss.JoinHorizontal(lipgloss.Bottom, m.inputs[i].View()))
	//case 2:
	//fmt.Println(3)
	//buff.WriteString(lipgloss.JoinVertical(lipgloss.Bottom, m.inputs[i].View()))
	//case 3:
	//fmt.Println(4)
	//buff.WriteString(lipgloss.JoinHorizontal(lipgloss.Right, m.inputs[i].View()))
	//}
	//}

	//lip
	//fmt.Println(buff.String())
	//os.Exit(1)
	//buff.WriteString("\n\n" + help)

	views = append(views, m.inputs[0].View(), m.inputs[1].View())
	s = lipgloss.JoinHorizontal(lipgloss.Top, views...) + "\n\n" + help
	//lipgloss.JoinVertical(lipgloss.Right, []string{m.inputs[2].View()}...) + "\n\n" +
	//lipgloss.JoinHorizontal(lipgloss.Right, []string{m.inputs[3].View()}...) + "\n\n" + help
	fmt.Println(s)
	//os.Exit(1)
	return s
	//return buff.String()
}
