package helpers

import (
	"fmt"
	"saarm/pkg/utilities"

	"github.com/robfig/cron/v3"
)

func InitCron() {
	c := cron.New()

	fmt.Println("START CRON JOB.......")
	c.AddFunc("@midnight", func() {
		fmt.Println("START | Clean resource every midnight........")

		if err := utilities.RemoveAllAssets(); err != nil {
			fmt.Println("[InitCron] Remove resources failed!!")
		}

		fmt.Println("Finish clean!!")

	})

	c.Start()
}
