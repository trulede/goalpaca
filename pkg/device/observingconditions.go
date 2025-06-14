package device

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

// RegisterObservingConditionsEndpoints registers all ObservingConditions endpoints to the given Gin router group.
func RegisterObservingConditionsEndpoints(rg *gin.RouterGroup, devices DeviceTree) {
	getProps := map[string]func(*ObservingConditionsInstance) interface{}{
		"cloudcover":     func(o *ObservingConditionsInstance) interface{} { return o.CloudCover },
		"dewpoint":       func(o *ObservingConditionsInstance) interface{} { return o.DewPoint },
		"humidity":       func(o *ObservingConditionsInstance) interface{} { return o.Humidity },
		"pressure":       func(o *ObservingConditionsInstance) interface{} { return o.Pressure },
		"rainrate":       func(o *ObservingConditionsInstance) interface{} { return o.RainRate },
		"skybrightness":  func(o *ObservingConditionsInstance) interface{} { return o.SkyBrightness },
		"skyquality":     func(o *ObservingConditionsInstance) interface{} { return o.SkyQuality },
		"skytemperature": func(o *ObservingConditionsInstance) interface{} { return o.SkyTemperature },
		"starfwhm":       func(o *ObservingConditionsInstance) interface{} { return o.StarFWHM },
		"temperature":    func(o *ObservingConditionsInstance) interface{} { return o.Temperature },
		"winddirection":  func(o *ObservingConditionsInstance) interface{} { return o.WindDirection },
		"windgust":       func(o *ObservingConditionsInstance) interface{} { return o.WindGust },
		"windspeed":      func(o *ObservingConditionsInstance) interface{} { return o.WindSpeed },
	}
	for prop, getter := range getProps {
		prop := prop
		getter := getter
		rg.GET("/:number/"+prop, func(c *gin.Context) {
			numStr := c.Param("number")
			num, err := strconv.Atoi(numStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid observingconditions number"})
				return
			}
			oc, ok := devices["ObservingConditions"][DeviceIndex(num)]
			if !ok {
				c.JSON(http.StatusNotFound, gin.H{"error": "observingconditions not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Value": getter(oc)})
		})
	}
}
