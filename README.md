# go-sysmon
a minimal tool for writing system metrics to an influxdb - written in golang

## Environment variables
- `SYSMON_INTERVAL_MS` the interval for metrics collection, defaults to `5000`
- `SYSMON_INFLUX_HOST` the influx host to write to
- `SYSMON_INFLUX_DB` the dateabase to write to, defaults to `go-sysmon`
- `SYSMON_INFLUX_TAG` the tag value for `name` in the timeseries

## Usage

`SYSMON_INFLUX_TAG=system1 SYSMON_INFLUX_HOST=192.168.1.60 ./go-sysmon`

## Docker

`docker build -t go-sysmon .`
`docker run -e SYSMON_INFLUX_TAG=sys1 -e SYSMON_INFLUX_HOST=192.168.1.60 -it --rm go-sysmon`