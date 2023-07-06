package rencent

import (
	"strconv"

	"github.com/ahsar/cli-chat/internal/chat"
	"github.com/ahsar/cli-chat/internal/ui/components/message"
	"github.com/ahsar/cli-chat/internal/ui/constant"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	// id means local id, not vxid
	id, msg, uname string
}

type keymap struct {
	enter key.Binding
}

func (i item) Title() string       { return i.uname }
func (i item) Description() string { return i.msg }
func (i item) FilterValue() string { return i.msg }

type listModel struct {
	items   []item
	list    list.Model
	keymap  keymap
	focused byte
	uMap    map[string]int // uid:index
}

func NewList() (m *listModel) {
	// TODO read from storage file
	m = &listModel{
		list: list.New([]list.Item{}, list.NewDefaultDelegate(), 170, 40),
		keymap: keymap{
			enter: key.NewBinding(
				key.WithKeys("enter"),
			),
		},
		uMap: make(map[string]int),
	}

	m.list.Title = "rencent contacts                                         "
	//m.list.SetShowTitle(false)
	m.list.SetShowFilter(false)
	m.list.SetShowPagination(false)
	m.list.SetShowHelp(false)
	m.list.SetShowStatusBar(false)
	Obj = m
	return
}

func (m listModel) Init() tea.Cmd {
	return nil
}

func (m *listModel) Update(msg tea.Msg) (_ tea.Model, cmd tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.enter):
			var (
				it list.Item
				i  item
				ok bool
			)
			if it = m.list.SelectedItem(); it == nil {
				return
			}

			if i, ok = it.(item); !ok {
				return
			}

			// 1. register message user
			message.Msg.SetUser(i.id)

			// 2. blur current panel
			m.Blur()
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *listModel) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.list.View(),
	)
}

func (m *listModel) SetSize(w, h int) {
	m.list.SetSize(w, h)
}

func (m *listModel) Focus() {
	m.focused = constant.RencentPanel
}

func (m *listModel) Blur() {
	m.focused = 0
}

func (m *listModel) Focused() byte {
	return m.focused
}

// 添加联系人
// 不存在: add
// 存在: set to top
func (m *listModel) AddUser(id string) {
	if v, ok := m.uMap[id]; ok {
		// 索引交换
		m.uMap[id] = 0
		m.uMap[m.items[0].id] = v

		// 数组交换
		m.items[0], m.items[v] = m.items[v], m.items[0]
	} else {
		i, _ := strconv.Atoi(id)
		n := chat.GetName(chat.FriendById(i).User)
		if n == "" {
			return
		}

		m.addItem(item{
			id:    id,
			uname: n,
		})
		return
	}

	m.render(m.items)
}

// todo
func (m *listModel) SetItems(i any) {
	x, ok := i.([]list.Item)
	if !ok {
		return
	}

	if len(x) <= 0 {
		return
	}

	// todo
	// make index
	// make items
	m.list.SetItems(x)
}

func (m *listModel) addItem(i item) {
	m.items = append(m.items, i)
	m.list.InsertItem(0, i)
}

func (m *listModel) render(i []item) {
	items := make([]list.Item, 0, len(i))

	for _, v := range i {
		items = append(items, v)
	}

	m.list.SetItems(items)
}

// SetUserMsg
// 设定用户最后聊天文字
func (m *listModel) SetUserMsg(id, msg string) {
	var (
		v  int
		ok bool
	)
	if v, ok = m.uMap[id]; !ok {
		return
	}

	m.items[v].msg = msg
	m.render(m.items)
}
