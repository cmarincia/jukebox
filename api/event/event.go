package event

import (
	// "github.com/go-gl/mathgl/mgl64"
    "github.com/jukebox-mc/jukebox/global"
	dfEvent "github.com/df-mc/dragonfly/server/event"
	dfPlayer "github.com/df-mc/dragonfly/server/player"
)

// event(priority(plugoName(functionNames)))
var eventSubs = make(map[int]map[int]map[string][]string)

// Intiailise maps
func init() {
	for i := 0; i < 5; i++ {
		eventSubs[i] = make(map[int]map[string][]string)
		for j := 0; j < 5; j++ {
			eventSubs[i][j] = make(map[string][]string)
		}
	}
}

const (
	Chat = iota
)

type EventHandler struct {
	dfPlayer.NopHandler
	Player *dfPlayer.Player
}

func NewEventHandler(p *dfPlayer.Player) *EventHandler {
	return &EventHandler{
		dfPlayer.NopHandler{},
		p,
	}
}

// Allows plugins to subscribe to events
func SubscribeEvent(plugoId, handlerId string, event, level int) {
	eventSubs[event][level][plugoId] = append(
        eventSubs[event][level][plugoId], 
        handlerId,
    )
}

func (e *EventHandler) HandleChat(ctx *dfEvent.Context, message *string) {
	for level := 0; level < 5; level++ {
		for pluginId, functionIds := range eventSubs[Chat][level] {
			for _, functionId := range functionIds {
				resp, _ := global.PluginServer.Call(
                    pluginId, 
                    functionId, 
                    e.Player.Name(), 
                    *message,
                )

                newMessage := resp[0].(string)
                cancel := resp[1].(bool)

				if cancel {
					ctx.Cancel()
					return
				}

				*message = newMessage
			}
		}
	}
}
