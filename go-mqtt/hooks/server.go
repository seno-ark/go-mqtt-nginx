package hooks

import "github.com/mochi-mqtt/server/v2/system"

// OnStarted is called when the server starts.
func (h *CustomHook) OnStarted() {
	// h.Log.Info("\n# OnStarted")
}

// OnStopped is called when the server stops.
func (h *CustomHook) OnStopped() {
	// h.Log.Info("\n# OnStopped")
}

func (h *CustomHook) OnSysInfoTick(sysInfo *system.Info) {
	// h.Log.Info("\n# OnSysInfoTick",
	// 	"sysInfo ", sysInfo,
	// )
}
