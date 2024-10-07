package cron

import (
	"test-fiber/lib"

	"github.com/robfig/cron/v3"
)

func Init() {
	c := cron.New(cron.WithSeconds())

	// 1초마다 실행
	c.AddFunc("* * * * * *", func() {
		lib.Info("Hello, World!")
	})

	c.Start()
}
