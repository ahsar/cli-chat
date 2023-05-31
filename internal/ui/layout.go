package ui

import (
	"github.com/ahsar/cli-chat/internal/logger"
	"github.com/jroimartin/gocui"
)

const (
	wx_view       = "wechat"
	msg_view      = "message"
	talk_view     = "dialog"
	contacts_view = "contacts"
)

var (
	viewArr = []string{
		wx_view,
		msg_view,
		talk_view,
		contacts_view}

	MaxX, MaxY int
)

type OutPut struct {
	G       *gocui.Gui
	Contact string
}

func New() *OutPut {
	ui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		logger.Panic("layout err %+v", err)
	}

	ui.Highlight = true
	ui.Cursor = true
	ui.SelFgColor = gocui.ColorGreen
	ui.SetManagerFunc(layout)

	MaxX, MaxY = ui.Size()
	keybinds(ui)
	setView(ui)

	//ui.SetCurrentView("contacts")
	return &OutPut{G: ui}
}

func (g *OutPut) Run() {
	if err := g.G.MainLoop(); err != nil && err != gocui.ErrQuit {
		logger.Panic("keybind err %+v", err)
	}
}

func layout(g *gocui.Gui) error {
	setView(g)
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
