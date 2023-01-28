package main

import (
	log "app/logging"
	"app/routers"
	"app/utils"
	"github.com/tarantool/go-tarantool"
)

func main() {
	env := "dev" // or "prod"

	log.InitLogging()

	utils.InitDB(env)
	defer func(Conn *tarantool.Connection) {
		err := Conn.Close()
		if err != nil {
			log.ErrorLogger.Println("Не удалось закрыть БД")
		}
	}(utils.Conn)

	routers.InitRouters()
}
