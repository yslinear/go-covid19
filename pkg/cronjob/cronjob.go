package cronjob

import (
	"yslinear/go-covid19/pkg/dataset"

	"github.com/robfig/cron/v3"
)

func Setup() {
	c := cron.New()
	c.AddFunc("@every 20m", dataset.Setup)
	c.Start()
}
