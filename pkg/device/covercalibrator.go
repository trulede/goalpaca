package device

// CoverCalibratorInstance contains ASCOM CoverCalibrator property getters.
type CoverCalibratorInstance struct {
	DeviceInstance
	Brightness      int
	CalibratorState int
	CoverState      int
	MaxBrightness   int
}

// CoverCalibrator is the interface for ASCOM CoverCalibrator set methods.
type CoverCalibrator interface {
	Device
	CalibratorOn(brightness int) error
	CalibratorOff() error
	OpenCover() error
	CloseCover() error
	HaltCover() error
}
