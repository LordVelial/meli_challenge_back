package db

import (
	"database/sql"
	"fmt"
	"meli_challenge_back/internal/entity"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Función para obtener la conexión a la base de datos PostgreSQL
func GetConnection(cfgDb entity.DataBase) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfgDb.DbHost, cfgDb.DbPort, cfgDb.DbUser, cfgDb.DbPass, cfgDb.DbName)

	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Función para cerrar la conexión a la base de datos PostgreSQL
func CloseConnection() error {
	if db != nil {
		err := db.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
