package cron

import (
	"chat/services/repository"
	"github.com/go-co-op/gocron"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func CronStatus(repo repository.UserRepository) {
	s := gocron.NewScheduler(time.UTC)

	s.Every(10).Seconds().Do(func() {
		//fmt.Println("hello world")
		listUser, err := repo.GetAll(bson.M{"status": 1})
		//fmt.Println("len user status", len(listUser))
		if err != nil {
			//fmt.Println("Cannot check status")
		} else {
			for i := range listUser {
				if time.Now().UTC().Unix()-listUser[i].RequestAt.Unix() > 10 {
					query := primitive.M{"_id": listUser[i].ID}
					err := repo.UpdateStatus(query, 0)
					if err != nil {
						//fmt.Println("Cannot update status for ", listUser[i].ID)
					} else {
						//fmt.Println("Update successful for", listUser[i].ID)
					}
				}
			}
		}
	})
	s.StartBlocking()
}
