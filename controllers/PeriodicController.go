package controllers

import (
	"log"
	"os"
	"strconv"
	"time"
)

func PeriodCalls() {
	log.Println("Inside PeriodicCalls ")
	p := os.Getenv("PERIOD")
	t, _ := strconv.Atoi(p)
	ticker := time.NewTicker(time.Duration(t) * time.Millisecond)
	go func() {
		for _ = range ticker.C {
			log.Println("Calling GetCatFacts API")
			GetCatFacts()
		}
	}()
	select {}
}
