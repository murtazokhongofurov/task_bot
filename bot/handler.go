package bot

import (
	"bytes"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.com/task_bot/storage"
	"gitlab.com/task_bot/storage/models"
)

const page = 1
const limit = 10

func (h *BotHandler)  DisplayWelcome(user *models.User) error {
	err := h.strg.ChangeStep(user.TgId, storage.EnterStartingStep)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(user.TgId, startMessage)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := h.bot.Send(msg); err != nil {
		log.Println(err)
	}
	return nil
}

func (h *BotHandler) DisplayAdminPage(user *models.User) error {
	err := h.strg.ChangeStep(user.TgId, storage.ChangeRole)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(user.TgId, adminStartMessage)
	msg.ReplyMarkup = adminMenuKeyboards
	if _, err := h.bot.Send(msg); err != nil {
		log.Println(err)
	}
	return nil
}


func (h *BotHandler) HandleEnterUsers(user *models.User, text string) error {
	if text == adminGetUsers {
		return h.DisplayAllUsers(user)
	}
	return h.DisplayAdminPage(user)
}

func (h *BotHandler) DisplayAllUsers(user *models.User) error {
	info, err := h.strg.GetAllUsers(page, limit)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	for i, usertgname := range info.Users {
		if i > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(usertgname.TgName)
		userNameString := buf.String()
		msg := tgbotapi.NewMessage(user.TgId, userNameString)
		if _, err := h.bot.Send(msg); err != nil {
			return err
		}
	}
	return nil
}

