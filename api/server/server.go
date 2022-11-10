package server

import (
    "fmt"
    dragonflyServer "github.com/df-mc/dragonfly/server"
)

var DragonflyServer *dragonflyServer.Server

func IsPlayerOnline(args ...any) []any {
    _, ok := DragonflyServer.PlayerByName(args[0].(string))
    return []any{ok}
}

func Players(args ...any) []any {
    players := DragonflyServer.Players()

    playerNames := make([]any, len(players))

    for i, player := range players {
        playerNames[i] = player.Name()
    }

    fmt.Println(playerNames)

    return playerNames
}
