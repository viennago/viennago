package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	r := time.Second
	l := rate.NewLimiter(rate.Every(r), 2)

	var count int
	for {
		if l.Allow() {
			count++
			log.WithFields(logrus.Fields{
				"count": count,
			}).Info("I may speak!")
		} else {
			log.WithFields(logrus.Fields{
				"time": r * 2,
			}).Debug("I waited silently")
			time.Sleep(r * 2)
		}
	}
}
