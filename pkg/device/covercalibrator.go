package device

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

// RegisterCoverCalibratorEndpoints registers all CoverCalibrator endpoints to the given Gin router group.
func RegisterCoverCalibratorEndpoints(rg *gin.RouterGroup, devices DeviceTree) {
	getProps := map[string]func(*CoverCalibratorInstance) interface{}{
		"brightness":      func(c *CoverCalibratorInstance) interface{} { return c.Brightness },
		"calibratorstate": func(c *CoverCalibratorInstance) interface{} { return c.CalibratorState },
		"coverstate":      func(c *CoverCalibratorInstance) interface{} { return c.CoverState },
		"maxbrightness":   func(c *CoverCalibratorInstance) interface{} { return c.MaxBrightness },
	}
	for prop, getter := range getProps {
		prop := prop
		getter := getter
		rg.GET("/:number/"+prop, func(c *gin.Context) {
			numStr := c.Param("number")
			num, err := strconv.Atoi(numStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid covercalibrator number"})
				return
			}
			cc, ok := devices["CoverCalibrator"][DeviceIndex(num)]
			if !ok {
				c.JSON(http.StatusNotFound, gin.H{"error": "covercalibrator not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Value": getter(cc)})
		})
	}
}
