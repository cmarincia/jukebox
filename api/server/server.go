package server

import (
    "github.com/jukebox-mc/jukebox/global"
)

var server = global.Server

func IsPlayerOnline(playerName string) bool {
    _, ok := server.PlayerByName(playerName)
    return ok
}

func Players() []string {
    players := server.Players()

    playerNames := make([]string, len(players))

    for i, player := range players {
        playerNames[i] = player.Name()
    }

    return playerNames
}
