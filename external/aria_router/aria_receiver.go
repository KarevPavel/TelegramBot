package aria_router

import (
	"bitbucket.org/y4cxp543/aria2c"
	"bitbucket.org/y4cxp543/telegram-bot/constants"
	"bitbucket.org/y4cxp543/telegram-bot/interfaces"
	"bitbucket.org/y4cxp543/telegram-bot/util"
	"encoding/json"
	"github.com/asaskevich/EventBus"
	"log"
)

type wsRouter struct {
	cache               interfaces.Cache
	respHandler         aria2c.AriaWSResponseHandler
	notificationHandler aria2c.AriaWSNotificationHandler
}

func NewWSRouter(cache interfaces.Cache, respHandler aria2c.AriaWSResponseHandler, notificationHandler aria2c.AriaWSNotificationHandler) *wsRouter {
	return &wsRouter{
		cache:               cache,
		respHandler:         respHandler,
		notificationHandler: notificationHandler,
	}
}

func (router *wsRouter) Route(resp *aria2c.Response) {
	if resp.Error != nil {
		router.HandleAriaError(resp)
		return
	}

	//Notification
	if resp.Method != constants.EmptyString && resp.Params != nil {
		var notificationHandler = aria2c.NotificationFuncMap[resp.Method]
		notificationHandler(router.notificationHandler, resp)
	} else { //Standard answer
		var tmp = router.cache.Get(resp.Id)
		if tmp != nil {
			var request = tmp.(*aria2c.Request)
			var receiveFunc = aria2c.MethodNameFuncMap[request.Method]
			receiveFunc(router.respHandler, request,  resp)
		}
	}
}

func (router *wsRouter) HandleError(err error) {
	log.Print("Oops, here is error: ", err)
}

func (router *wsRouter) HandleAriaError(resp *aria2c.Response) {
	var byteArr, _ = json.MarshalIndent(resp, "", "    ")
	log.Println("WTF?! Error received! Response:\n", string(byteArr))
}

type NotificationHandler struct{
	Cache interfaces.Cache
	EventBus EventBus.Bus
}

func NewNotificationHandler(Cache interfaces.Cache, EventBus EventBus.Bus) aria2c.AriaWSNotificationHandler {
	return &NotificationHandler{
		Cache:    Cache,
		EventBus: EventBus,
	}
}

func (n *NotificationHandler) OnDownloadStart(resp *aria2c.Response) {
	var gid = util.GetGid(resp)
	n.EventBus.Publish(gid)
}

func (n *NotificationHandler) OnDownloadPause(resp *aria2c.Response) {
	log.Println("OnDownloadPause: ", resp)
}

func (n *NotificationHandler) OnDownloadStop(resp *aria2c.Response) {
	log.Println("OnDownloadStop: ", resp)
}

func (n *NotificationHandler) OnDownloadComplete(resp *aria2c.Response) {
	log.Println("OnDownloadComplete: ", resp)
}

func (n *NotificationHandler) OnDownloadError(resp *aria2c.Response) {
	log.Println("OnDownloadError: ", resp)
}

func (n *NotificationHandler) OnBtDownloadComplete(resp *aria2c.Response) {
	var gid = util.GetGid(resp)
	n.EventBus.Publish(gid)
}

type ResponseHandler struct {
	Cache    interfaces.Cache
	EventBus EventBus.Bus
}

func NewResponseHandler(Cache interfaces.Cache, EventBus EventBus.Bus) aria2c.AriaWSResponseHandler {
	return &ResponseHandler {
		Cache: Cache,
		EventBus: EventBus,
	}
}

func (w ResponseHandler) ReceiveAddUri(req *aria2c.Request, resp *aria2c.Response) {
	log.Println(resp)
}

func (w ResponseHandler) ReceiveAddTorrent(req *aria2c.Request, resp *aria2c.Response) {
	w.EventBus.Publish(req.Id, req, resp)
}
func (w ResponseHandler) ReceiveGetPeers(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveAddMetalink(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveRemove(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceivePause(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveForcePause(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceivePauseAll(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveForcePauseAll(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveUnpause(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveUnpauseAll(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveForceRemove(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveChangePosition(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveTellStatus(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveGetUris(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveGetFiles(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveGetServers(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveTellActive(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveTellWaiting(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveTellStopped(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveGetOption(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveChangeUri(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveChangeOption(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveGetGlobalOption(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveChangeGlobalOption(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceivePurgeDownloadResult(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveRemoveDownloadResult(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveGetVersion(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveGetSessionInfo(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveShutdown(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveForceShutdown(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveGetGlobalStat(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveSaveSession(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveMulticall(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveListMethods(req *aria2c.Request, resp *aria2c.Response) {
}
func (w ResponseHandler) ReceiveListNotifications(req *aria2c.Request, resp *aria2c.Response) {
	log.Println(resp.Id)
}
