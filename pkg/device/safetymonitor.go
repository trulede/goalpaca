package device

// SafetyMonitorInstance contains ASCOM SafetyMonitor property getters.
type SafetyMonitorInstance struct {
	DeviceInstance
	IsSafe bool
}

// SafetyMonitor is the interface for ASCOM SafetyMonitor set methods (none for this type).
type SafetyMonitor interface {
	Device
}
