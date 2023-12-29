package database

import (
	log "github.com/sirupsen/logrus"
	productentity "go-template-project/module/product/entity"
	userentity "go-template-project/module/user/entity"
	"go-template-project/pkg"
	"gorm.io/gorm"
)

func createSeeds(db *gorm.DB) error {

	if data := db.Find(&userentity.Users{}); data.RowsAffected < 1 {

		UserAdmin := userentity.Users{
			Name:     "Super Admin",
			Email:    "super.admin@gmail.com",
			Password: pkg.HashPassword("12345678"),
		}

		UserAri := userentity.Users{
			Name:     "Malik",
			Email:    "malik@gmail.id",
			Password: pkg.HashPassword("12345678"),
		}
		db.Create(&UserAdmin)
		db.Create(&UserAri)
	}

	if data := db.Find(&productentity.Products{}); data.RowsAffected < 1 {
		db.Create(&productentity.Products{
			Name:     "Mouse",
			Quantity: 1000,
			Price:    10000})
		db.Create(&productentity.Products{
			Name:     "Keyboard",
			Quantity: 1000,
			Price:    10000})
		db.Create(&productentity.Products{
			Name:     "Laptop",
			Quantity: 1000,
			Price:    10000})
		db.Create(&productentity.Products{
			Name:     "Printer",
			Quantity: 1000,
			Price:    10000})
		db.Create(&productentity.Products{
			Name:     "Monitor",
			Quantity: 1000,
			Price:    10000})

		log.Println("✅ Seed Products inserted")
	}

	return nil
}
