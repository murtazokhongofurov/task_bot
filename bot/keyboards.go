package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
var adminMenuKeyboards = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(adminGetUsers),
		tgbotapi.NewKeyboardButton(messageStatus),
		tgbotapi.NewKeyboardButton(sendMessage),
	),
)