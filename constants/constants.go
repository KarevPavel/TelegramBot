package constants

import (
	"bitbucket.org/y4cxp543/telegram-bot/external/utils"
	"bitbucket.org/y4cxp543/telegram-bot/models"
)

var Config = utils.ReadConfig(ConfigurationFile, &models.Conf{}).(*models.Conf)

/**************************************
	   SIMPLE CONSTANTS
***************************************/
const ConfigurationFile string = "config.toml"

const JSONContentType string = "application/json"
const Method string = "method"
const Space string = " "
const EmptyString string = ""
const BaseUrl string = "https://www.torrentz.eu.com"
const SearchOrderByPeers = BaseUrl + "/search?f="
const TelegramMaxPollSize int = 10
const NextPagePollIndex = TelegramMaxPollSize - 1
const TelegramMaxPollTextSize int = 100
const TreeDots = "..."


/**************************************
   BotCommands STRUCTURE
***************************************/
type BotCommands string

const (
	ByFile       BotCommands = "byFile"
	Search       BotCommands = "search"
	ByMagnetLink BotCommands = "byMagnet"
)

func (b BotCommands) Equals(string2 string) bool {
	if string(b) == string2 {
		return true
	}
	return false
}

/**************************************
   BotCommandVariables ID STRUCTURE
***************************************/
type BotCommandVariables string

const (
	Command   BotCommandVariables = "command"
	Argument  BotCommandVariables = "argument"
	MessageId BotCommandVariables = "messageId"
	ChatId    BotCommandVariables = "chatId"
	Response  BotCommandVariables = "response"
	TFunction BotCommandVariables = "tFunction"
)

/**************************************
   GROUP ID STRUCTURE
***************************************/
type GroupID string

const (
	BotCommand     GroupID = "bot_command"
	UpdateResponse GroupID = "UpdateResponse"
)

func (b GroupID) Equals(groupId string) bool {
	if string(b) == groupId {
		return true
	}
	return false
}

/**************************************
   TELEGRAM IMPLEMENTED METHODS
***************************************/
type TelegramMethods string

const (
	GetMe       TelegramMethods = "getMe"
	GetUpdates  TelegramMethods = "getUpdates"
	SendMessage TelegramMethods = "sendMessage"
	SendPoll    TelegramMethods = "sendPoll"
	GetFile     TelegramMethods = "getFile"
)

func (b TelegramMethods) String() string {
	return string(b)
}

/**************************************
		QUERY PARAMS
***************************************/
type QueryParams string

const (
	Offset         QueryParams = "offset"
	Timeout        QueryParams = "timeout"
	Limit          QueryParams = "limit"
	AllowedUpdates             = "allowed_updates"
)
