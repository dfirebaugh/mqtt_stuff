
# mqtt_stuff

This is example code for making an mqtt device to act against [github.com/hackrva/memberdashboard](github.com/hackrva/memberdashboard)


# Resource MQTT Events
## Subscribe
| topic | subscribe_data |
| ----- | ---- |
| [resourcename]/adduser | *rfidtag |
| [resourcename]/deleteruser | *rfidtag |
| [resourcename]/request_verify | *rfidtag |
| [resourcename]/request_hash | |

## Publish
| topic | publish_data |
| ----- | ---- |
| [resourcename]/access | *rfidtag |
| [resourcename]/heartbeat | *time |
| [resourcename]/respond_verify | *rfidtag || maybe something useful |
| [resourcename]/respond_hash | a hash of the rfid tags |

