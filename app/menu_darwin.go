package app

import (
	"github.com/kbinani/robot/ax"
)

type Menu struct {
	axMenu *ax.UIElement
}

func newMenu(menu *ax.UIElement) *Menu {
	m := new(Menu)
	m.axMenu = menu
	return m
}

func (menu *Menu) items() []*MenuItem {
	items := []*MenuItem{}
	if menu.axMenu == nil {
		return items
	}
	children := menu.axMenu.Children()
	for _, child := range children {
		items = append(items, newMenuItem(child))
	}
	return items
}
