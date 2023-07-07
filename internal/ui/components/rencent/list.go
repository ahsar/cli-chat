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

type (
	item struct {
		// id means local id, not vxid
		id, msg, uname string
	}

	keymap struct {
		enter key.Binding
	}

	listModel struct {
		items   []item
		list    list.Model
		keymap  keymap
		focused byte
		uMap    map[string]int // 记录local user id in rencent list index
	}
)

func (i item) Title() string       { return i.uname }
func (i item) Description() string { return i.msg }
func (i item) FilterValue() string { return i.msg }

func NewList() (m *listModel) {
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
// id means local id
func (m *listModel) AddUser(id, msg string) {
	if v, ok := m.uMap[id]; ok {
		newest := m.items[v]
		l := len(m.items)

		if l-1 != v {
			for i := v; i < l-1; i++ {
				m.items[i] = m.items[i+1]
				m.uMap[m.items[i].id] = i + 1
			}
		}
		newest.msg = msg
		m.items[l-1] = newest
		m.uMap[id] = l - 1
	} else {
		i, _ := strconv.Atoi(id)
		n := chat.GetName(chat.FriendById(i).User)
		if n == "" {
			return
		}

		m.addItem(item{
			id:    id,
			uname: n,
			msg:   msg,
		})
		return
	}

	m.render(m.items)
}

// SetItems todo
//
// 初始化最近联系人
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

// 新增元素
// 逆序插入, 减少数组头部插入, 元素移动
func (m *listModel) addItem(i item) {
	m.items = append(m.items, i)
	m.uMap[i.id] = len(m.items) - 1
	m.list.InsertItem(0, i)
}

func (m *listModel) render(i []item) {
	items := make([]list.Item, 0, len(i))
	for i := len(m.items) - 1; i >= 0; i-- {
		items = append(items, m.items[i])
	}

	m.list.SetItems(items)
}

// SetUserMsg
// 设定用户最后聊天文字
// id means vxid
func (m *listModel) SetUserMsg(id, msg string) {
	i := chat.ConverVxid2Id(id)
	if i < 0 {
		return
	}

	m.AddUser(strconv.Itoa(i), msg)
}
