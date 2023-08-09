package migration

import (
	"fmt"
	"test-absensi/models"
	"test-absensi/pkg/mysql"
)

func RunAutoMigrate() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Migration")
}