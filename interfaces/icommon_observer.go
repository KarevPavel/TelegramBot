package interfaces

import "telegram-bot-long-polling/constants"

type IObserver interface {
	Register(Func ICommonObserverFunc, observerID string, groupID constants.GroupID)
	Unregister(GroupID, ObserverID string)
	NotifyAll(command constants.GroupID, wrapper map[string]interface{})
}

type ICommonObserverFunc func(params map[string]interface{})
