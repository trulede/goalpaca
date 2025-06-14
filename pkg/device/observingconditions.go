package device

// ObservingConditionsInstance contains ASCOM ObservingConditions property getters.
type ObservingConditionsInstance struct {
	DeviceInstance
	CloudCover     float64
	DewPoint       float64
	Humidity       float64
	Pressure       float64
	RainRate       float64
	SkyBrightness  float64
	SkyQuality     float64
	SkyTemperature float64
	StarFWHM       float64
	Temperature    float64
	WindDirection  float64
	WindGust       float64
	WindSpeed      float64
}

// ObservingConditions is the interface for ASCOM ObservingConditions set methods.
type ObservingConditions interface {
	Device
}
