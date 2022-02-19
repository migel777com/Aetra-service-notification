package tgbot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleCommand(update tgbotapi.Update, server Server) {

	chatID := update.Message.Chat.ID
	username := update.Message.Chat.UserName
	message := update.Message.Text
	contact := ""

	if update.Message.Contact != nil {
		contact = "+" + update.Message.Contact.PhoneNumber
	}


	fmt.Println(chatID, username + "(" + contact + "): " + message)

	if server.Db.CheckUser(chatID) {
		if contact != "" {
			server.Db.AddPhone(chatID, contact)

			msg := tgbotapi.NewMessage(chatID, "Номер телефона добавлен")
			_, _ = server.Bot.Send(msg)
			return
		}

		msg := tgbotapi.NewMessage(chatID, "Бот для отправки уведомлений")
		_, _ = server.Bot.Send(msg)
		return
	} else {
		server.Db.AddNewUser(chatID, username)

		msg := tgbotapi.NewMessage(chatID, "Предоствьте номер телефона, нажав на кнопку в меню")

		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButtonContact("Предоставить номер телефона"),
			),
		)
		_, _ = server.Bot.Send(msg)
		return
	}

}

func HandleKey(update tgbotapi.Update, server Server) {
	/*server := Server{Bot: bot, Db: db}

	chatID := update.CallbackQuery.Message.Chat.ID
	message := update.CallbackQuery.Data*/

}
