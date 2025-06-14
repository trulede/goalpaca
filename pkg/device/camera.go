package device

// CameraInstance contains ASCOM Camera property getters.
type CameraInstance struct {
	DeviceInstance
	CanAbortExposure      bool
	CanAsymmetricBin      bool
	CanFastReadout        bool
	CanGetCoolerPower     bool
	CanPulseGuide         bool
	CanSetCCDTemperature  bool
	CanStopExposure       bool
	CCDTemperature        float64
	CoolerOn              bool
	CoolerPower           float64
	ElectronsPerADU       float64
	ExposureMax           float64
	ExposureMin           float64
	ExposureResolution    float64
	FastReadout           bool
	FullWellCapacity      float64
	HasShutter            bool
	HeatSinkTemperature   float64
	ImageReady            bool
	LastExposureDuration  float64
	LastExposureStartTime string
	MaxADU                int
	MaxBinX               int
	MaxBinY               int
	NumX                  int
	NumY                  int
	PixelSizeX            float64
	PixelSizeY            float64
	ReadoutMode           int
	ReadoutModes          []string
	SensorName            string
	SensorType            int
}

// Camera is the interface for ASCOM Camera set methods.
type Camera interface {
	Device
	AbortExposure() error
	BeginExposure(duration float64, light bool) error
	SetCCDTemperature(temp float64) error
	SetCoolerOn(on bool) error
	SetFastReadout(fast bool) error
	SetReadoutMode(mode int) error
	SetStartX(x int) error
	SetStartY(y int) error
	SetNumX(x int) error
	SetNumY(y int) error
	StartExposure(duration float64, light bool) error
	StopExposure() error
}
