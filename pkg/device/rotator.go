package device

// RotatorInstance contains ASCOM Rotator property getters.
type RotatorInstance struct {
	DeviceInstance
	CanReverse         bool
	IsMoving           bool
	MechanicalPosition float64
	Position           float64
	Reverse            bool
	StepSize           float64
	TargetPosition     float64
}

// Rotator is the interface for ASCOM Rotator set methods.
type Rotator interface {
	Device
	Move(position float64) error
	MoveAbsolute(position float64) error
	MoveMechanical(position float64) error
	SetReverse(reverse bool) error
	Halt() error
}
