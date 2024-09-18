package hooks

import (
	"bytes"

	mqtt "github.com/mochi-mqtt/server/v2"
)

type CustomHookOptions struct {
	Server *mqtt.Server
}

type CustomHook struct {
	mqtt.HookBase
	config *CustomHookOptions
}

func (h *CustomHook) ID() string {
	return "custom-hook"
}

func (h *CustomHook) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt.OnConnect,
		mqtt.OnDisconnect,
		mqtt.OnConnectAuthenticate,
		mqtt.OnACLCheck,
		mqtt.OnSubscribe,
		mqtt.OnPublish,
		mqtt.OnPublished,
		mqtt.OnWill,
		mqtt.OnSysInfoTick,
		mqtt.OnStarted,
		mqtt.OnStopped,
	}, []byte{b})
}

func (h *CustomHook) Init(config any) error {
	h.Log.Info("initialised")
	if _, ok := config.(*CustomHookOptions); !ok && config != nil {
		return mqtt.ErrInvalidConfigType
	}

	h.config = config.(*CustomHookOptions)
	if h.config.Server == nil {
		return mqtt.ErrInvalidConfigType
	}
	return nil
}
