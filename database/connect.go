package database

import (
	"context"
	"fmt"
	"log"

	"pentag.kr/api-server/configs"
	"pentag.kr/api-server/ent"

	_ "github.com/lib/pq"
)

var (
	dbClient *ent.Client
)

func GetDBClient() *ent.Client {
	return dbClient
}

func NewDBClient() {

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		configs.Env.DBHost, configs.Env.DBPort, configs.Env.DBUser, configs.Env.DBName, configs.Env.DBPassword)

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}
	dbClient = client
	log.Println("Database connected")
}