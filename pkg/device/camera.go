package device

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

// RegisterCameraEndpoints registers all Camera endpoints to the given Gin router group.
func RegisterCameraEndpoints(rg *gin.RouterGroup, devices DeviceTree) {
	getProps := map[string]func(*CameraInstance) interface{}{
		"canabortexposure":      func(c *CameraInstance) interface{} { return c.CanAbortExposure },
		"canasymmetricbin":      func(c *CameraInstance) interface{} { return c.CanAsymmetricBin },
		"canfastreadout":        func(c *CameraInstance) interface{} { return c.CanFastReadout },
		"cangetcoolerpower":     func(c *CameraInstance) interface{} { return c.CanGetCoolerPower },
		"canpulseguide":         func(c *CameraInstance) interface{} { return c.CanPulseGuide },
		"cansetccdtemperature":  func(c *CameraInstance) interface{} { return c.CanSetCCDTemperature },
		"canstopexposure":       func(c *CameraInstance) interface{} { return c.CanStopExposure },
		"ccdtemperature":        func(c *CameraInstance) interface{} { return c.CCDTemperature },
		"cooleron":             func(c *CameraInstance) interface{} { return c.CoolerOn },
		"coolerpower":           func(c *CameraInstance) interface{} { return c.CoolerPower },
		"electronsperadu":       func(c *CameraInstance) interface{} { return c.ElectronsPerADU },
		"exposuremax":           func(c *CameraInstance) interface{} { return c.ExposureMax },
		"exposuremin":           func(c *CameraInstance) interface{} { return c.ExposureMin },
		"exposureresolution":    func(c *CameraInstance) interface{} { return c.ExposureResolution },
		"fastreadout":           func(c *CameraInstance) interface{} { return c.FastReadout },
		"fullwellcapacity":      func(c *CameraInstance) interface{} { return c.FullWellCapacity },
		"hasshutter":            func(c *CameraInstance) interface{} { return c.HasShutter },
		"heatsinktemperature":   func(c *CameraInstance) interface{} { return c.HeatSinkTemperature },
		"imageready":            func(c *CameraInstance) interface{} { return c.ImageReady },
		"lastexposureduration":  func(c *CameraInstance) interface{} { return c.LastExposureDuration },
		"lastexposurestarttime": func(c *CameraInstance) interface{} { return c.LastExposureStartTime },
		"maxadu":                func(c *CameraInstance) interface{} { return c.MaxADU },
		"maxbinx":               func(c *CameraInstance) interface{} { return c.MaxBinX },
		"maxbiny":               func(c *CameraInstance) interface{} { return c.MaxBinY },
		"numx":                  func(c *CameraInstance) interface{} { return c.NumX },
		"numy":                  func(c *CameraInstance) interface{} { return c.NumY },
		"pixelsizex":            func(c *CameraInstance) interface{} { return c.PixelSizeX },
		"pixelsizey":            func(c *CameraInstance) interface{} { return c.PixelSizeY },
		"readoutmode":           func(c *CameraInstance) interface{} { return c.ReadoutMode },
		"readoutmodes":          func(c *CameraInstance) interface{} { return c.ReadoutModes },
		"sensorname":            func(c *CameraInstance) interface{} { return c.SensorName },
		"sensortype":            func(c *CameraInstance) interface{} { return c.SensorType },
	}
	for prop, getter := range getProps {
		prop := prop
		getter := getter
		rg.GET("/:number/"+prop, func(c *gin.Context) {
			numStr := c.Param("number")
			num, err := strconv.Atoi(numStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid camera number"})
				return
			}
			camera, ok := devices["Camera"][DeviceIndex(num)]
			if !ok {
				c.JSON(http.StatusNotFound, gin.H{"error": "camera not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Value": getter(camera)})
		})
	}
}
