package hooks

import (
	"fmt"
	"strings"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

func (h *CustomHook) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
	// username := string(cl.Properties.Username)
	// password := string(pk.Connect.Password)

	// h.Log.Info("\n# OnConnectAuthenticate",
	// 	"Client ID", cl.ID,
	// 	"Username", username,
	// 	"Password", password,
	// )
	// dummy auth
	return true
}

func (h *CustomHook) OnACLCheck(cl *mqtt.Client, topic string, write bool) bool {
	username := string(cl.Properties.Username)

	// h.Log.Info("\n# OnACLCheck",
	// 	"Client ID", cl.ID,
	// 	"Username", username,
	// 	"topic", topic,
	// )

	// dummy acl
	if write && !strings.HasPrefix(topic, "/device/"+username) {
		h.Log.Debug(fmt.Sprintf("%s doesnt have write permission to %s topic", username, topic))
		return false
	}
	return true
}

func (h *CustomHook) OnSubscribe(cl *mqtt.Client, pk packets.Packet) packets.Packet {
	// username := string(cl.Properties.Username)

	// h.Log.Info("\n# OnSubscribe",
	// 	"Client ID", cl.ID,
	// 	"Username", username,
	// 	"TopicName", pk.TopicName,
	// )
	return pk
}

func (h *CustomHook) OnPublish(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	// username := string(cl.Properties.Username)

	// h.Log.Info("\n# OnPublish",
	// 	"Client ID", cl.ID,
	// 	"Username", username,
	// 	"TopicName", pk.TopicName,
	// )
	return pk, nil
}

func (h *CustomHook) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	// username := string(cl.Properties.Username)

	// h.Log.Info("\n# OnConnect",
	// 	"Client ID", cl.ID,
	// 	"Username", username,
	// )
	return nil
}

func (h *CustomHook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	// username := string(cl.Properties.Username)

	// h.Log.Info("\n# OnDisconnect",
	// 	"Client ID", cl.ID,
	// 	"Username", username,
	// 	"expire", expire,
	// 	"err", err,
	// )
	// payload, _ := json.Marshal(map[string]any{
	// 	"timestamp": time.Now().UTC().Unix(),
	// 	"status":    "offline",
	// })
	// h.config.Server.Publish(fmt.Sprintf("device/%s/status", username), payload, true, 1)
}
