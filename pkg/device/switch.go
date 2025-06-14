package device

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SwitchInstance contains ASCOM Switch property getters.
type SwitchInstance struct {
	DeviceInstance
	MaxSwitch int
}

// Switch is the interface for ASCOM Switch set methods.
type Switch interface {
	Device
	SetValue(index int, value float64) error
}

// SwitchProperties provides property accessors for ASCOM Switch.
type SwitchProperties interface {
	CanWrite(index int) bool
	Description(index int) (string, error)
	MaxValue(index int) (float64, error)
	MinValue(index int) (float64, error)
	Name(index int) (string, error)
	State(index int) (bool, error)
	Value(index int) (float64, error)
}

// RegisterSwitchEndpoints registers all Switch endpoints to the given Gin router group.
func RegisterSwitchEndpoints(rg *gin.RouterGroup, devices DeviceTree) {
	getProps := map[string]func(*SwitchInstance) interface{}{
		"maxswitch": func(s *SwitchInstance) interface{} { return s.MaxSwitch },
	}
	for prop, getter := range getProps {
		prop := prop
		getter := getter
		rg.GET("/:number/"+prop, func(c *gin.Context) {
			numStr := c.Param("number")
			num, err := strconv.Atoi(numStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid switch number"})
				return
			}
			sw, ok := devices["Switch"][DeviceIndex(num)]
			if !ok {
				c.JSON(http.StatusNotFound, gin.H{"error": "switch not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Value": getter(sw)})
		})
	}
}
