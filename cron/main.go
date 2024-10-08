package cron

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/robfig/cron/v3"

	"test-fiber/config"
)

func Init() {
	c := cron.New(cron.WithSeconds())

	// 1초마다 실행
	c.AddFunc("* * * * * *", func() {
		log.Info("Hello, World! env: ", config.AppConfig.GO_ENV)
	})

	c.Start()
}
