package ui

import (
	"bytes"

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
	rencent  rencent.Model
	contacts contacts.Model
	dialog   dialog.Model
	message  message.Model
	width    int
	height   int
	keymap   keymap
	help     help.Model
	inputs   []textarea.Model
	focus    int
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

	//m.inputs[m.focus].Focus()
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
			//cmd := m.inputs[m.focus].Focus()
			//cmds = append(cmds, cmd)
		case key.Matches(msg, m.keymap.prev):
			m.inputs[m.focus].Blur()
			m.focus--
			if m.focus < 0 {
				m.focus = len(m.inputs) - 1
			}
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

	return m, tea.Batch(cmds...)
}

func getTable() (t table.Model) {
	columns := []table.Column{
		{Title: "Rank", Width: 4},
		{Title: "City", Width: 10},
		{Title: "Country", Width: 10},
		{Title: "Population", Width: 10},
	}

	rows := []table.Row{
		{"1", "Tokyo", "Japan", "37,274,000"},
		{"2", "Delhi", "India", "32,065,760"},
	}

	t = table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)
	return
}

func textA(i int) (t textarea.Model) {
	t = textarea.New()
	t.SetHeight(i)
	t.Prompt = ""
	t.SetValue("1212")
	//t.ShowLineNumbers = false
	t.Cursor.Style = constant.CursorStyle
	t.FocusedStyle.Placeholder = constant.FocusedPlaceholderStyle
	t.BlurredStyle.Placeholder = constant.PlaceholderStyle
	t.FocusedStyle.CursorLine = constant.CursorLineStyle
	t.FocusedStyle.Base = constant.FocusedBorderStyle
	t.BlurredStyle.Base = constant.BlurredBorderStyle
	t.FocusedStyle.EndOfBuffer = constant.EndOfBufferStyle
	t.BlurredStyle.EndOfBuffer = constant.EndOfBufferStyle
	t.KeyMap.LineNext = key.NewBinding(key.WithKeys("down"))
	t.KeyMap.LinePrevious = key.NewBinding(key.WithKeys("up"))
	t.Blur()
	return t
}

func (m model) View() string {
	help := m.help.ShortHelpView([]key.Binding{
		m.keymap.next,
		m.keymap.prev,
		m.keymap.quit,
	})

	var buff bytes.Buffer

	buff.WriteString(
		lipgloss.JoinHorizontal(
			lipgloss.Right,
			textA(15).View(),
			m.message.View(),
			//textA(5).View(),
			textA(15).View(),
		))
	buff.WriteString("\n\n" + help)
	//buff.WriteString(m.message.View())

	//buff.WriteString(
	lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Align(lipgloss.Center).
		//Background(lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}).
		Height(20).
		Border(lipgloss.NormalBorder()).
		//AlignHorizontal(lipgloss.Center).
		Render(buff.String())
		//m.rencent.View(),
		//"\n\n",
		//m.message.View(),
		////getTable().View(),
		////))

		////buff.WriteString(m.rencent.View())

	return buff.String()
}
