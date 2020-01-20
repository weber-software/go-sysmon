package main

import (
	"time"
)

func IntervalIterator(interval int64) <-chan time.Time {
	ch := make(chan time.Time)
	go func() {
		for {
			now := time.Now().UnixNano()
			next := ((now / interval) + 1) * interval

			time.Sleep(time.Nanosecond * time.Duration(next-now))

			ch <- time.Unix(0, next)
		}
	}()
	return ch
}
