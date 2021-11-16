package services

import (
	"github.com/go-co-op/gocron"
	"github.com/rwcoding/mrng/config"
	"time"
)

func SyncTimer() {
	sp := config.GetTimer()
	if sp <= 0 {
		sp = 3600
	}

	s := gocron.NewScheduler(time.UTC)
	s.Every(sp).Seconds().Do(func() {
		Sync()
	})
	s.StartAsync()
}
