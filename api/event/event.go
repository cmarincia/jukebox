package event

import (
	"github.com/antiquark/plugo"
	// "github.com/go-gl/mathgl/mgl64"
	dragonflyServer "github.com/df-mc/dragonfly/server"
	dragonflyWorld "github.com/df-mc/dragonfly/server/entity"
	dragonflyEvent "github.com/df-mc/dragonfly/server/event"
	dragonflyPlayer "github.com/df-mc/dragonfly/server/player"
)

var Plugo *plugo.Plugo

var DragonflyServer *dragonflyServer.Server

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
	AttackEntity = iota
	Chat
)

type EventHandler struct {
	dragonflyPlayer.NopHandler
	Player *dragonflyPlayer.Player
}

func NewEventHandler(p *dragonflyPlayer.Player) *EventHandler {
	return &EventHandler{
		dragonflyPlayer.NopHandler{},
		p,
	}
}

// Allows plugins to subscribe to events
func SubscribeEvent(args ...any) []any {
	plugoName := args[0].(string)
	handlerName := args[1].(string)
	event := args[2].(int)
	level := args[3].(int)

	eventSubs[event][level][plugoName] = append(eventSubs[event][level][plugoName], handlerName)

	return nil
}

func (e *EventHandler) HandleAttackEntity(ctx *dragonflyEvent.Context, entity dragonflyWorld.Entity, force, height *float64, critical *bool) {
	for level := 0; level < 5; level++ {
		for plugoName, functionNames := range eventSubs[Chat][level] {
			for _, functionName := range functionNames {
                var resp []any

                // Include entity type and entity name (player username if entity type is player
                resp, _ = Plugo.Call(plugoName, functionName, e.Player.Name(), *force, *height, *critical, entity.EncodeEntity(), entity.Name())

                // If length of resp is one, then it is a boolean indicating whether or not this event should be cancelled
                if len(resp) == 1 {
                    shouldCancel = resp[0].(bool)

                    if shouldCancel {
                        ctx.Cancel()
                        return
                    } else {
                        continue
                    }
                }

                *force = resp[0].(float64)
                *height = resp[1].(float64)
                *critical = resp[2].(bool)
			}
		}
	}
}

func (e *EventHandler) HandleChat(ctx *dragonflyEvent.Context, message *string) {
	for level := 0; level < 5; level++ {
		for plugoName, functionNames := range eventSubs[Chat][level] {
			for _, functionName := range functionNames {
				resp, _ := Plugo.Call(plugoName, functionName, e.Player.Name(), *message)

				// If length of string returned from plugin is zero, then cancel
				if len(resp[0].(string)) == 0 {
					ctx.Cancel()
					return
				}

				*message = resp[0].(string)
			}
		}
	}
}
