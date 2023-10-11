package controllers

import (
	"TgDbMai/internal/psql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Диспетчер", "dispatcher"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Сотрудник дпс", "dps_employee"),
	),
)

type Command struct {
	Name string
	f    func(args ...interface{})
}
type CommandHandler struct {
	commands []Command
}

func (h *CommandHandler) FindCommand(msg *tgbotapi.Message) *Command {
	for _, command := range h.commands {
		if command.Name == msg.Command() {
			return &command
		}
	}
	return nil
}
func (h *CommandHandler) Handle(msg *tgbotapi.Message) {
	command := h.FindCommand(msg)
	command.f()

}

func (h *CommandHandler) AddHandler(name string, f func(args ...interface{})) {
	command := Command{
		Name: name,
		f:    f,
	}
	h.commands = append(h.commands, command)
}
func Controller(bot *tgbotapi.BotAPI, repository psql.Repository) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

			if callback.Text == "dispatcher" {

				msg.Text = "Выберите"
				msg.ReplyMarkup = dispacherActions

			} else if callback.Text == "call_patrul" {
				//repository.
			}
			// And finally, send a message containing the data received.
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.Message.IsCommand() {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, " ")

			// Extract the command from the Message.
			switch update.Message.Command() {
			case "start":
				msg.ReplyMarkup = numericKeyboard
				msg.Text = "Выберите вашу должность"
			case "sayhi":
				msg.Text = "Hi :)"
			case "status":
				msg.Text = "I'm ok."
			default:
				msg.Text = "I don't know that command"
			}

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)

			}

		} else if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = numericKeyboard
			case "close":
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			}

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}
