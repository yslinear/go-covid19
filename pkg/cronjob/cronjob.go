package cronjob

import (
	"yslinear/go-covid19/pkg/dataset"

	"github.com/robfig/cron/v3"
)

func Setup() {
	c := cron.New()
	c.AddFunc("*/20 7-22 * * *", dataset.Setup)
	c.Start()
}
