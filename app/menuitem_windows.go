package app

import (
	"github.com/kbinani/win"
	"unsafe"
	// "fmt"
	u "github.com/kbinani/robot/unsafe"
	// "image"
)

type MenuItem struct {
	hMenu       win.HMENU
	index       uint32
	hWnd        win.HWND
	hMenuParent win.HMENU
}

func newMenuItem(hMenu win.HMENU, index uint32, hWnd win.HWND, hMenuParent win.HMENU) *MenuItem {
	m := new(MenuItem)
	m.hMenu = hMenu
	m.index = index
	m.hWnd = hWnd
	m.hMenuParent = hMenuParent
	return m
}

func (item *MenuItem) text() string {
	var info win.MENUITEMINFO
	info.FMask = win.MIIM_STRING
	info.FType = win.MFT_STRING
	info.CbSize = uint32(unsafe.Sizeof(info))

	if !win.GetMenuItemInfo(item.hMenu, item.index, win.TRUE, &info) {
		return ""
	}

	info.Cch++
	storage := make([]uint16, info.Cch)
	info.DwTypeData = (*uint16)(u.Ptr(storage))

	if !win.GetMenuItemInfo(item.hMenu, item.index, win.TRUE, &info) {
		return ""
	}

	return win.UTF16PtrToString((*uint16)(u.Ptr(storage)))
}

func (item *MenuItem) click() {
	if !item.isEnabled() {
		return
	}
	id := win.GetMenuItemID(item.hMenu, int32(item.index))
	win.PostMessage(item.hWnd, win.WM_COMMAND, uintptr(id), 0)
}

func (item *MenuItem) sub() *Menu {
	var info win.MENUITEMINFO
	info.FMask = win.MIIM_SUBMENU
	info.CbSize = uint32(unsafe.Sizeof(info))

	if !win.GetMenuItemInfo(item.hMenu, item.index, win.TRUE, &info) {
		return nil
	}

	if info.HSubMenu == win.HMENU(0) {
		return nil
	}

	return newMenu(info.HSubMenu, item.hWnd, item.hMenu)
}

func (item *MenuItem) isEnabled() bool {
	state := item.state()
	disabled := (state & win.MF_GRAYED) == win.MF_GRAYED
	return !disabled
}

func (item *MenuItem) isSelected() bool {
	state := item.state()
	checked := ((state & win.MFS_CHECKED) == win.MFS_CHECKED)
	return checked
}

func (item *MenuItem) state() uint32 {
	var info win.MENUITEMINFO
	info.FMask = win.MIIM_STATE
	info.CbSize = uint32(unsafe.Sizeof(info))
	if !win.GetMenuItemInfo(item.hMenu, item.index, win.TRUE, &info) {
		return 0
	}
	return info.FState
}
