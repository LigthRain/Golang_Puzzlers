package main

import "time"

type rateLimter struct {
	timer *time.Ticker
}

func NewRateLimter(count int, duration time.Duration) *rateLimter {
	return &rateLimter{
		timer: time.NewTicker(duration / time.Duration(count)),
	}
}

func (r rateLimter) Limit() {
	<- r.timer.C
}

func main() {
	
}
