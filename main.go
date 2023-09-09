package main

import (
	"os"
	"time"

	"github.com/dungnguyen/bank-transfer/infrastructure"
	"github.com/dungnguyen/bank-transfer/infrastructure/database"
	"github.com/dungnguyen/bank-transfer/infrastructure/log"
	"github.com/dungnguyen/bank-transfer/infrastructure/router"
	"github.com/dungnguyen/bank-transfer/infrastructure/validation"
)

func main() {
	var app = infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(log.InstanceLogrusLogger).
		Validator(validation.InstanceGoPlayground).
		DBSQL(database.InstancePostgres).
		DbNoSQL(database.InstanceMongoDB)

	app.WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGorillaMux).
		Start()
}
