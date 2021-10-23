package events

import "go-echo/pkg/db"

func StartApp() {
	db.ConnectMongo()
	db.ConnectRedis()
}

func StopApp() {
	db.CloseMongo()
	db.ConnectRedis()
}
