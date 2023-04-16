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
const limit = 100

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
	err := h.strg.ChangeStep(user.TgId, storage.AdminRole)
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
		return h.DisplayGetAllUsers(user)
	}
	return nil
}

func (h *BotHandler) DisplayGetAllUsers(user *models.User) error {
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

func (h *BotHandler) HandleMessage(user *models.User, text string) error {
	if text == sendMessage {
		return h.DisplayChangeStep(user)
	}
	return nil
}

func (h *BotHandler) DisplayChangeStep(user *models.User) error {
	err := h.strg.ChangeStep(user.TgId, storage.SendMessage)
	if err != nil {
		return err
	}
	return h.DisplayMessage(user)
}

func (h *BotHandler) DisplayMessage(user *models.User) error {
	err := h.strg.ChangeStep(user.TgId, storage.SendMessage)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(user.TgId, enterMessage)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := h.bot.Send(msg); err != nil {
		log.Println(err)
	}
	return nil
}

func (h *BotHandler) HandleSendMessage(user *models.User, text string) error {
	err := h.strg.ChangeStep(user.TgId, storage.StatusStep)
	if err != nil {
		return err
	}

	err = h.ReceiveMessageToQueue(text, "Receive message")
	if err != nil {
		log.Println("Info is not found", err.Error())
	}
	ids, err := h.strg.GetAllTgIds()
	if err != nil {
		return err
	}
	message, err := h.GetMessageFromQueue("Receive message")
	if err != nil {
		return err
	}
	for _, id := range ids.TgIds {
		msg := tgbotapi.NewMessage(int64(id.TgId), message)
		if _, err := h.bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
	return nil
}

