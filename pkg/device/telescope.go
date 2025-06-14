package device

// TelescopeInstance contains ASCOM Telescope property getters.
type TelescopeInstance struct {
	DeviceInstance
	AlignmentMode         int
	Altitude              float64
	ApertureArea          float64
	ApertureDiameter      float64
	AtHome                bool
	AtPark                bool
	Azimuth               float64
	CanFindHome           bool
	CanPark               bool
	CanPulseGuide         bool
	CanSetDeclinationRate bool
	CanSetGuideRates      bool
	CanSetPark            bool
	CanSetPierSide        bool
	CanSetRightAscensionRate bool
	CanSetTracking        bool
	CanSlew               bool
	CanSlewAltAz          bool
	CanSlewAltAzAsync     bool
	CanSlewAsync          bool
	CanSync               bool
	CanSyncAltAz          bool
	CanUnpark             bool
	Declination           float64
	DeclinationRate       float64
	DoesRefraction        bool
	EquatorialSystem      int
	GuideRateDeclination  float64
	GuideRateRightAscension float64
	IsPulseGuiding        bool
	PierSide              int
	RightAscension        float64
	RightAscensionRate    float64
	TargetDeclination     float64
	TargetRightAscension  float64
	Tracking              bool
	TrackingRate          int
	TrackingRates         []int
	UTCDate               string
}

// Telescope is the interface for ASCOM Telescope set methods.
type Telescope interface {
	Device
	AbortSlew() error
	FindHome() error
	Park() error
	PulseGuide(direction int, duration int) error
	SetDeclinationRate(rate float64) error
	SetGuideRates(guideRateRA, guideRateDec float64) error
	SetPark() error
	SetPierSide(pierSide int) error
	SetRightAscensionRate(rate float64) error
	SetTracking(tracking bool) error
	SlewToAltAz(azimuth, altitude float64) error
	SlewToAltAzAsync(azimuth, altitude float64) error
	SlewToCoordinates(rightAscension, declination float64) error
	SlewToCoordinatesAsync(rightAscension, declination float64) error
	SlewToTarget() error
	SlewToTargetAsync() error
	SyncToAltAz(azimuth, altitude float64) error
	SyncToCoordinates(rightAscension, declination float64) error
	SyncToTarget() error
	Unpark() error
}
