package global_services

import (
	"bitbucket.org/y4cxp543/aria2c"
	"bitbucket.org/y4cxp543/telegram-bot/cache"
	"bitbucket.org/y4cxp543/telegram-bot/constants"
	"bitbucket.org/y4cxp543/telegram-bot/telegram"
	"bitbucket.org/y4cxp543/telegram-bot/telegram/commands"
	"github.com/asaskevich/EventBus"
)

/*var wsConn = aria2c.NewAriaWsConnector("localhost", strconv.Itoa(constants.Config.Aria2C.Port), "/jsonrpc")
var notificationHandler = aria_router.NewNotificationHandler(GlobalCache, EBus)
var responseHandler = aria_router.NewResponseHandler(GlobalCache, EBus)*/
var commandsCache = new(cache.TemporaryCache)

var EBus = EventBus.New()

var CommandProcessor = commands.NewCommandProcessor(commandsCache, EBus, TFunctions)

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

/*var GlobalCache interfaces.Cache = new(cache.TemporaryCache)*/

//Configure Aria2C API
/*var AriaCache interfaces.Cache = new(cache.TemporaryCache)
var router = aria_router.NewWSRouter(AriaCache, responseHandler, notificationHandler)
var AriaApi = wsConn.ConnectAndRoute(aria_router.NewLocalAriaWS(AriaCache, constants.Config.Aria2C.Secret), router)*/



var TFunctions = telegram.NewTFunctions(constants.Config.Client.RequestURL, constants.Config.Client.RequestFile)

var TelegramBot = telegram.NewBot(constants.Config.Client.RequestFile, constants.Config.Client.RequestFile, TFunctions)

func GlobalServicesStop() {
	/*_ = AriaApi.Disconnect()*/
	/*_ = AriaDaemon.Process.Kill()*/
}
