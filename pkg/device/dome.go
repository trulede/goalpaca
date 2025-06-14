package device

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

// RegisterDomeEndpoints registers all Dome endpoints to the given Gin router group.
func RegisterDomeEndpoints(rg *gin.RouterGroup, devices DeviceTree) {
	getProps := map[string]func(*DomeInstance) interface{}{
		"altitude":       func(d *DomeInstance) interface{} { return d.Altitude },
		"athome":         func(d *DomeInstance) interface{} { return d.AtHome },
		"atpark":         func(d *DomeInstance) interface{} { return d.AtPark },
		"azimuth":        func(d *DomeInstance) interface{} { return d.Azimuth },
		"canfindhome":    func(d *DomeInstance) interface{} { return d.CanFindHome },
		"canpark":        func(d *DomeInstance) interface{} { return d.CanPark },
		"cansetaltitude": func(d *DomeInstance) interface{} { return d.CanSetAltitude },
		"cansetazimuth":  func(d *DomeInstance) interface{} { return d.CanSetAzimuth },
		"cansetpark":     func(d *DomeInstance) interface{} { return d.CanSetPark },
		"cansetshutter":  func(d *DomeInstance) interface{} { return d.CanSetShutter },
		"canslave":       func(d *DomeInstance) interface{} { return d.CanSlave },
		"cansyncazimuth": func(d *DomeInstance) interface{} { return d.CanSyncAzimuth },
		"shutterstatus":  func(d *DomeInstance) interface{} { return d.ShutterStatus },
		"slaved":         func(d *DomeInstance) interface{} { return d.Slaved },
		"slewing":        func(d *DomeInstance) interface{} { return d.Slewing },
	}
	for prop, getter := range getProps {
		prop := prop
		getter := getter
		rg.GET("/:number/"+prop, func(c *gin.Context) {
			numStr := c.Param("number")
			num, err := strconv.Atoi(numStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dome number"})
				return
			}
			dome, ok := devices["Dome"][DeviceIndex(num)]
			if !ok {
				c.JSON(http.StatusNotFound, gin.H{"error": "dome not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Value": getter(dome)})
		})
	}
}
