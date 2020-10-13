# Go demo repo
Go lang demo repo containing an API and tests.

## Requirements
Check the GOPATH environment variable so that the build passes. Set it via `go env -w GOPATH={your path}`

## Testing
The assertion package used comes from `stretchr/testify` https://github.com/stretchr/testify

To update Testify to the latest version, use `go get -u github.com/stretchr/testify`.

In order to run the tests, first local server needs to be started(this runs on port 9090 by default): `go run main`.
Running the actual tests can be done by executing the following command: `go test` from the root of the project.