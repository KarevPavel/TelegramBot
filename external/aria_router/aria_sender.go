package aria_router

import (
	"bitbucket.org/y4cxp543/aria2c"
	"bitbucket.org/y4cxp543/telegram-bot/interfaces"
	"bitbucket.org/y4cxp543/telegram-bot/util"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type AuthAriaWS struct {
	Conn              *websocket.Conn
	Cache             interfaces.Cache
	handlerRegistered bool
	Token             string
}

func NewLocalAriaWS(cache interfaces.Cache, token string) aria2c.AriaWSSender {
	return &AuthAriaWS{handlerRegistered: false, Cache: cache, Token: token}
}

func (ws *AuthAriaWS) SetConnection(conn *websocket.Conn){
	ws.Conn = conn
}

func (ws *AuthAriaWS) Disconnect() error {
	return ws.Conn.Close()
}

func (ws *AuthAriaWS) Router(router aria2c.Router) {
	if ws.handlerRegistered {
		log.Print("Another handler already registered")
		return
	}
	ws.handlerRegistered = true
	go func() {
		for {
			var response = new(aria2c.Response)
			_, p, err := ws.Conn.ReadMessage()
			if err != nil {
				router.HandleError(err)
				return
			}
			log.Println("Aria2c Response: ", string(p))
			err = json.Unmarshal(p, &response)
			if err != nil {
				router.HandleError(err)
				return
			}
			router.Route(response)
		}
	}()
}

func (ws *AuthAriaWS) execute(method string, param interface{}) string {
	var request = aria2c.NewRequest(ws.Token, method, util.Guid(), param)
	var byteArr,_ = json.MarshalIndent(request, "", "    ")
	log.Println("Aria2c Request: ", string(byteArr))
	ws.Cache.Put(request.Id, request)
	_ = ws.Conn.WriteJSON(request)
	return request.Id
}

func (ws *AuthAriaWS) AddUri(uri []string) string {
	return ws.execute(aria2c.AddUri, uri)
}

func (ws *AuthAriaWS) AddTorrent(b64 string) string {
	return ws.execute(aria2c.AddTorrent, b64)
}

func (ws *AuthAriaWS) GetPeers() string {
	return ws.execute(aria2c.GetPeers, nil)
}

func (ws *AuthAriaWS) AddMetalink(metalink string) string {
	return ws.execute(aria2c.GetPeers, metalink)
}

func (ws *AuthAriaWS) Remove(guid string) string {
	return ws.execute(aria2c.Remove, guid)
}

func (ws *AuthAriaWS) Pause(guid string) string {
	return ws.execute(aria2c.Pause, guid)
}

func (ws *AuthAriaWS) ForcePause(guid string) string {
	return ws.execute(aria2c.ForcePause, guid)
}

func (ws *AuthAriaWS) PauseAll() string {
	return ws.execute(aria2c.PauseAll, nil)
}

func (ws *AuthAriaWS) ForcePauseAll() string {
	return ws.execute(aria2c.ForcePauseAll, nil)
}

func (ws *AuthAriaWS) Unpause() string {
	return ws.execute(aria2c.Unpause, nil)
}

func (ws *AuthAriaWS) UnpauseAll() string {
	return ws.execute(aria2c.UnpauseAll, nil)
}

func (ws *AuthAriaWS) ForceRemove() string {
	return ws.execute(aria2c.ForceRemove, nil)
}

func (ws *AuthAriaWS) ChangePosition() string {
	return ws.execute(aria2c.ChangePosition, nil)
}

func (ws *AuthAriaWS) TellStatus() string {
	return ws.execute(aria2c.TellStatus, nil)
}

func (ws *AuthAriaWS) GetUris() string {
	return ws.execute(aria2c.GetUris, nil)
}

func (ws *AuthAriaWS) GetFiles() string {
	return ws.execute(aria2c.GetFiles, nil)
}

func (ws *AuthAriaWS) GetServers() string {
	return ws.execute(aria2c.GetServers, nil)
}

func (ws *AuthAriaWS) TellActive() string {
	return ws.execute(aria2c.TellActive, nil)
}

func (ws *AuthAriaWS) TellWaiting() string {
	return ws.execute(aria2c.TellWaiting, nil)
}

func (ws *AuthAriaWS) TellStopped() string {
	return ws.execute(aria2c.TellStopped, nil)
}

func (ws *AuthAriaWS) GetOption() string {
	return ws.execute(aria2c.GetOption, nil)
}

func (ws *AuthAriaWS) ChangeUri() string {
	return ws.execute(aria2c.ChangeUri, nil)
}

func (ws *AuthAriaWS) ChangeOption() string {
	return ws.execute(aria2c.ChangeOption, nil)
}

func (ws *AuthAriaWS) GetGlobalOption() string {
	return ws.execute(aria2c.GetGlobalOption, nil)
}

func (ws *AuthAriaWS) ChangeGlobalOption() string {
	return ws.execute(aria2c.ChangeGlobalOption, nil)
}

func (ws *AuthAriaWS) PurgeDownloadResult() string {
	return ws.execute(aria2c.PurgeDownloadResult, nil)
}

func (ws *AuthAriaWS) RemoveDownloadResult() string {
	return ws.execute(aria2c.RemoveDownloadResult, nil)
}

func (ws *AuthAriaWS) GetVersion() string {
	return ws.execute(aria2c.GetVersion, nil)
}

func (ws *AuthAriaWS) GetSessionInfo() string {
	return ws.execute(aria2c.GetSessionInfo, nil)
}

func (ws *AuthAriaWS) Shutdown() string {
	return ws.execute(aria2c.Shutdown, nil)
}

func (ws *AuthAriaWS) ForceShutdown() string {
	return ws.execute(aria2c.ForceShutdown, nil)
}

func (ws *AuthAriaWS) GetGlobalStat() string {
	return ws.execute(aria2c.GetGlobalStat, nil)
}

func (ws *AuthAriaWS) SaveSession() string {
	return ws.execute(aria2c.SaveSession, nil)
}

func (ws *AuthAriaWS) Multicall() string {
	return ws.execute(aria2c.Multicall, nil)
}

func (ws *AuthAriaWS) ListMethods() string {
	return ws.execute(aria2c.ListMethods, nil)
}

func (ws *AuthAriaWS) ListNotifications() string {
	return ws.execute(aria2c.ListNotifications, nil)
}
