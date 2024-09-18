package hooks

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

func (h *CustomHook) OnPublished(cl *mqtt.Client, pk packets.Packet) {
	username := string(cl.Properties.Username)

	h.Log.Info("\n# OnPublished",
		"Client ID", cl.ID,
		"Username", username,
		"Topic", pk.TopicName,
		"Payload", string(pk.Payload),
	)
}

func (h *CustomHook) OnWill(cl *mqtt.Client, will mqtt.Will) (mqtt.Will, error) {
	username := string(cl.Properties.Username)

	h.Log.Info("\n# OnWill",
		"Client ID", cl.ID,
		"Username", username,
		"will", will,
	)
	return will, nil
}
