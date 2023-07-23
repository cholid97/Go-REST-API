package seed

import (
	"log"

	"example/web-service-gin/api/models"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "Ahmad Sahroni",
		Email:    "sahroni@gmail.com",
		Password: "as123",
	},
	models.User{
		Nickname: "King Azis",
		Email:    "kingslanding@gmail.com",
		Password: "kings1232",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Lorem Ipsum",
		Content: "Dolor sit amet",
	},
	models.Post{
		Title:   "Lorem Lore",
		Content: "Dolor Amet sit",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
