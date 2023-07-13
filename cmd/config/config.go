package config

import (
	"fmt"
	"meli_challenge_back/internal/entity"
)

func GetConfig(args []string) (dbEntity entity.DataBase, err error) {
	if len(args) < 6 {
		return entity.DataBase{}, fmt.Errorf("Usage: go run cmd/main.go <dbname> <dbpass> <dbuser> <dbhost> <dbport>")
	}

	dbEntity = entity.DataBase{
		DbName: args[1],
		DbPass: args[2],
		DbUser: args[3],
		DbHost: args[4],
		DbPort: args[5],
	}

	return dbEntity, nil
}
