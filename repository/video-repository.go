package repository

import (
	"github.com/sheryarbutt/Learn-Go-Gin/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video *entity.Video)
	Update(video *entity.Video)
	Delete(video *entity.Video)
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() *database {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&entity.Video{}, &entity.Person{})

	return &database{
		connection: db,
	}
}

func (db *database) Save(video *entity.Video) {
	db.connection.Create(video)
}

func (db *database) Update(video *entity.Video) {
	db.connection.Save(video)
}

func (db *database) Delete(video *entity.Video) {
	db.connection.Set("gorm:auto_preload", true).Delete(video)
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}

func (db *database) CloseDB() {
	DB, err := db.connection.DB()
	if err != nil {
		panic("Failed to close database")
	}
	DB.Close()
}
