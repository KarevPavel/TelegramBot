package observer

import (
	"bitbucket.org/y4cxp543/telegram-bot/constants"
	"bitbucket.org/y4cxp543/telegram-bot/interfaces"
)


type SystemObserver struct {
	handlers map[string][]map[string]interfaces.ICommonObserverFunc
}

type observerConfig struct {
	Func       interfaces.ICommonObserverFunc
	ObserverID string
	GroupID    constants.GroupID
}

func newObserverConfig(Func interfaces.ICommonObserverFunc, observerID string, groupID constants.GroupID) observerConfig {
	return observerConfig{
		Func:       Func,
		ObserverID: observerID,
		GroupID:    groupID,
	}
}

func (o *SystemObserver) Register(Func interfaces.ICommonObserverFunc, observerID string, groupID constants.GroupID) {
	var cfg = newObserverConfig(Func, observerID, groupID)

	if o.handlers == nil {
		o.handlers = make(map[string][]map[string]interfaces.ICommonObserverFunc)
	}
	var tmpMap = make(map[string]interfaces.ICommonObserverFunc)
	tmpMap[cfg.ObserverID] = cfg.Func
	o.handlers[string(cfg.GroupID)] = append(o.handlers[string(cfg.GroupID)], tmpMap)
}

func (o *SystemObserver) Unregister(GroupID, ObserverID string){
	var mapList = o.handlers[GroupID]
	for index, _map := range mapList {
		var handler = _map[ObserverID]
		if handler != nil {
			if index == 1 || index + 1 == len(mapList) {
				if index == 1 {
					mapList = mapList[1:]
				} else {
					mapList = mapList[:len(mapList) - 1]
				}
			} else {
				mapList = append(mapList[:index], mapList[index:]...)
			}
			if len(o.handlers[GroupID]) == 0 {
				delete(o.handlers, GroupID)
			}
		}
	}
	o.handlers[GroupID] = mapList
}

func (o *SystemObserver) NotifyAll(GroupID constants.GroupID, params map[string]interface{}){
	var mapList = o.handlers[string(GroupID)]
	for _, _map := range mapList {
		for _, function := range _map {
			function(params)
		}
	}
}