package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

var active = 0

func setView(g *gocui.Gui) (err error) {
	var v *gocui.View

	// 最近联系的人
	if v, err = g.SetView(wx_view, 0, 0, MaxX/5-1, MaxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "recent"
		v.Editable = true
		v.Autoscroll = true
	}

	// 对话框
	if v, err = g.SetView(msg_view, MaxX/5, 0, MaxX*2/3-1, MaxY/2+3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Message(nickname%s)"
		v.Autoscroll = true
		//v.Editable = true
	}

	// 聊天输入框
	if v, err = g.SetView(talk_view, MaxX/5, MaxY/2+4, MaxX*2/3-1, MaxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Msg"
		v.Editable = true
		if _, err := setCurrentViewOnTop(g, talk_view); err != nil {
			fmt.Println("set cur err", err)
			return err
		}
	}

	return
}

func (g *OutPut) SetContacts(friends string) (err error) {
	// 通讯录
	if v, err := g.G.SetView(contacts_view, MaxX*2/3, 0, MaxX-6, MaxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.BgColor = 129
		v.Title = "contacts"
		v.Wrap = true
		setCurrentViewOnTop(g.G, contacts_view)

		fmt.Fprintln(v, friends)
	}
	return
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]

	//out, _ := g.View(msg_view)
	//fmt.Fprintf(out, "gui pointer %p ", g)
	//fmt.Fprintln(out, "from", active, v.Name(), "to", nextIndex, name)
	if _, err := setCurrentViewOnTop(g, name); err != nil {
		fmt.Println("set cur err", err)
		return err
	}

	active = nextIndex
	return nil
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func selectUser(g *gocui.Gui, v *gocui.View) (err error) {
	_, row := v.Cursor()

	var l string
	if l, err = v.Line(row); err != nil {
		l = ""
	}

	// todo show recent
	//v.Clear()
	out, err := g.View(msg_view)
	out.Clear()
	fmt.Fprintf(out, "#TODO %s 聊天记录\n", l)

	setCurrentViewOnTop(g, talk_view)
	return
}
