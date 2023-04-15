package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/task_bot/bot"
	"gitlab.com/task_bot/config"
	"gitlab.com/task_bot/storage"
)
func main() {
	fmt.Println("Server is running")
	cfg := config.Load()
	db, err := sql.Open("sqlite3", cfg.SqliteUrl)
	if err != nil {
		log.Println("Error connection sqlite3", err.Error())
	}

	statment, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY, 
		tg_id INTEGER, 
		tg_name VARCHAR(100), 
		step VARCHAR(100),
		created_at NOT NULL DEFAULT CURRENT_TIMESTAMP)`)

	if err != nil {
		log.Println("Error while creating table", err.Error())
	}
	statment.Exec()
	strg := storage.NewStoragePg(db)

	botHandler := bot.New(cfg, strg)

	botHandler.Start()
}