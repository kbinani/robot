package app

import (
	"github.com/kbinani/win"
	// "fmt"
)

type Menu struct {
	hMenu       win.HMENU
	hWnd        win.HWND
	hMenuParent win.HMENU
}

func newMenu(hMenu win.HMENU, hWnd win.HWND, hMenuParent win.HMENU) *Menu {
	m := new(Menu)
	m.hMenu = hMenu
	m.hMenuParent = hMenuParent
	m.hWnd = hWnd
	return m
}

func (menu *Menu) items() []*MenuItem {
	items := []*MenuItem{}
	num := int(win.GetMenuItemCount(menu.hMenu))
	for i := 0; i < num; i++ {
		items = append(items, newMenuItem(menu.hMenu, uint32(i), menu.hWnd, menu.hMenuParent))
	}
	return items
}
