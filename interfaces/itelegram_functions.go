package interfaces

import (
	models2 "telegram-bot-long-polling/telegram/models"
)

type ITelegramFunctions interface {
	GetMe() (models2.User, error)
	GetUpdates(query models2.GetUpdates, response chan []models2.Update)
	SendMessage(request models2.SendMessage) (models2.Message, error)
	SendPoll(poll models2.SendPoll) (models2.Message, error)
	GetFile(fileId string) (models2.File, error)
	DownloadFile(filePath string) []byte
}
