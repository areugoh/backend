# Backend

## Data sources

- [ipinfo.io](https://ipinfo.io/): IP geolocation API to avoid asking the user for permission to access the location
- [kbgeo.com](https://www.kbgeo.com/): API to get the nearest sea/ocean from the given coordinates

## Setup

### Requirements

Environment variables:

- `MOCK`: if `true` the API will return mock data (the ML model generates real data, only INFOIP and KBGEO are mocked due to the API limits on the free plan)
- `DEBUG`: if `true` the API will print debug messages
- `INFOIP_TOKEN`: token for the [ipinfo.io](https://ipinfo.io/) API
- `KBGEO_TOKEN`: token for the [kbgeo.com](https://www.kbgeo.com/) API
- `KBGEO_URL`: url for the [kbgeo.com](https://www.kbgeo.com/) API

## Endpoints

The location will be returned as a JSON object except `/api/v2/animal/:name` with the following structure:

```json
"location": {
    "longitude": 0,
    "latitude": 0,
    "name": "name"
}
```

### GET /api/v1/area

Returns:
- the area name
- km from the closest sea/ocean
- avg condition of the surrounded water
- avg temperature of the surrounded water

For the given coordinates.

Response example:

```json
{
    "location": {
        "longitude": 35.4729335,
        "latitude": 139.6146366,
        "name": "Yokohama"
    },
    "nearest_aquatic_location": {
        "name": "Tokyo Bay",
        "condition": "good",
        "temperature": 17.8,
        "distance": 14.97
    }
}
```

| Condition | Description |
| --- | --- |
| good | The water quality is good |
| moderate | The water quality is moderate |
| contaminated | The water quality is contaminated |
| unknown | The water quality is unknown |
| radiation | The water quality is contaminated with radiation |
| danger | There are indigenous fauna or flora living in the proximities dangerous for humans |
| protected | There are indigenous fauna or flora living in the proximities protected by law |

- `temperature` is in Celsius degrees
- `distance` is in km
- `name` is the name of the closest sea/ocean, only available if `distance` is less than 100km

### GET /api/v1/graph

Returns:
- an array with the previous and future prediction of the water quality based on parameters like temperature, pH,pollution, etc.

For the given coordinates.

### GET /api/v1/animals

Returns:
- an array with the animals and plants that living in the surrounding water.

For the given coordinates.

### GET /api/v2/animal/:name [WIP for V2]

Returns:
- a description about the animal/plant
- a picture of the animal/plant
- the water quality that the animal/plant needs to survive
- human danger level
    


    Address: ftp.ptree.jaxa.jp
UID: laket79103_fesgrid.com
PW: SP+wari8