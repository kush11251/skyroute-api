# Skyroute API
The Skyroute API is a production-ready system designed to optimize flight routes. It provides a RESTful interface for creating, reading, updating, and deleting flight routes.
## Getting Started
1. Initialize the project by running `go mod init skyroute-api` in the root directory.
2. Install dependencies by running `go get`.
3. Build and run the application by running `go build` and `./skyroute-api`.
## API Endpoints
* `POST /routes`: Create a new flight route
* `GET /routes`: Retrieve all flight routes
* `GET /routes/{id}`: Retrieve a single flight route by ID
* `PUT /routes/{id}`: Update a single flight route by ID
* `DELETE /routes/{id}`: Delete a single flight route by ID