package main

import (
	"kitchenpager/internal/apichecker"
	"kitchenpager/internal/config"
	"kitchenpager/pkg/dapnet"
	"log"
	"time"
)

func main() {
	c, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	ticker := time.NewTicker(1 * time.Minute)
	pager := dapnet.NewPager(c.Username, c.Password)
	go pager.Start()
	apichecker.CheckandPageOpenStatusperiodically(ticker, pager, c.SpaceAPI, c.Callsigns)
}
