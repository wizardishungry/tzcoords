# tzcoords

Go library for getting rough lat/lon coords for a `time.Location` based on data from the IANA Time Zone Database.

[![Go Reference](https://pkg.go.dev/badge/github.com/WIZARDISHUNGRY/tzcoords.svg)](https://pkg.go.dev/github.com/WIZARDISHUNGRY/tzcoords)

## Uppdating

Place an updated `zone1970.tab` from the [IANA Time Zone Database](https://www.iana.org/time-zones) inside `./gen/`  and run `go generate ./...`.
