package device

// FilterWheelInstance contains ASCOM FilterWheel property getters.
type FilterWheelInstance struct {
	DeviceInstance
	FocusOffsets []int
	Names        []string
	Position     int
}

// FilterWheel is the interface for ASCOM FilterWheel set methods.
type FilterWheel interface {
	Device
	SetPosition(position int) error
}
