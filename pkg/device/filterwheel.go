package device

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FilterWheelInstance contains ASCOM FilterWheel property getters.
type FilterWheelInstance struct {
	DeviceInstance
	FocusOffsets []int
	Names        []string
	Position     int
}

// FilterWheel is the interface for ASCOM FilterWheel set methods.
type FilterWheel interface {
	Device
	SetPosition(position int) error
}

// RegisterFilterWheelEndpoints registers all FilterWheel endpoints to the given Gin router group.
func RegisterFilterWheelEndpoints(rg *gin.RouterGroup, devices DeviceTree) {
	getProps := map[string]func(*FilterWheelInstance) interface{}{
		"focusoffsets": func(f *FilterWheelInstance) interface{} { return f.FocusOffsets },
		"names":        func(f *FilterWheelInstance) interface{} { return f.Names },
		"position":     func(f *FilterWheelInstance) interface{} { return f.Position },
	}
	for prop, getter := range getProps {
		prop := prop
		getter := getter
		rg.GET("/:number/"+prop, func(c *gin.Context) {
			numStr := c.Param("number")
			num, err := strconv.Atoi(numStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid filterwheel number"})
				return
			}
			fw, ok := devices["FilterWheel"][DeviceIndex(num)]
			if !ok {
				c.JSON(http.StatusNotFound, gin.H{"error": "filterwheel not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Value": getter(fw)})
		})
	}
}
