package player

import (
    "time"
    dragonflyServer "github.com/df-mc/dragonfly/server"
    dragonflyEffect "github.com/df-mc/dragonfly/server/entity/effect"
)

var DragonflyServer *dragonflyServer.Server

func AbortBreaking(args ...any) []any {
    dragonflyPlayer, ok := DragonflyServer.PlayerByName(args[0].(string))

    if !ok {
        return nil
    }

    dragonflyPlayer.AbortBreaking()
    return nil
}

func Absorption(args ...any) []any {
    dragonflyPlayer, ok := DragonflyServer.PlayerByName(args[0].(string))

    if !ok {
        return nil
    }

    return []any{dragonflyPlayer.Absorption()}
}

func AddEffect(args ...any) []any {
    dragonflyPlayer, ok := DragonflyServer.PlayerByName(args[0].(string))

    if !ok {
        return nil
    }

    strength := args[2].(int)
    duration := args[3].(int)

    var effect dragonflyEffect.Effect

    switch args[1] {
        case "absorption":
            effect = dragonflyEffect.New(
                dragonflyEffect.Absorption{},
                strength,
                time.Duration(duration) * time.Millisecond,
            )
        case "blindness":
            effect = dragonflyEffect.New(
                dragonflyEffect.Blindness{},
                strength,
                time.Duration(duration) * time.Millisecond,
            )
        case "conduitpower":
            effect = dragonflyEffect.New(
                dragonflyEffect.ConduitPower{},
                strength,
                time.Duration(duration) * time.Millisecond,
            )
        case "darkness":
            effect = dragonflyEffect.New(
                dragonflyEffect.Darkness{},
                strength,
                time.Duration(duration) * time.Millisecond,
            )
        case "fatalpoison":
            effect = dragonflyEffect.New(
                dragonflyEffect.FatalPoison{},
                strength,
                time.Duration(duration) * time.Millisecond,
            )
        default:
    }
    
    dragonflyPlayer.AddEffect(effect)
    return nil
}

func AddExperience(args ...any) []any {
    dragonflyPlayer, ok := DragonflyServer.PlayerByName(args[0].(string))

    if !ok {
        return nil
    }

    dragonflyPlayer.AddExperience(args[1].(int))
    return nil
}

func AddFood(args ...any) []any {
    dragonflyPlayer, ok := DragonflyServer.PlayerByName(args[0].(string))

    if !ok {
        return nil
    }

    dragonflyPlayer.AddFood(args[1].(int))
    return nil
}

func AirSupply(args ...any) []any {
    dragonflyPlayer, ok := DragonflyServer.PlayerByName(args[0].(string))

    if !ok {
        return nil
    }

    return []any{int(dragonflyPlayer.AirSupply())}
}
