package server

import (
    dragonfly "github.com/df-mc/dragonfly/server"
)

var Server *dragonfly.Server

func IsPlayerOnline(args ...any) []any {
    _, ok := Server.PlayerByName(args[0].(string))
    return []any{ok}
}

func Players(args ...any) []any {
    players := Server.Players()

    playerNames := make([]any, len(players))

    for i, player := range players {
        playerNames[i] = player.Name()
    }

    return playerNames
}
