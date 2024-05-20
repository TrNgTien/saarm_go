package helpers

import (
	"fmt"
	"saarm/pkg/utilities"

	"github.com/robfig/cron/v3"
)

func InitCron() error {
	c := cron.New()

	fmt.Println("START CRON JOB.......")

	// Run cronjob every midnight
	c.AddFunc("0 0 * * *", func() {
		fmt.Println("START | Clean resource every midnight........")

		if err := utilities.RemoveAllAssets(); err != nil {
			fmt.Println("[InitCron] Remove resources failed!!")
		}

		fmt.Println("Finish clean!!")

	})

	c.Start()

	// keep job run until server is off
	// var forever chan struct{}
	// <-forever
	return nil
}
