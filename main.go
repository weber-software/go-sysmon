package main

import (
	"go-sysmon/gatherer"
	"log"
	"os"
	"strconv"
)

// /proc/meminfo
// /proc/interrupts
// /proc/vmstat
// /proc/uptime
// /proc/loadavg

func run(influx Influx, gatherers []gatherer.Gatherer, interval int64, tag string) {
	for currentTime := range IntervalIterator(interval * 1000 * 1000) {
		total := 0
		// log.Println(currentTime)
		for _, g := range gatherers {
			points, err := g.Read()

			if err != nil {
				log.Println(err)
				return
			}

			for _, point := range points {
				influx.Add(g.GetName(), tag, point.Name, point.Value, currentTime)
			}
			total += len(points)

			// log.Println(points)
		}
		err := influx.Post()
		influx.Clear()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("posted %v points", total)
	}
}

func main() {
	memInfo := gatherer.MemInfo{}
	uptime := gatherer.Uptime{}
	vmstat := gatherer.VmStat{}
	loadavg := gatherer.LoadAvg{}
	gatherers := []gatherer.Gatherer{&memInfo, &uptime, &vmstat, &loadavg}

	tag := os.Getenv("SYSMON_INFLUX_TAG")
	if tag == "" {
		log.Println("please provide env SYSMON_INFLUX_TAG")
		os.Exit(1)
		return
	}
	influxHost := os.Getenv("SYSMON_INFLUX_HOST")
	if influxHost == "" {
		log.Println("please provide env SYSMON_INFLUX_HOST")
		os.Exit(1)
		return
	}

	intervalMs := int64(5000)

	intervalStr := os.Getenv("SYSMON_INTERVAL_MS")
	if intervalStr != "" {
		value, err := strconv.ParseInt(intervalStr, 10, 64)
		if err != nil {
			log.Printf("failed to parse interval: %v", err)
			os.Exit(1)
			return
		}
		intervalMs = value
	}

	influx := Influx{
		Host: influxHost,
		Port: 8086,
		Db:   "go-sysmon",
	}

	influxDb := os.Getenv("SYSMON_INFLUX_DB")
	if influxDb != "" {
		influx.Db = influxDb
	}

	log.Printf("gathering system metrics all %vms and writing them to db \"%v\" on host \"%v\" with tag \"%v\"\n", intervalMs, influx.Db, influx.Host, tag)
	run(influx, gatherers, intervalMs, tag)
}
