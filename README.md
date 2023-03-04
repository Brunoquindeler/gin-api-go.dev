# Tutorial: Developing a RESTful API with Go and Gin

Ref.: https://go.dev/doc/tutorial/web-service-gin

## Endpoints
### /albums
    - GET: Get a list of all albums, returned as JSON.
    - POST: Add a new album from request data sent as JSON.
### /albums/:id
    - GET: Get an album by its ID, returning the album data as JSON.