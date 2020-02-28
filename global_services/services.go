package global_services

import (
	"aria2c"
	"github.com/asaskevich/EventBus"
	"strconv"
	"telegram-bot-long-polling/cache"
	"telegram-bot-long-polling/constants"
	"telegram-bot-long-polling/external/aria_router"
	"telegram-bot-long-polling/interfaces"
	"telegram-bot-long-polling/telegram"
	"telegram-bot-long-polling/telegram/commands"
)

var wsConn = aria2c.NewAriaWsConnector("localhost", strconv.Itoa(constants.Config.Aria2C.Port), "/jsonrpc")
var notificationHandler = aria_router.NewNotificationHandler(GlobalCache, EBus)
var responseHandler = aria_router.NewResponseHandler(GlobalCache, EBus)
var commandsCache = new(cache.TemporaryCache)

var EBus = EventBus.New()

var CommandProcessor = commands.NewCommandProcessor(commandsCache, EBus, TFunctions, AriaApi)

var AriaDaemon = aria2c.NewAriaDaemon(aria2c.AriaConfig{
	EnableRPC:              true,
	DownloadsDir:           constants.Config.Aria2C.DownloadDir,
	LogFile:                constants.Config.Aria2C.LogDir,
	MaxConcurrentDownloads: constants.Config.Aria2C.MaxConcurrentDownloads,
	MaxConnPerServer:       constants.Config.Aria2C.MaxConnectionsPerServer,
	Port:                   constants.Config.Aria2C.Port,
	RpcSecret:              constants.Config.Aria2C.Secret,
	LogLevel:               constants.Config.Aria2C.LogLevel,
}).Start()

var GlobalCache interfaces.Cache = new(cache.TemporaryCache)

//Configure Aria2C API
var AriaCache interfaces.Cache = new(cache.TemporaryCache)
var router = aria_router.NewWSRouter(AriaCache, responseHandler, notificationHandler)
var AriaApi = wsConn.ConnectAndRoute(aria_router.NewLocalAriaWS(AriaCache, constants.Config.Aria2C.Secret), router)



var TFunctions = telegram.NewTFunctions(constants.Config.Client.RequestURL, constants.Config.Client.RequestFile)

var TelegramBot = telegram.NewBot(constants.Config.Client.RequestFile, constants.Config.Client.RequestFile, TFunctions)

func GlobalServicesStop() {
	_ = AriaApi.Disconnect()
	_ = AriaDaemon.Process.Kill()
}
