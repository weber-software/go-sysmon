package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Influx struct {
	Host string
	Port int
	Db   string

	builder strings.Builder
}

func (influx *Influx) GetWriteUrl() string {
	return fmt.Sprintf("http://%s:%v/write?db=%s", influx.Host, strconv.Itoa(influx.Port), influx.Db)
}

func (influx *Influx) GetQueryUrl() string {
	return fmt.Sprintf("http://%s:%v/query?db=%s", influx.Host, strconv.Itoa(influx.Port), influx.Db)
}

func (influx *Influx) Add(measurement string, name string, field string, value float64, timestamp time.Time) {
	fmt.Fprintf(&influx.builder, "%s,name=%s %s=%f %d\n", measurement, name, field, value, timestamp.UnixNano())
}

func (influx *Influx) Post() error {
	_, err := http.Post(influx.GetWriteUrl(), "application/text", bytes.NewBufferString(influx.builder.String()))
	return err
}

func (influx *Influx) Query(query string) (string, error) {
	resp, err := http.Get(influx.GetQueryUrl() + "&q=" + url.QueryEscape(query))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func (influx *Influx) Clear() {
	influx.builder.Reset()
}
