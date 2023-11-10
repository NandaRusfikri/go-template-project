package database

import (
	log "github.com/sirupsen/logrus"
	auth_entity "go-template-project/module/auth/entity"
	item_entity "go-template-project/module/product/entity"
	user_entity "go-template-project/module/user/entity"
	"go-template-project/pkg"
	"gorm.io/gorm"
)

func MigrateDBSQL(db *gorm.DB) error {
	err := db.AutoMigrate(
		&user_entity.EntityUser{},
		&auth_entity.EntityForgotPassword{},
		&item_entity.MSProduct{},
	)

	if err != nil {
		log.Errorln("❌ Error Migrate ", err.Error())
		return err
	}
	if data := db.Find(&user_entity.EntityUser{}); data.RowsAffected < 1 {

		UserAdmin := user_entity.EntityUser{
			Name:     "Super Admin",
			Email:    "super.admin@gmail.com",
			Password: pkg.HashPassword("12345678"),
		}

		UserAri := user_entity.EntityUser{
			Name:     "Malik",
			Email:    "malik@gmail.id",
			Password: pkg.HashPassword("12345678"),
		}
		db.Create(&UserAdmin)
		db.Create(&UserAri)
	}

	if data := db.Find(&item_entity.MSProduct{}); data.RowsAffected < 1 {
		db.Create(&item_entity.MSProduct{
			Name:     "Mouse",
			Quantity: 1000,
			Price:    10000})
		db.Create(&item_entity.MSProduct{
			Name:     "Keyboard",
			Quantity: 1000,
			Price:    10000})
		db.Create(&item_entity.MSProduct{
			Name:     "Laptop",
			Quantity: 1000,
			Price:    10000})
		db.Create(&item_entity.MSProduct{
			Name:     "Printer",
			Quantity: 1000,
			Price:    10000})
		db.Create(&item_entity.MSProduct{
			Name:     "Monitor",
			Quantity: 1000,
			Price:    10000})

		log.Println("✅ Seed MSProduct inserted")
	}

	return err
}
