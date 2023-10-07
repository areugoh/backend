# Backend

## Endpoints

### GET /api/v1/area/:long/:lat

Returns:
- the area name
- km from the closest sea/ocean
- avg condition of the surrounded water
- avg temperature of the surrounded water

For the given coordinates.

### GET /api/v1/graph/:long/:lat/

Returns:
- an array with the previous and future prediction of the water quality based on parameters like temperature, pH,pollution, etc.

For the given coordinates.

### GET /api/v1/animals/:long/:lat/

Returns:
- an array with the animals and plants that living in the surrounding water.

For the given coordinates.

### GET /api/v1/animal/:name

Returns:
- a description about the animal/plant
- a picture of the animal/plant
- the water quality that the animal/plant needs to survive
- human danger level