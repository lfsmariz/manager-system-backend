package repository

import (
	"sync"

	"github.com/lfsmariz/manager-system-backend/app/repository/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once

type dbConn struct {
	Db *gorm.DB
}

var singleInstance *dbConn

func GetInstance() *dbConn {
	if singleInstance == nil {
		once.Do(
			func() {
				singleInstance = &dbConn{
					Db: connectRepository(),
				}
			})
	}

	return singleInstance
}

func connectRepository() *gorm.DB {

	dsn := "host=db user=testeuser password=testepass dbname=home port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// m := &model.Medicine{Animal: "Godofredo", MedicineName: "Rabica", ApplicationDate: time.Now(), NextApplicationDate: time.Now()}
	// n := &model.Medicine{Animal: "Alfredo", MedicineName: "Rabica", ApplicationDate: time.Now(), NextApplicationDate: time.Now()}
	// p := &model.Medicine{Animal: "Marlindo", MedicineName: "Rabica", ApplicationDate: time.Now(), NextApplicationDate: time.Now()}
	// q := &model.Medicine{Animal: "Safado", MedicineName: "Rabica", ApplicationDate: time.Now(), NextApplicationDate: time.Now()}

	db.AutoMigrate(&model.Medicine{})

	// db.Create(m)
	// db.Create(n)
	// db.Create(p)
	// db.Create(q)

	return db
}
