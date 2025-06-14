package device

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SafetyMonitorInstance contains ASCOM SafetyMonitor property getters.
type SafetyMonitorInstance struct {
	DeviceInstance
	IsSafe bool
}

// SafetyMonitor is the interface for ASCOM SafetyMonitor set methods (none for this type).
type SafetyMonitor interface {
	Device
}

// RegisterSafetyMonitorEndpoints registers all SafetyMonitor endpoints to the given Gin router group.
func RegisterSafetyMonitorEndpoints(rg *gin.RouterGroup, devices DeviceTree) {
	getProps := map[string]func(*SafetyMonitorInstance) interface{}{
		"issafe": func(s *SafetyMonitorInstance) interface{} { return s.IsSafe },
	}
	for prop, getter := range getProps {
		prop := prop
		getter := getter
		rg.GET("/:number/"+prop, func(c *gin.Context) {
			numStr := c.Param("number")
			num, err := strconv.Atoi(numStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid safetymonitor number"})
				return
			}
			sm, ok := devices["SafetyMonitor"][DeviceIndex(num)]
			if !ok {
				c.JSON(http.StatusNotFound, gin.H{"error": "safetymonitor not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Value": getter(sm)})
		})
	}
}
