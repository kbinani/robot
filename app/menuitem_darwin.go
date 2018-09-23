package app

import (
	"github.com/kbinani/robot/ax"
)

type MenuItem struct {
	axMenuItem *ax.UIElement
}

func newMenuItem(menu *ax.UIElement) *MenuItem {
	m := new(MenuItem)
	m.axMenuItem = menu
	return m
}

func (item *MenuItem) text() string {
	if item.axMenuItem == nil {
		return ""
	}
	return item.axMenuItem.Title()
}

func (item *MenuItem) click() {
	if item.axMenuItem != nil {
		item.axMenuItem.Perform(ax.PressAction)
	}
}

func (item *MenuItem) sub() *Menu {
	children := item.axMenuItem.Children()
	if len(children) != 1 {
		return nil
	}
	child := children[0]
	if child.Role() != ax.MenuRole {
		return nil
	}
	return newMenu(child)
}

func (item *MenuItem) isEnabled() bool {
	return item.axMenuItem.IsEnabled()
}

func (item *MenuItem) isSelected() bool {
	return item.axMenuItem.IsSelected()
}
