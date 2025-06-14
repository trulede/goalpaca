package device

// SwitchInstance contains ASCOM Switch property getters.
type SwitchInstance struct {
	DeviceInstance
	MaxSwitch int
}

// Switch is the interface for ASCOM Switch set methods.
type Switch interface {
	Device
	SetValue(index int, value float64) error
}

// SwitchProperties provides property accessors for ASCOM Switch.
type SwitchProperties interface {
	CanWrite(index int) bool
	Description(index int) (string, error)
	MaxValue(index int) (float64, error)
	MinValue(index int) (float64, error)
	Name(index int) (string, error)
	State(index int) (bool, error)
	Value(index int) (float64, error)
}
