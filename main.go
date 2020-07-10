package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	cron "github.com/robfig/cron/v3"
)

var config struct {
	Quiet     bool
	PrintDots bool

	cronExpressions    []string
	schedules          []cron.Schedule
	evaluationInterval time.Duration
}

var app = "wait-for-cron-expression-match"

func init() {
	log.SetFlags(log.Ltime | log.Ldate | log.Lmicroseconds)
	log.SetPrefix(fmt.Sprintf("[%s] ", app))
	flag.BoolVar(&config.Quiet, "quiet", config.Quiet, "Suppress all output")
	flag.BoolVar(&config.Quiet, "q", config.Quiet, "(alias for -quiet)")
	flag.BoolVar(&config.PrintDots, "dots", config.PrintDots, "Print dots to stdout while waiting")
	flag.Parse()

	config.evaluationInterval = 1 * time.Second
	config.cronExpressions = flag.Args()

	if len(config.cronExpressions) == 0 {
		log.Fatal("at least one expression is required")
	}
	parser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	for _, e := range config.cronExpressions {
		schedule, err := parser.Parse(e)
		if err != nil {
			log.Fatalf("parse cron expression %q: %v", e, err)
		}
		config.schedules = append(config.schedules, schedule)
	}

	if config.Quiet {
		log.SetOutput(ioutil.Discard)
		config.PrintDots = false
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
		if config.PrintDots {
			fmt.Print(".")
		}
		if now.After(next) {
			if config.PrintDots {
				fmt.Println()
			}
			return
		}
	}
}
