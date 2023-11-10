package main

import cron "github.com/robfig/cron/v3"

func init() {
	Cron.Start()
}

var Cron = cron.New()
