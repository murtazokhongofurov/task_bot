package cmd

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/task_bot/bot"
	"gitlab.com/task_bot/config"
	"gitlab.com/task_bot/storage"
)
func main(){
	cfg := config.Load(".")

	sqliteConn, err := sql.Open("sqlite3", cfg.SqliteUrl)
	if err != nil {
		log.Println("Error connection sqlite3", err.Error())
		return
	}
	defer sqliteConn.Close()

	strg := storage.NewStoragePg(sqliteConn)

	botHandler := bot.New(cfg, strg)

	botHandler.Start()
}