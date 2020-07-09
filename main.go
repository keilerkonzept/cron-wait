package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron"
)

var config struct {
	cronExpressions    []string
	schedules          []cron.Schedule
	evaluationInterval time.Duration
}

var app = "wait-for-cron-expression-match"

func init() {
	log.SetFlags(log.Ltime | log.Ldate | log.Lmicroseconds)
	log.SetPrefix(fmt.Sprintf("[%s] ", app))
	flag.Parse()

	config.evaluationInterval = 250 * time.Millisecond
	config.cronExpressions = flag.Args()

	if len(config.cronExpressions) == 0 {
		log.Fatal("at least one expression is required")
	}
	for _, e := range config.cronExpressions {
		schedule, err := cron.Parse(e)
		if err != nil {
			log.Fatalf("parse cron expression %q: %v", e, err)
		}
		config.schedules = append(config.schedules, schedule)
	}
}

func main() {
	now := time.Now()
	tickInterval := config.evaluationInterval
	var next time.Time
	for i, s := range config.schedules {
		nextNew := s.Next(now)
		if i == 0 || nextNew.Before(next) {
			next = nextNew
		}
	}
	delta := next.Sub(now)
	if delta < tickInterval {
		tickInterval = delta
	}
	tick := time.Tick(tickInterval)

	plural := ""
	if len(config.cronExpressions) > 1 {
		plural = "s"
	}
	log.Printf("waiting %v until next match (%v) of cron expression%s %q", delta, next.Format(time.RFC3339Nano), plural, config.cronExpressions)
	defer log.Print("done")
	for now = range tick {
		if now.After(next) {
			return
		}
	}
}
