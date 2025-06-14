package device

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

// RegisterTelescopeEndpoints registers all Telescope endpoints to the given Gin router group.
func RegisterTelescopeEndpoints(rg *gin.RouterGroup, devices DeviceTree) {
	getProps := map[string]func(*TelescopeInstance) interface{}{
		"alignmentmode":         func(t *TelescopeInstance) interface{} { return t.AlignmentMode },
		"altitude":              func(t *TelescopeInstance) interface{} { return t.Altitude },
		"aperturearea":          func(t *TelescopeInstance) interface{} { return t.ApertureArea },
		"aperturediameter":      func(t *TelescopeInstance) interface{} { return t.ApertureDiameter },
		"athome":                func(t *TelescopeInstance) interface{} { return t.AtHome },
		"atpark":                func(t *TelescopeInstance) interface{} { return t.AtPark },
		"azimuth":               func(t *TelescopeInstance) interface{} { return t.Azimuth },
		"canfindhome":           func(t *TelescopeInstance) interface{} { return t.CanFindHome },
		"canpark":               func(t *TelescopeInstance) interface{} { return t.CanPark },
		"canpulseguide":         func(t *TelescopeInstance) interface{} { return t.CanPulseGuide },
		"cansetdeclinationrate": func(t *TelescopeInstance) interface{} { return t.CanSetDeclinationRate },
		"cansetguiderates":      func(t *TelescopeInstance) interface{} { return t.CanSetGuideRates },
		"cansetpark":            func(t *TelescopeInstance) interface{} { return t.CanSetPark },
		"cansetpierside":        func(t *TelescopeInstance) interface{} { return t.CanSetPierSide },
		"cansetrightascensionrate": func(t *TelescopeInstance) interface{} { return t.CanSetRightAscensionRate },
		"cansettracking":        func(t *TelescopeInstance) interface{} { return t.CanSetTracking },
		"canslew":               func(t *TelescopeInstance) interface{} { return t.CanSlew },
		"canslewaltaz":          func(t *TelescopeInstance) interface{} { return t.CanSlewAltAz },
		"canslewaltazasync":     func(t *TelescopeInstance) interface{} { return t.CanSlewAltAzAsync },
		"canslewasync":          func(t *TelescopeInstance) interface{} { return t.CanSlewAsync },
		"cansync":               func(t *TelescopeInstance) interface{} { return t.CanSync },
		"cansyncaltaz":          func(t *TelescopeInstance) interface{} { return t.CanSyncAltAz },
		"canunpark":             func(t *TelescopeInstance) interface{} { return t.CanUnpark },
		"declination":           func(t *TelescopeInstance) interface{} { return t.Declination },
		"declinationrate":       func(t *TelescopeInstance) interface{} { return t.DeclinationRate },
		"doesrefraction":        func(t *TelescopeInstance) interface{} { return t.DoesRefraction },
		"equatorialsystem":      func(t *TelescopeInstance) interface{} { return t.EquatorialSystem },
		"guideratedeclination":  func(t *TelescopeInstance) interface{} { return t.GuideRateDeclination },
		"guideraterightascension": func(t *TelescopeInstance) interface{} { return t.GuideRateRightAscension },
		"ispulseguiding":        func(t *TelescopeInstance) interface{} { return t.IsPulseGuiding },
		"pierside":              func(t *TelescopeInstance) interface{} { return t.PierSide },
		"rightascension":        func(t *TelescopeInstance) interface{} { return t.RightAscension },
		"rightascensionrate":    func(t *TelescopeInstance) interface{} { return t.RightAscensionRate },
		"targetdeclination":     func(t *TelescopeInstance) interface{} { return t.TargetDeclination },
		"targetrightascension":  func(t *TelescopeInstance) interface{} { return t.TargetRightAscension },
		"tracking":              func(t *TelescopeInstance) interface{} { return t.Tracking },
		"trackingrate":          func(t *TelescopeInstance) interface{} { return t.TrackingRate },
		"trackingrates":         func(t *TelescopeInstance) interface{} { return t.TrackingRates },
		"utcdate":               func(t *TelescopeInstance) interface{} { return t.UTCDate },
	}
	for prop, getter := range getProps {
		prop := prop
		getter := getter
		rg.GET("/:number/"+prop, func(c *gin.Context) {
			numStr := c.Param("number")
			num, err := strconv.Atoi(numStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid telescope number"})
				return
			}
			telescope, ok := devices["Telescope"][DeviceIndex(num)]
			if !ok {
				c.JSON(http.StatusNotFound, gin.H{"error": "telescope not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Value": getter(telescope)})
		})
	}
}
