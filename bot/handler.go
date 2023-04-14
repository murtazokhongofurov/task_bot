package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.com/task_bot/storage"
	"gitlab.com/task_bot/storage/models"
)


func (h *BotHandler)  DisplayEnterFullNameMenu(user *models.User) error {
	err := h.strg.ChangeStep(user.TgId, storage.EnterFullnameStep)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(user.TgId, "Ism Familiyangizni kiriting ⬇️")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := h.bot.Send(msg); err != nil {
		log.Println(err)
	}
	return nil
}