package device

// DeviceType represents the type of an ASCOM device (e.g., "Focuser", "Camera").
type DeviceType string

// DeviceIndex represents the device number/index.
type DeviceIndex int

// DeviceInstance contains common ASCOM device properties (getters).
type DeviceInstance struct {
	Type               DeviceType
	Number             DeviceIndex
	
	Connected          bool
	Description        string
	DriverInfo         string
	DriverVersion      string
	InterfaceVersion   int
	Name               string
	SupportedActions   []string
}

// DeviceTree is a map where the key is the device type (DeviceType),
// and the value is another map from device number (DeviceIndex) to DeviceInstance.
type DeviceTree map[DeviceType]map[DeviceIndex]*DeviceInstance

// Device is the interface for common ASCOM device set methods (setters).
type Device interface {
	SetConnected(connected bool) error
	Connect() error
	Disconnect() error
	Action(actionName string, actionParameters string) (string, error)
	CommandBlind(command string, raw bool) error
	CommandBool(command string, raw bool) (bool, error)
	CommandString(command string, raw bool) (string, error)
	Refresh(instance *DeviceInstance) error
}
