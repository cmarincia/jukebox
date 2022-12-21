package server

import (
    "github.com/jukebox-mc/jukebox/global"
)

func IsPlayerOnline(playerName string) bool {
    _, ok := global.Server.PlayerByName(playerName)
    return ok
}

func Players() []string {
    players := global.Server.Players()

    playerNames := make([]string, len(players))

    for i, player := range players {
        playerNames[i] = player.Name()
    }

    return playerNames
}
