package robot

// Op represents actions for buttons or keys.
type Op int

// Op codes for button/key actions.
const (
	Click Op = iota
	Down
	Up
)
