package observer

import (
	"telegram-bot-long-polling/interfaces"
)

type BotCommandObserver struct {
	handlers map[string]interfaces.IBotCommandFunc
}

func (o *BotCommandObserver) Register(Func interfaces.IBotCommandFunc, observerID string){
	if o.handlers == nil {
		o.handlers = make(map[string]interfaces.IBotCommandFunc)
	}
	o.handlers[observerID] = Func
}

func (o *BotCommandObserver) Unregister(observerID string){
	delete(o.handlers, observerID)
}

func (o *BotCommandObserver) NotifyAll(arg interfaces.BotCommandArgument) {
	for _,function := range o.handlers {
		function(arg)
	}
}