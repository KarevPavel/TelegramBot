package commands

import (
	"aria2c"
	"encoding/base64"
	"github.com/asaskevich/EventBus"
	"log"
	"regexp"
	"telegram-bot-long-polling/constants"
	"telegram-bot-long-polling/external/torrentz2"
	"telegram-bot-long-polling/interfaces"
	"telegram-bot-long-polling/telegram"
	"telegram-bot-long-polling/telegram/models"
)

type commandProcessor struct {
	Cache      interfaces.Cache
	EventBus   EventBus.Bus
	TFunctions interfaces.ITelegramFunctions
	AriaApi    aria2c.AriaWSSender
}

func NewCommandProcessor(Cache interfaces.Cache, EventBus EventBus.Bus, TFunctions interfaces.ITelegramFunctions, AriaApi aria2c.AriaWSSender) *commandProcessor {
	return &commandProcessor{Cache: Cache,
		EventBus:   EventBus,
		TFunctions: TFunctions,
		AriaApi:    AriaApi,
	}
}

func (command *commandProcessor) ProcessDocument(botCommandArg interfaces.BotCommandArgument) {
	if constants.ByFile.Equals(botCommandArg.Command) {
		if botCommandArg.Response.Message != nil && botCommandArg.Response.Message.Document != nil {
			var document = botCommandArg.Response.Message.Document
			if !regexp.MustCompile(".*\\.torrent$").MatchString(document.FileName) {
				_, _ = command.TFunctions.SendMessage(models.SendMessage{
					ChatId:           botCommandArg.ChatId,
					Text:             "Wrong file format. Pattern '.*\\.torrent&'",
					ReplyToMessageId: botCommandArg.MessageId,
				})
				return
			}
			var file, _ = command.TFunctions.GetFile(document.FileId)
			var fileBytes = command.TFunctions.DownloadFile(file.FilePath)
			var b64 = base64.StdEncoding.EncodeToString(fileBytes)
			var uuid = command.AriaApi.AddTorrent(b64)
			command.Cache.Put(uuid, botCommandArg)
			command.EventBus.Subscribe(uuid, command.AriaReceived)
		}
	}
}

func (command *commandProcessor) AriaReceived(req *aria2c.Request, resp *aria2c.Response) {

	var tmp = command.Cache.Get(req.Id)
	if tmp != nil {
		var gid = resp.Result.(string)
		var botArguments = tmp.(interfaces.BotCommandArgument)
		command.TFunctions.SendMessage(models.SendMessage{
			ChatId:                botArguments.ChatId,
			Text:                  "Aria Received. Gid: " + gid,
			DisableWebPagePreview: false,
			DisableNotification:   false,
			ReplyToMessageId:      botArguments.MessageId,
			ReplyMarkup:           models.InlineKeyboardMarkup{},
		})
		command.Cache.Put(gid, botArguments)
		command.EventBus.Subscribe(gid, command.DownloadStarted)

	}
	command.EventBus.Unsubscribe(req.Id, command.AriaReceived)
}

func (command *commandProcessor) DownloadStarted(gid string) {

	var tmp = command.Cache.Get(gid)
	if tmp != nil {
		var botArguments = tmp.(interfaces.BotCommandArgument)
		command.TFunctions.SendMessage(models.SendMessage{
			ChatId:                botArguments.ChatId,
			Text:                  "Download started. Gid: " + gid,
			DisableWebPagePreview: false,
			DisableNotification:   false,
			ReplyToMessageId:      botArguments.MessageId,
			ReplyMarkup:           models.InlineKeyboardMarkup{},
		})
		command.Cache.Put(gid, botArguments)
	}
	command.EventBus.Unsubscribe(gid, command.DownloadStarted)
	command.EventBus.Subscribe(gid, command.BtDownloadCompleted)
}

func (command *commandProcessor) BtDownloadCompleted(gid string) {

	var tmp = command.Cache.Get(gid)
	if tmp != nil {
		var botArguments = tmp.(interfaces.BotCommandArgument)
		command.TFunctions.SendMessage(models.SendMessage{
			ChatId:                botArguments.ChatId,
			Text:                  "Download Completed. Gid: " + gid,
			DisableWebPagePreview: false,
			DisableNotification:   false,
			ReplyToMessageId:      botArguments.MessageId,
			ReplyMarkup:           models.InlineKeyboardMarkup{},
		})
	}
	command.EventBus.Unsubscribe(gid, command.BtDownloadCompleted)
}

func (command *commandProcessor) ProcessSearchTorrents(botCommandArg interfaces.BotCommandArgument) {
	if constants.Search.Equals(botCommandArg.Command) {
		var results = torrentz2.Search(botCommandArg.Argument)
		var poll = telegram.CreatePollFromResults(botCommandArg.ChatId, botCommandArg.MessageId, 1, results)
		var response, _ = command.TFunctions.SendPoll(poll)
		command.Cache.Put(response.Poll.Id, results)
	}
}

func (command *commandProcessor) ProcessMagnetLink(botCommandArg interfaces.BotCommandArgument) {
	if constants.ByMagnetLink.Equals(botCommandArg.Command) {
		var magnetLink = botCommandArg.Argument
		log.Println(magnetLink)
		var wrapper = make([]string, 1)
		wrapper[0] = magnetLink
		command.AriaApi.AddUri(wrapper)
	}
}
