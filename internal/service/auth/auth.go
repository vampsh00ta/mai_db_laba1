package auth

import (
	"TgDbMai/internal/keyboard"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	Anyone = iota
	Gashnik
	Gai
)

type Auth struct {
	DB map[int64]*User
}
type User struct {
	Id          int
	accessLevel int
}

func (auth *Auth) LogIn(userTgId int64, userId int, accessLvl int) {
	auth.DB[userTgId] = &User{Id: userId, accessLevel: accessLvl}

}
func (auth *Auth) LogOut(userTgId int64) {

	delete(auth.DB, userTgId)
}

func (auth *Auth) IsLogged(userTgId int64) bool {
	_, ok := auth.DB[userTgId]
	fmt.Println("daun")

	fmt.Println(ok)
	return ok
}
func (auth *Auth) GetUser(userTgId int64) *User {
	return auth.DB[userTgId]
}
func (auth *Auth) GetAccess(userTgId int64) int {
	return auth.DB[userTgId].accessLevel
}

func (auth *Auth) AuthMiddleware(privateCommand ...string) func(next tgbotapi.HandlerFunc) tgbotapi.HandlerFunc {
	allCommands := make(map[string]int)

	allCommands[keyboard.AddParticipantDtpCommand] = Gashnik
	allCommands[keyboard.CheckVehicleCommand] = Gashnik

	return func(next tgbotapi.HandlerFunc) tgbotapi.HandlerFunc {
		return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
			msg := update.Message.Text
			me, err := bot.GetMe(ctx)
			userTgId := me.ID
			if err != nil {
				return
			}
			for command, access := range allCommands {
				if msg == command && (!auth.IsLogged(userTgId) || auth.GetAccess(userTgId) < access) {
					bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
						ChatID: update.Message.Chat.ID,
						Text:   "Ввойдите в аккаунт",
					})
					bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
						ChatID:      update.Message.Chat.ID,
						ReplyMarkup: keyboard.Main(),
					})
					return
				}
			}

			next(ctx, bot, update)
		}
	}
}
