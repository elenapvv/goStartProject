package utils

import (
	"app/config"
	log "app/logging"
	"fmt"
	"github.com/tarantool/go-tarantool"
)

var Conn *tarantool.Connection

func InitDB(env string) {
	configuration := config.GetConfig(env)

	var err error
	Conn, err = tarantool.Connect(fmt.Sprintf("%s:%s", configuration.DbHost, configuration.DbPort),
		tarantool.Opts{
			User: configuration.DbUsername,
			Pass: configuration.DbPassword,
		})

	if err != nil {
		log.ErrorLogger.Println("Connection refused")
	} else {
		log.InfoLogger.Println("Connected to DB")
	}

	//defer func(Conn *tarantool.Connection) {
	//	err := Conn.Close()
	//	if err != nil {
	//		log.ErrorLogger.Println("Не удалось закрыть соединение с БД")
	//	}
	//}(Conn)
}
