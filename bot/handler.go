package bot

import (
	"bytes"
	"log"
	"strconv"

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


func (h *BotHandler) HandleGetUsers(user *models.User, text string) error {
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
	k, number := 1, ""
	for i, usertgname := range info.Users {
		if i > 0 {
			buf.WriteString("\n")
		}
		number = strconv.Itoa(k)
		buf.WriteString(number +"." + usertgname.TgName)
		k++
		userNameString := buf.String()
		if len(info.Users) == i+1 {
			msg := tgbotapi.NewMessage(user.TgId, userNameString)
			if _, err := h.bot.Send(msg); err != nil {
				return err
			}
		}
	}
	return nil
}

