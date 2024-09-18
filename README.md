# go-mqtt-nginx

MQTT Implementation with Mochi-MQTT, NGINX, MQTT.js, and Paho Python

## Run with Docker Compose

```
docker-compose up
```

## Access the Dashboard

http://localhost/?device_id=peach&password=peach1

## Run the device
```
cd device
python3 -m venv venv
venv/bin/activate
pip install -r requirements.txt

python3 main.py peach peach1
```