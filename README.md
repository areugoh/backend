# Backend

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
        "name": "Baltic Sea",
        "condition": "good",
        "temperature": 17.8,
        "distance": 14.97
    }
}
```

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