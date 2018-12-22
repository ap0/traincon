# traincon

Used to control AC model train switches via a Raspberry Pi.  API is compatible with Homebridge plugin.

![switch in action](resources/working.gif)

This project uses `go mod` and thus requires Go 1.11.

```sh
# build locally
go build
# build for raspberry pi
GOOS=linux GOARM=7 GOARCH=arm go build
./traincon config.yml
```

Example config file:

```yaml
port: 8080
switches:
  1: # ID of switch for below routes
    on_off_pin: 35 # Pin for on/off relay
    direction_pin: 37 # Pin for changing current direction
  2:
    on_off_pin: 11
    direction_pin: 13
  3:
    on_off_pin: 16
    direction_pin: 18
  4:
    on_off_pin: 8
    direction_pin: 10
```

Switch IDs can be any string; I used numbers for simplicity.

The routes are simple:

```
GET /switch/{id}/on
GET /switch/{id}/off
GET /switch/{id}/toggle
GET /switch/{id}/status
```

It's meant to be used with the [Homebridge](https://github.com/nfarina/homebridge) plugin [homebridge-http-switch](https://github.com/Supereg/homebridge-http-switch).

Sample Homebridge config:

```json
{
    "accessories": [
        {
            "accessory": "HTTP-SWITCH",
            "name": "Turnout Switch 1",
            "switchType": "stateful",
            "onUrl": "http://trainpi:8080/switch/1/on",
            "offUrl": "http://trainpi:8080/switch/1/off",
            "statusUrl": "http://trainpi:8080/switch/1/status"
        }
    ]
}
```

## Wiring

Each turnout switch should be wired to two SPDT relays.  The relays are meant to simulate a momentary DPDT switch that is used to change the direction of the turnout.  Because they are powered using AC, you need to use two diodes as half-wave rectifiers to only allow the current to flow in the desired direction.

The code will change the direction of the track (the top relay), then momentarily energize the on/off relay to make the track move.

![wiring diagram](resources/diagram.png)
![board layout](resources/board_layout.png)