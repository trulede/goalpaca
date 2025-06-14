package server

import (
	"github.com/gin-gonic/gin"
	"workspaces/goalpaca/pkg/device"
)

// Example in-memory device tree for demonstration
var devices = device.DeviceTree{
	"Focuser": {
		0: &device.FocuserInstance{DeviceInstance: device.DeviceInstance{Type: "Focuser", Number: 0, Name: "Focuser0"}},
	},
	"Camera": {
		0: &device.CameraInstance{DeviceInstance: device.DeviceInstance{Type: "Camera", Number: 0, Name: "Camera0"}},
	},
	"Telescope": {
		0: &device.TelescopeInstance{DeviceInstance: device.DeviceInstance{Type: "Telescope", Number: 0, Name: "Telescope0"}},
	},
	"FilterWheel": {
		0: &device.FilterWheelInstance{DeviceInstance: device.DeviceInstance{Type: "FilterWheel", Number: 0, Name: "FilterWheel0"}},
	},
	"Dome": {
		0: &device.DomeInstance{DeviceInstance: device.DeviceInstance{Type: "Dome", Number: 0, Name: "Dome0"}},
	},
	"Rotator": {
		0: &device.RotatorInstance{DeviceInstance: device.DeviceInstance{Type: "Rotator", Number: 0, Name: "Rotator0"}},
	},
	"CoverCalibrator": {
		0: &device.CoverCalibratorInstance{DeviceInstance: device.DeviceInstance{Type: "CoverCalibrator", Number: 0, Name: "CoverCalibrator0"}},
	},
	"Switch": {
		0: &device.SwitchInstance{DeviceInstance: device.DeviceInstance{Type: "Switch", Number: 0, Name: "Switch0"}},
	},
	"SafetyMonitor": {
		0: &device.SafetyMonitorInstance{DeviceInstance: device.DeviceInstance{Type: "SafetyMonitor", Number: 0, Name: "SafetyMonitor0"}},
	},
	"ObservingConditions": {
		0: &device.ObservingConditionsInstance{DeviceInstance: device.DeviceInstance{Type: "ObservingConditions", Number: 0, Name: "ObservingConditions0"}},
	},
}

func StartServer() {
	r := gin.Default()
	device.RegisterFocuserEndpoints(r.Group("/api/v1/focuser"), devices)
	device.RegisterCameraEndpoints(r.Group("/api/v1/camera"), devices)
	device.RegisterTelescopeEndpoints(r.Group("/api/v1/telescope"), devices)
	device.RegisterFilterWheelEndpoints(r.Group("/api/v1/filterwheel"), devices)
	device.RegisterDomeEndpoints(r.Group("/api/v1/dome"), devices)
	device.RegisterRotatorEndpoints(r.Group("/api/v1/rotator"), devices)
	device.RegisterCoverCalibratorEndpoints(r.Group("/api/v1/covercalibrator"), devices)
	device.RegisterSwitchEndpoints(r.Group("/api/v1/switch"), devices)
	device.RegisterSafetyMonitorEndpoints(r.Group("/api/v1/safetymonitor"), devices)
	device.RegisterObservingConditionsEndpoints(r.Group("/api/v1/observingconditions"), devices)
	r.Run(":8080")
}
