import json
import time
import sys
import random
import threading
import paho.mqtt.client as mqtt,sys

class Device:

    def __init__(self) -> None:
        self.device_id = ""
        self.client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2)

        self.client.on_connect = self.on_connect
        self.client.on_message = self.on_message
        self.client.on_subscribe = self.on_subscribe
        self.client.on_unsubscribe = self.on_unsubscribe
    
    def connect(self, device_id, password):
        self.device_id = device_id

        self.client.will_set("device/"+self.device_id+"/status", json.dumps({
            "timestamp": int(time.time()),
            "status": "offline (last will)",
        }), qos=1, retain=True)
        
        self.client.username_pw_set(device_id, password)
        self.client.connect("localhost", 1883, 60)

        self.scheduler()

        self.client.loop_forever()
    
    def on_connect(self, client, userdata, flags, reason_code, properties):
        print(f"Connected with result code {reason_code}")

        # Subscribing in on_connect() means that if we lose the connection and
        # reconnect then subscriptions will be renewed.
        # client.subscribe("$SYS/#")
        # client.subscribe("$SYS/#")
        # client.subscribe("presence")
        # client.subscribe("paho/temperature")
        # client.subscribe("paho/test/topic")

        self.client.subscribe("device/"+self.device_id+"/control/gateLevel")

        topic = "device/"+self.device_id+"/status"
        data = json.dumps({
            "timestamp": int(time.time()),
            "status": "online",
        })

        print(f"Publish {topic} {data}")
        self.client.publish(topic, data, qos=1, retain=True)

    def on_message(self, client, userdata, msg):
        payload = json.loads(msg.payload)
        print(f"Got message {msg.topic} {payload}")

        if msg.topic == "device/"+self.device_id+"/control/gateLevel":
            # set gate
            # publish value
            topic = "device/"+self.device_id+"/value/gateLevel"
            data = json.dumps({
                "timestamp": int(time.time()),
                "value": float(payload["value"])+0.00123456789,
            })

            print(f"Publish {topic} {data}")
            client.publish(topic, data, qos=1, retain=True)

    def on_subscribe(self, client, userdata, mid, reason_code_list, properties):
        if reason_code_list[0].is_failure:
            print(f"Broker rejected you subscription: {reason_code_list[0]}")
        else:
            print(f"Broker granted the following QoS: {reason_code_list[0].value}")

    def on_unsubscribe(self, client, userdata, mid, reason_code_list, properties):
        # Be careful, the reason_code_list is only present in MQTTv5.
        # In MQTTv3 it will always be empty
        if len(reason_code_list) == 0 or not reason_code_list[0].is_failure:
            print("unsubscribe succeeded (if SUBACK is received in MQTTv3 it success)")
        else:
            print(f"Broker replied with failure: {reason_code_list[0]}")

    def scheduler(self):
        random_float = random.uniform(1.0, 10.0)

        topic = "device/"+self.device_id+"/value/waterLevel"
        data = json.dumps({
            "timestamp": int(time.time()),
            "value": random_float,
        })

        print(f"Publish {topic} {data}")
        self.client.publish(topic, data, qos=1, retain=True)

        threading.Timer(10, self.scheduler).start()

if __name__ == "__main__":
    if len(sys.argv) > 2:
        device_id = sys.argv[1]
        password = sys.argv[2]

        device = Device()
        device.connect(device_id, password)
