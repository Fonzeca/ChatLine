package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Sirve para obener el objeto para interactuar con la base de datos
func ObtenerConexionDb() (*gorm.DB, func() error, error) {
	//Cambiarlo por Viper
	host := os.Getenv("CHATLINE_DB_HOST")
	if host == "" {
		host = "vps-1791261-x.dattaweb.com:3306"
	}

	user := os.Getenv("chatlineDbUser")
	if user == "" {
		user = "root"
	}
	pass := os.Getenv("chatlineDbPass")
	if pass == "" {
		pass = "almacen.C12"
	}

	dsn := user + ":" + pass + "@tcp(" + host + ")/chatline?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, nil, err
	}
	sqlDb, _ := db.DB()
	return db, sqlDb.Close, nil
}
