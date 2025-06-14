package device

// DomeInstance contains ASCOM Dome property getters.
type DomeInstance struct {
	DeviceInstance
	Altitude        float64
	AtHome          bool
	AtPark          bool
	Azimuth         float64
	CanFindHome     bool
	CanPark         bool
	CanSetAltitude  bool
	CanSetAzimuth   bool
	CanSetPark      bool
	CanSetShutter   bool
	CanSlave        bool
	CanSyncAzimuth  bool
	ShutterStatus   int
	Slaved          bool
	Slewing         bool
}

// Dome is the interface for ASCOM Dome set methods.
type Dome interface {
	Device
	CloseShutter() error
	FindHome() error
	OpenShutter() error
	Park() error
	SetPark() error
	SetShutter(status int) error
	SetSlaved(slaved bool) error
	SlewToAltitude(altitude float64) error
	SlewToAzimuth(azimuth float64) error
	SyncToAzimuth(azimuth float64) error
}
