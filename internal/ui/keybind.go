package ui

import (
	"github.com/ahsar/cli-chat/internal/logger"
	"github.com/jroimartin/gocui"
)

func keybinds(g *gocui.Gui) {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		logger.Panic("keybind err %+v", err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlW, gocui.ModNone, nextView); err != nil {
		logger.Panic("keybind err %+v", err)
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		logger.Panic("keybind err %+v", err)
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		logger.Panic("keybind err %+v", err)
	}
	if err := g.SetKeybinding(contacts_view, gocui.KeyEnter, gocui.ModNone, selectUser); err != nil {
		logger.Panic("keybind err %+v", err)
	}
}
