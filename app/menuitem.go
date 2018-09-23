package app

// Text returns the label of the menu item.
func (item *MenuItem) Text() string {
	return item.text()
}

// Click menu item.
func (item *MenuItem) Click() {
	item.click()
}

// Sub returns sub menu. If menu item does not have sub item, returns nil.
func (item *MenuItem) Sub() *Menu {
	return item.sub()
}

func (item *MenuItem) IsEnabled() bool {
	return item.isEnabled()
}

func (item *MenuItem) IsSelected() bool {
	return item.isSelected()
}
