package device

// FocuserInstance contains ASCOM Focuser property getters.
type FocuserInstance struct {
	DeviceInstance
	Absolute           bool
	IsMoving           bool
	MaxIncrement       int
	MaxStep            int
	Position           int
	StepSize           float64
	TempComp           bool
	TempCompAvailable  bool
	Temperature        float64
}

// Focuser is the interface for ASCOM Focuser set methods.
type Focuser interface {
	Device
	Halt() error
	Move(position int) error
	SetTempComp(tempComp bool) error
}
