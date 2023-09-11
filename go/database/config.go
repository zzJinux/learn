package database

import (
	"os"
	"strconv"
)

var EnvPort uint16

func init() {
	envPort1 := os.Getenv("DB_PORT")
	if envPort1 == "" {
		EnvPort = 3306
	} else {
		v, err := strconv.Atoi(envPort1)
		if err != nil {
			panic(err)
		}
		EnvPort = uint16(v)
	}
}
