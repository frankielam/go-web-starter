package infra

import (
	"gowebstarter/infra/config"
	"log"
	"sync"
	"time"
)

const TimeZone = "Asia/Shanghai"

var timezoneOnce sync.Once

func Init() {
	initTimeZone()
	config.GetConf()
}

func initTimeZone() {
	timezoneOnce.Do(func() {
		ll, err := time.LoadLocation(TimeZone)
		if err != nil {
			log.Printf("Failed to load timezone %q: %v", TimeZone, err)
			time.Local = time.UTC
		} else {
			time.Local = ll
			log.Printf("Set timezone to %q", TimeZone)
		}
	})
}
