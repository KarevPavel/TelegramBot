package interfaces

import (
	"telegram-bot-long-polling/telegram/models"
)

type IBotCommandObserver interface {
	Register(Func IBotCommandFunc, observerID string)
	Unregister(observerID string)
	NotifyAll(arg BotCommandArgument)
}

type IBotCommandFunc func(BotCommandArgument)

type BotCommandArgument struct {
	Command, Argument string
	MessageId         int
	ChatId            uint64
	//TFunction         *ITelegramFunctions
	Response          *models.Update
	//Cache             Cache
}