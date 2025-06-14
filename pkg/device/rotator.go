package device

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RotatorInstance contains ASCOM Rotator property getters.
type RotatorInstance struct {
	DeviceInstance
	CanReverse         bool
	IsMoving           bool
	MechanicalPosition float64
	Position           float64
	Reverse            bool
	StepSize           float64
	TargetPosition     float64
}

// Rotator is the interface for ASCOM Rotator set methods.
type Rotator interface {
	Device
	Move(position float64) error
	MoveAbsolute(position float64) error
	MoveMechanical(position float64) error
	SetReverse(reverse bool) error
	Halt() error
}

// RegisterRotatorEndpoints registers all Rotator endpoints to the given Gin router group.
func RegisterRotatorEndpoints(rg *gin.RouterGroup, devices DeviceTree) {
	getProps := map[string]func(*RotatorInstance) interface{}{
		"canreverse":         func(r *RotatorInstance) interface{} { return r.CanReverse },
		"ismoving":           func(r *RotatorInstance) interface{} { return r.IsMoving },
		"mechanicalposition": func(r *RotatorInstance) interface{} { return r.MechanicalPosition },
		"position":           func(r *RotatorInstance) interface{} { return r.Position },
		"reverse":            func(r *RotatorInstance) interface{} { return r.Reverse },
		"stepsize":           func(r *RotatorInstance) interface{} { return r.StepSize },
		"targetposition":     func(r *RotatorInstance) interface{} { return r.TargetPosition },
	}
	for prop, getter := range getProps {
		prop := prop
		getter := getter
		rg.GET("/:number/"+prop, func(c *gin.Context) {
			numStr := c.Param("number")
			num, err := strconv.Atoi(numStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid rotator number"})
				return
			}
			rotator, ok := devices["Rotator"][DeviceIndex(num)]
			if !ok {
				c.JSON(http.StatusNotFound, gin.H{"error": "rotator not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Value": getter(rotator)})
		})
	}
}
