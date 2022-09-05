package telegram

import (
	"bitbucket.org/y4cxp543/telegram-bot/cache"
	"bitbucket.org/y4cxp543/telegram-bot/constants"
	"bitbucket.org/y4cxp543/telegram-bot/external/torrentz2"
	"bitbucket.org/y4cxp543/telegram-bot/interfaces"
	"bitbucket.org/y4cxp543/telegram-bot/observer"
	"bitbucket.org/y4cxp543/telegram-bot/telegram/models"
	"bitbucket.org/y4cxp543/telegram-bot/util"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//TelegramBot wrapper for delegating HTTP calls
type Bot struct {
	url                 string
	fileRequestUrl      string
	cache               interfaces.Cache
	systemObserver      interfaces.IObserver
	botCommandsObserver interfaces.IBotCommandObserver
	tFunctions          interfaces.ITelegramFunctions
}

func NewBot(url, fileRequest string, iFunc interfaces.ITelegramFunctions) Bot {
	var telegramBot = Bot{
		url:                 url,
		fileRequestUrl:      fileRequest,
		systemObserver:      new(observer.SystemObserver),
		botCommandsObserver: new(observer.BotCommandObserver),
		cache:               new(cache.TemporaryCache),
		tFunctions:          iFunc,
	}
	telegramBot.systemObserver.Register(telegramBot.processPoll, "processPoll", constants.UpdateResponse)
	telegramBot.systemObserver.Register(telegramBot.processUpdateResponses, "processUpdateResponses", constants.UpdateResponse)
	return telegramBot
}

func parseCommand(text string) (string, string){
	regex := regexp.MustCompile("^/([a-zA-z]+)")
	match := regex.FindStringSubmatch(text)
	fullMatch := match[0]
	command := match[1]
	if len(text) < len(fullMatch) {
		return constants.EmptyString, constants.EmptyString
	}
	if len(text) == len(fullMatch) {
		return command, constants.EmptyString
	}
	arguments := text[len(fullMatch) + 1:]
	return command, arguments
}

func (t Bot) processPoll(paramWrapper map[string]interface{}) {
	if len(paramWrapper) == 1 {
		var updateResponses = paramWrapper[string(constants.Response)].([]models.Update)
		for _, upd := range updateResponses {
			if upd.PollAnswer != nil {
				//IF answer is "next page"
				if upd.PollAnswer.OptionIds[0] == constants.NextPagePollIndex {
					var cachedVal = t.cache.Get(upd.PollAnswer.PollId)
					if cachedVal != nil {
						var results = cachedVal.([]torrentz2.Result)
						var chatId = paramWrapper[string(constants.ChatId)].(uint64)
						var messageId = paramWrapper[string(constants.MessageId)].(int)
						CreatePollFromResults(chatId, messageId, 2, results)
					}
				}
			}
		}
	}
}


func CreatePollFromResults(chatId uint64, messageId, pageNumber int, results []torrentz2.Result) models.SendPoll {
	var pollTextSb = new(strings.Builder)
	var needAddNextPageOption = true
	var length = constants.NextPagePollIndex
	var optionsSize = constants.TelegramMaxPollSize
	if len(results) <= constants.TelegramMaxPollSize {
		length = len(results)
		optionsSize = length
		needAddNextPageOption = false
	}
	if pageNumber != 1 {
		results = results[constants.NextPagePollIndex * pageNumber:]
	}
	var options = make([]string, optionsSize)
	for index, r := range results {
		if index == length {
			break
		}
		var resultText = util.AddSpacesBetweenStrings(r.Name, r.Size, r.Age)
		var totalLength = len(resultText)
		if totalLength > constants.TelegramMaxPollTextSize {
			r.Name = r.Name[:len(r.Name)-(totalLength - constants.TelegramMaxPollTextSize) - len(constants.TreeDots)] + constants.TreeDots
			resultText = util.AddSpacesBetweenStrings(r.Name, r.Size, r.Age)
		}
		options[index] = resultText
	}
	if needAddNextPageOption {
		options[constants.NextPagePollIndex] = "Next page"
		var pageCount = math.Ceil(float64(len(results) / constants.NextPagePollIndex))
		pollTextSb.WriteString("Page " + strconv.Itoa(pageNumber) + " from " + strconv.Itoa(int(pageCount)))
	}

	return models.SendPoll{
		ChatId:              chatId,
		Question:            pollTextSb.String(),
		Options:             options,
		IsAnonymous:         false,
		IsClosed:            false,
		DisableNotification: false,
		ReplyToMessageId:    messageId,
	}
}


func (t Bot) processUpdateResponses(paramWrapper map[string]interface{}) {
	if len(paramWrapper) == 1 {
		var updateResponses = paramWrapper[string(constants.Response)].([]models.Update)
		for _, upd := range updateResponses {
			if upd.Message != nil && len(upd.Message.Entities) > 0 {
				if constants.BotCommand.Equals(upd.Message.Entities[0].Type) {
					var command, argument = parseCommand(upd.Message.Text)
					var botCommandArgument = t.createBotCommandArgument(command, argument, upd)
					t.botCommandsObserver.NotifyAll(botCommandArgument)
				}
			}
			if upd.Message != nil && len(upd.Message.CaptionEntities) > 0 {
				if constants.BotCommand.Equals(upd.Message.CaptionEntities[0].Type) {
					var command, argument = parseCommand(upd.Message.Caption)
					var botCommandArgument = t.createBotCommandArgument(command, argument, upd)
					t.botCommandsObserver.NotifyAll(botCommandArgument)
				}
			}
		}
	}
}

func (t *Bot) createBotCommandArgument(command, argument string, upd models.Update) interfaces.BotCommandArgument {
	return interfaces.BotCommandArgument{
		Command:   command,
		Argument:  argument,
		MessageId: upd.Message.MessageId,
		ChatId:    upd.Message.Chat.Id,
		//TFunction: &t.tFunctions,
		Response:  &upd,
		//Cache:	   t.cache,
	}
}

func nextOffset(response []models.Update) int {
	answer := 0
	if response != nil && len(response) > 0 {
		for _, upd := range response {
			if answer < upd.UpdateId {
				answer = upd.UpdateId + 1
			}
		}
	}
	return answer
}

func (t Bot) Start() {
	//offset - это идентификатор последнего сообщения.
	//т.е. если не указать то прийдут вообще все сообщения.
	//а offset это нижняя граница ид сообщения.
	var offset = 908895178
	for {
		ch := make(chan []models.Update)
		go t.tFunctions.GetUpdates(models.GetUpdates{
			Offset:  offset,
			Limit:   0,
			Timeout: 30,
		}, ch)
		var response = <-ch
		offset = nextOffset(response)
		var wrapper = make(map[string]interface{})
		wrapper[string(constants.Response)] = response
		t.systemObserver.NotifyAll(constants.UpdateResponse, wrapper)
	}
}

/**
Each Handler will receive parameters map witch contains keys:
see BotCommandVariables
"command"   - Command name
"argument"  - Command arguments
"chatId"    - Current chat id
"messageId" - Current message id
"response"  - Update Response
"tFunction" - Telegram API
"cache" 	- TemporaryCache to store any data
 */
func (t Bot) RegisterBotCommand(torrents interfaces.IBotCommandFunc, observerId string) {
	t.botCommandsObserver.Register(torrents, observerId)
}