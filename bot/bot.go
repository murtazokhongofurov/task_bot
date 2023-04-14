package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.com/task_bot/config"
	"gitlab.com/task_bot/storage"
	"gitlab.com/task_bot/storage/models"
)


type BotHandler struct {
    cfg config.Config
    strg storage.StorageI
    bot     *tgbotapi.BotAPI
}


func New(cfg config.Config, strg storage.StorageI) BotHandler {
    bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true
    return BotHandler{
        cfg: cfg,
        strg: strg,
        bot: bot,
    }
}

func (h *BotHandler) Start() {
    log.Printf("Authorized on account %s", h.bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)

    u.Timeout = 60

    _ = h.bot.GetUpdatesChan(u)

    // for update := range updates {
    //     go 
    // }

}

func (h *BotHandler) HandleBot(update tgbotapi.Update) {
    user, err := h.strg.GetOrCreate(update.Message.From.ID, update.Message.From.FirstName)
    if err != nil {
        h.SendMessage(user, "Error happened")
    }

    if update.Message.Command() == "start" {
        
    }
}


func (h *BotHandler) SendMessage(user *models.User, message string) {
	msg := tgbotapi.NewMessage(user.TgId, message)
	if _, err := h.bot.Send(msg); err != nil {
		log.Println(err)
	}
}

// func main() {

//     bot, err := tgbotapi.NewBotAPI("5707916239:AAGJoPr07ddUYl6_Osvbb5A3XtGYHZphPD8")
//     if err != nil {
//         log.Fatal(err)
//     }

//     bot.Debug = true

//     log.Printf("Authorized on account %s", bot.Self.UserName)
//     u := tgbotapi.NewUpdate(0)
//     u.Timeout = 60

//     updates, err := bot.GetUpdatesChan(u)
// 	if err != nil {
// 		log.Println("Error while GetUpdateChan(u)", err.Error())
// 	}

//     for update := range updates {
// 		fmt.Println("Message ====>>>> \n", update.Message.Chat.FirstName)
//         if update.Message == nil {
//             continue
//         }

//         msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You said: " + update.Message.Text)
//         bot.Send(msg)
//     }
// }
