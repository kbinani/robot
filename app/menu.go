package app

// Items returns sub items of the menu.
func (menu *Menu) Items() []*MenuItem {
	return menu.items()
}
