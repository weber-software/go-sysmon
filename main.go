package main

import (
	"go-sysmon/gatherer"
	"log"
	"os"
)

// /proc/meminfo
// /proc/interrupts
// /proc/vmstat
// /proc/uptime
// /proc/loadavg

func run(influx Influx, gatherers []gatherer.Gatherer) {
	tag := "sys1"

	for currentTime := range IntervalIterator(5 * 1000 * 1000 * 1000) {
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

	tag := os.Getenv("SYSMON_TAG")
	if tag == "" {
		log.Println("please provide env SYSMON_TAG")
		return
	}
	influxHost := os.Getenv("SYSMON_INFLUX_HOST")
	if influxHost == "" {
		log.Println("please provide env SYSMON_INFLUX_HOST")
		return
	}

	influx := Influx{
		Host: influxHost,
		Port: 8086,
		Db:   "go-sysmon",
	}

	log.Println(tag)
	run(influx, gatherers)
}
