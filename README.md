# traincon

Used to control AC model train switches via a Raspberry Pi.

This project uses [gb](https://getgb.io/) to build.  After installing `gb`, simply:

```
gb build app/traincon
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

Very WIP.  Will eventually post diagrams and pictures.