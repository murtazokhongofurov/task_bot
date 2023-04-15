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

    updates := h.bot.GetUpdatesChan(u)
    for update := range updates {
        go h.HandleBot(update)
    }
}

func (h *BotHandler) HandleBot(update tgbotapi.Update) {
    user, err := h.strg.GetOrCreate(update.Message.From.ID, update.Message.From.FirstName)
    if err != nil {
        h.SendMessage(user, "Error happened")
    }

    if update.Message.Command() == "start" {
        err = h.DisplayWelcome(user)
        if err != nil {
            log.Println("Error while welcome to bot: ", err.Error())
        }
    }
    if update.Message.Command() == "admin" {
       err = h.DisplayAdminPage(user)
       if err != nil {
            log.Println("Error /admin start: ", err.Error())
       }
    } else if update.Message.Text != "" {
        switch user.Step {
        case storage.ChangeRole:
            err = h.HandleEnterUsers(user, update.Message.Text)
            CheckError(err)
        default:
            h.SendMessage(user, errorMessage)
        }
    }
   
}

func (h *BotHandler) SendMessage(user *models.User, message string) {
	msg := tgbotapi.NewMessage(user.TgId, message)
	if _, err := h.bot.Send(msg); err != nil {
		log.Println(err)
	}
}

func CheckError(err error) {
    if err != nil {
        log.Println(err)
    }
}