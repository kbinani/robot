package robot

// PwOp represents actions for various power management systems.
type PwOp int

// PwOp codes for power management actions.
const (
	MonitorOff = iota
	MonitorOn
)

// Pw is a function to control various power management system.
func Pw(op PwOp) {
	pw(op)
}
