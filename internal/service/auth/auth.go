package auth

import (
	"TgDbMai/internal/keyboard"
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	Anyone = iota
	Gashnik
	Gai
)

type Auth interface {
	LogIn(chatid int64, userId, accessLvl int)
	LogOut(chatid int64)
	IsLogged(chatid int64) bool
	GetUser(chatid int64) *User
	GetAccess(chatid int64) int
	GetTgIdsByPersonId(personId ...int) map[int]int64
	AuthMiddleware(privateCommand ...string) func(next tgbotapi.HandlerFunc) tgbotapi.HandlerFunc
}

type AuthMap struct {
	DB map[int64]*User
}
type User struct {
	PersonId    int
	OfficerId   int
	accessLevel int
}

func (auth *AuthMap) LogIn(chatid int64, userId, accessLvl int) {
	auth.DB[chatid] = &User{PersonId: userId, accessLevel: accessLvl}

}
func (auth *AuthMap) LogOut(chatid int64) {

	delete(auth.DB, chatid)
}

func (auth *AuthMap) IsLogged(chatid int64) bool {
	_, ok := auth.DB[chatid]

	return ok
}
func (auth *AuthMap) GetUser(chatid int64) *User {
	return auth.DB[chatid]
}
func (auth *AuthMap) GetAccess(chatid int64) int {
	return auth.DB[chatid].accessLevel
}
func (auth *AuthMap) GetTgIdsByPersonId(personId ...int) map[int]int64 {
	res := make(map[int]int64)
	for _, id := range personId {
		res[id] = 0
	}
	for tgId, user := range auth.DB {
		_, ok := res[user.PersonId]
		if ok {
			res[user.PersonId] = tgId
		}

	}
	return res
}

func (auth *AuthMap) AuthMiddleware(privateCommand ...string) func(next tgbotapi.HandlerFunc) tgbotapi.HandlerFunc {
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
