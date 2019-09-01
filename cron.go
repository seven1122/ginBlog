package main

/*
import (
	"github.com/robfig/cron"
	"seven1122/ginBlog/models"
	"seven1122/ginBlog/pkg/logging"
	"time"
)

func main()  {
	logging.Info("cron is starting !!!")
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		logging.Info("clearAllTag ing !")
		models.ClearAllTag()
	})

	c.AddFunc("* * * * * *", func() {
		logging.Info("clearAllArticle ing !")
		models.ClearAllAriticle()
	})

	c.Start()

	t := time.NewTimer(time.Second * 10)
	for{
		select {
		case <- t.C:
			t.Reset(time.Second * 10)

		}
	}
}
*/
