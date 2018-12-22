# traincon

Used to control AC model train switches via a Raspberry Pi.  API is compatible with Homebridge plugin.

![switch in action](resources/working.gif)

This project uses `go mod` and thus requires Go 1.11.

```
# build locally
go build
# build for raspberry pi
GOOS=linux GOARM=7 GOARCH=arm go build
./traincon
```

The routes are simple:

```
GET /toggle/{switchNum}
GET /on/{switchNum}
GET /off/{switchNum}
GET /status/{switchNum}
```

Right now everything's hardcoded. TODO would be to have this run via config.

It's meant to be used with the [Homebridge](https://github.com/nfarina/homebridge) plugin [homebridge-http-switch](https://github.com/Supereg/homebridge-http-switch).

Sample Homebridge config:

```
{
    "accessories": [
        {
            "accessory": "HTTP-SWITCH",
            "name": "Train Switch 1",
            "switchType": "stateful",
            "onUrl": "http://raspberrypi:8080/on/1",
            "offUrl": "http://raspberrypi:8080/off/1",
            "statusUrl": "http://raspberrypi:8080/status/1"
        }
    ]
}
```

## Wiring

Each turnout switch should be wired to two SPDT relays.  The relays are meant to simulate a momentary DPDT switch that is used to change the direction of the turnout.  Because they are powered using AC, you need to use two diodes as half-wave rectifiers to only allow the current to flow in the desired direction.

The code will change the direction of the track (the top relay), then momentarily energize the on/off relay to make the track move.

![wiring diagram](resources/diagram.png)
![board layout](resources/board_layout.png)