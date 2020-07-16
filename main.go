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
	Quiet                     bool
	PrintDots                 bool
	PrintNextTimestampAndExit bool
	PrintDeltaAndExit         bool
	TimestampFormat           string
	CronExpressions           []string

	schedules          []cron.Schedule
	evaluationInterval time.Duration
}

var app = "cron-wait"

func init() {
	log.SetFlags(log.Ltime | log.Ldate | log.Lmicroseconds)
	log.SetPrefix(fmt.Sprintf("[%s] ", app))
	config.TimestampFormat = time.RFC3339
	flag.BoolVar(&config.PrintDeltaAndExit, "print-delta-and-exit", config.PrintDeltaAndExit, "Only print the duration (in seconds) until the next expression match and exit (without waiting)")
	flag.BoolVar(&config.PrintNextTimestampAndExit, "print-next-match-and-exit", config.PrintNextTimestampAndExit, "Only print the timestamp of the next expression match and exit (without waiting)")
	flag.StringVar(&config.TimestampFormat, "format", config.TimestampFormat, "Timestamp format")
	flag.BoolVar(&config.Quiet, "quiet", config.Quiet, "Suppress all output")
	flag.BoolVar(&config.Quiet, "q", config.Quiet, "(alias for -quiet)")
	flag.BoolVar(&config.PrintDots, "dots", config.PrintDots, "Print dots to stdout while waiting")
	flag.Parse()

	config.evaluationInterval = 1 * time.Second
	config.CronExpressions = flag.Args()

	if len(config.CronExpressions) == 0 {
		log.Fatal("at least one expression is required")
	}
	parser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	for _, e := range config.CronExpressions {
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
	next := nextMatch(now)
	if config.PrintDeltaAndExit {
		fmt.Println(time.Until(next).Seconds())
		return
	}
	if config.PrintNextTimestampAndExit {
		fmt.Println(next.Format(config.TimestampFormat))
		return
	}
	waitUntil(next)
}

func waitUntil(next time.Time) {
	defer func() {
		if config.PrintDots {
			fmt.Println()
		}
		log.Print("done")
	}()

	tickInterval := config.evaluationInterval
	delta := time.Until(next)
	if delta < tickInterval {
		tickInterval = delta
	}
	tick := time.Tick(tickInterval)

	plural := ""
	if len(config.CronExpressions) > 1 {
		plural = "s"
	}
	log.Printf("waiting %v until next match (%v) of cron expression%s %q", time.Until(next), next.Format(config.TimestampFormat), plural, config.CronExpressions)

	for now := range tick {
		if config.PrintDots {
			fmt.Print(".")
		}
		if now.After(next) {
			return
		}
	}
}

func nextMatch(now time.Time) time.Time {
	var next time.Time
	for i, s := range config.schedules {
		nextNew := s.Next(now)
		if i == 0 || nextNew.Before(next) {
			next = nextNew
		}
	}
	return next
}
