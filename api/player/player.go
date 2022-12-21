package player

import (
    "time"
    "errors"
    "github.com/jukebox-mc/jukebox/global"
    dfEffect "github.com/df-mc/dragonfly/server/entity/effect"
)

var server = global.Server

func AbortBreaking(username string) {
    p, ok := server.PlayerByName(username)

    if !ok { return }

    p.AbortBreaking()
}

func Absorption(username string) (float64, error) {
    p, ok := server.PlayerByName(username)

    if !ok {
        var empty float64
        return empty, errors.New(
            "Player with username " + username + " is not online.",
        )
    }

    return p.Absorption(), nil
}

func AddEffect(username, effectName string, strength, duration int) {
    p, ok := server.PlayerByName(username)

    if !ok {
        return
    }

    var effect dfEffect.Effect

    switch effectName {
        case "absorption":
            effect = dfEffect.New(
                dfEffect.Absorption{},
                strength,
                time.Duration(duration) * time.Millisecond,
            )
        case "blindness":
            effect = dfEffect.New(
                dfEffect.Blindness{},
                strength,
                time.Duration(duration) * time.Millisecond,
            )
        case "conduit_power":
            effect = dfEffect.New(
                dfEffect.ConduitPower{},
                strength,
                time.Duration(duration) * time.Millisecond,
            )
        case "darkness":
            effect = dfEffect.New(
                dfEffect.Darkness{},
                strength,
                time.Duration(duration) * time.Millisecond,
            )
        case "fatal_poison":
            effect = dfEffect.New(
                dfEffect.FatalPoison{},
                strength,
                time.Duration(duration) * time.Millisecond,
            )
        default:
            return
    }
    
    p.AddEffect(effect)
    return
}

func AddExperience(username string, levels int) {
    player, ok := server.PlayerByName(username)

    if !ok {
        return
    }

    player.AddExperience(levels)
}

func AddFood(username string, points int) {
    p, ok := server.PlayerByName(username)

    if !ok {
        return
    }

    p.AddFood(points)
}

func AirSupply(username string) int {
    p, ok := server.PlayerByName(username)

    if !ok {
        return -1
    }

    return int(p.AirSupply())
}
