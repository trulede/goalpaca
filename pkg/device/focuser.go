package device

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FocuserInstance contains ASCOM Focuser property getters.
type FocuserInstance struct {
	DeviceInstance
	Absolute           bool
	IsMoving           bool
	MaxIncrement       int
	MaxStep            int
	Position           int
	StepSize           float64
	TempComp           bool
	TempCompAvailable  bool
	Temperature        float64
}

// Focuser is the interface for ASCOM Focuser set methods.
type Focuser interface {
	Device
	Halt() error
	Move(position int) error
	SetTempComp(tempComp bool) error
}

// RegisterFocuserEndpoints registers all Focuser endpoints to the given Gin router group.
func RegisterFocuserEndpoints(rg *gin.RouterGroup, devices DeviceTree) {
	rg.GET("/:number", func(c *gin.Context) {
		numStr := c.Param("number")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid focuser number"})
			return
		}
		focuser, ok := devices["Focuser"][DeviceIndex(num)]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "focuser not found"})
			return
		}
		c.JSON(http.StatusOK, focuser)
	})

	// GET endpoints for Focuser properties (ASCOM API)
	getProps := map[string]func(*FocuserInstance) interface{}{
		"absolute":           func(f *FocuserInstance) interface{} { return f.Absolute },
		"ismoving":           func(f *FocuserInstance) interface{} { return f.IsMoving },
		"maxincrement":       func(f *FocuserInstance) interface{} { return f.MaxIncrement },
		"maxstep":            func(f *FocuserInstance) interface{} { return f.MaxStep },
		"position":           func(f *FocuserInstance) interface{} { return f.Position },
		"stepsize":           func(f *FocuserInstance) interface{} { return f.StepSize },
		"tempcomp":           func(f *FocuserInstance) interface{} { return f.TempComp },
		"tempcompavailable":  func(f *FocuserInstance) interface{} { return f.TempCompAvailable },
		"temperature":        func(f *FocuserInstance) interface{} { return f.Temperature },
	}
	for prop, getter := range getProps {
		// e.g. /api/v1/focuser/:number/absolute
		prop := prop
		getter := getter
		rg.GET("/:number/"+prop, func(c *gin.Context) {
			numStr := c.Param("number")
			num, err := strconv.Atoi(numStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid focuser number"})
				return
			}
			focuser, ok := devices["Focuser"][DeviceIndex(num)]
			if !ok {
				c.JSON(http.StatusNotFound, gin.H{"error": "focuser not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Value": getter(focuser)})
		})
	}

	rg.PUT("/:number/halt", func(c *gin.Context) {
		numStr := c.Param("number")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid focuser number"})
			return
		}
		focuser, ok := devices["Focuser"][DeviceIndex(num)]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "focuser not found"})
			return
		}
		// Here you would call the actual Halt() method on the real device implementation
		c.JSON(http.StatusOK, gin.H{"result": "Focuser halted (stub)"})
	})

	rg.PUT("/:number/move", func(c *gin.Context) {
		numStr := c.Param("number")
		posStr := c.Query("Position")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid focuser number"})
			return
		}
		pos, err := strconv.Atoi(posStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid position"})
			return
		}
		focuser, ok := devices["Focuser"][DeviceIndex(num)]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "focuser not found"})
			return
		}
		// Here you would call the actual Move(position) method on the real device implementation
		focuser.Position = pos
		c.JSON(http.StatusOK, gin.H{"result": "Focuser moved (stub)", "position": pos})
	})
}
