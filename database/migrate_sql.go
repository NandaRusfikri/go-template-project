package database

import (
	log "github.com/sirupsen/logrus"
	authentity "go-template-project/module/auth/entity"
	productentity "go-template-project/module/product/entity"
	userentity "go-template-project/module/user/entity"
	"go-template-project/pkg"
	"gorm.io/gorm"
)

func MigrateDBSQL(db *gorm.DB) error {
	err := db.AutoMigrate(
		&userentity.EntityUser{},
		&authentity.EntityForgotPassword{},
		&productentity.MSProduct{},
	)

	if err != nil {
		log.Errorln("❌ Error Migrate ", err.Error())
		return err
	}
	if data := db.Find(&userentity.EntityUser{}); data.RowsAffected < 1 {

		UserAdmin := userentity.EntityUser{
			Name:     "Super Admin",
			Email:    "super.admin@gmail.com",
			Password: pkg.HashPassword("12345678"),
		}

		UserAri := userentity.EntityUser{
			Name:     "Malik",
			Email:    "malik@gmail.id",
			Password: pkg.HashPassword("12345678"),
		}
		db.Create(&UserAdmin)
		db.Create(&UserAri)
	}

	if data := db.Find(&productentity.MSProduct{}); data.RowsAffected < 1 {
		db.Create(&productentity.MSProduct{
			Name:     "Mouse",
			Quantity: 1000,
			Price:    10000})
		db.Create(&productentity.MSProduct{
			Name:     "Keyboard",
			Quantity: 1000,
			Price:    10000})
		db.Create(&productentity.MSProduct{
			Name:     "Laptop",
			Quantity: 1000,
			Price:    10000})
		db.Create(&productentity.MSProduct{
			Name:     "Printer",
			Quantity: 1000,
			Price:    10000})
		db.Create(&productentity.MSProduct{
			Name:     "Monitor",
			Quantity: 1000,
			Price:    10000})

		log.Println("✅ Seed MSProduct inserted")
	}

	return err
}
