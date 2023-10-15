# Backend

As a weather forecast app try to predict the water quality of the surrounding.

Our water quality app aims to help people to decide if it is safe to swim in the sea or ocean, if indigenous fauna or flora are in danger or if the water is contaminated with radiation.

If fauna or flora are in danger, the app will recommend to avoid the area or to be careful with the animals and plants living in the water giving information about them and how your day to day actions can affect them.

Our prediction is based on a ML model gathering data from JAXA (Japan Aerospace Exploration Agency), open source datasets about water quality and weather conditions combined with our custom algorithm to generate what we call "water score" or "condition".

## Data sources

- [ipinfo.io](https://ipinfo.io/): IP geolocation API to avoid asking the user for permission to access the location
- [kbgeo.com](https://www.kbgeo.com/): API to get the nearest sea/ocean from the given coordinates
- [eorc.jaxa.jp](https://www.eorc.jaxa.jp/ptree/LORA/index.html): API to get the water quality/pollution/temperature of the given coordinates
- OpenAPI: ML model to generate the description of the animals and plants from our data source

## Setup

### Requirements

Environment variables:

- `MOCK`: if `true` the API will return mock data (the ML model generates real data, only INFOIP and KBGEO are mocked due to the API limits on the free plan)
- `DEBUG`: if `true` the API will print debug messages
- `INFOIP_TOKEN`: token for the [ipinfo.io](https://ipinfo.io/) API
- `KBGEO_TOKEN`: token for the [kbgeo.com](https://www.kbgeo.com/) API
- `KBGEO_URL`: url for the [kbgeo.com](https://www.kbgeo.com/) API
- `OPENAI_API_KEY`: token for the [OpenAI](https://openai.com/) API
- `CHATGPT_URL`: url for the [OpenAI](https://openai.com/) API

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

> Cache for 24h

Returns:
- an array with the animals and plants that living in the surrounding water.

For the given coordinates.

Response example:

```json
{
    "location": {
        "longitude": 139.6146366,
        "latitude": 35.4729335,
        "name": "Yokohama"
    },
    "animals": [
        {
            "name": "Black-tailed gull",
            "in_danger": false,
            "description": "The Black-tailed gull is a common bird found in Yokohama. It is not considered to be in danger. However, human activities such as pollution and habitat destruction can negatively impact their nesting sites."
        },
        {
            "name": "Japanese spiny lobster",
            "in_danger": false,
            "description": "The Japanese spiny lobster is a species of lobster found in the waters around Yokohama. It is not considered to be in danger. However, overfishing and habitat destruction can pose a threat to their population."
        },
        {
            "name": "Japanese sea nettle",
            "in_danger": true,
            "description": "The Japanese sea nettle is a species of jellyfish found in the waters around Yokohama. It is considered to be in danger due to the impacts of pollution and climate change. These factors can disrupt their reproductive cycles and reduce their food sources."
        }
    ]
}
```

### GET /api/v2/animal/:name [WIP for V2]

Returns:
- a description about the animal/plant
- a picture of the animal/plant
- the water quality that the animal/plant needs to survive
- human danger level
