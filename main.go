package main

import (
    "os"
    "log"
    "errors"
    "github.com/antiquark/plugo"
    dragonflyServer "github.com/df-mc/dragonfly/server"
    dragonflyPlayer "github.com/df-mc/dragonfly/server/player"
    jukeboxServerAPI "github.com/jukebox-mc/jukebox/api/server"
    jukeboxPlayerAPI "github.com/jukebox-mc/jukebox/api/player"
    jukeboxCommandAPI "github.com/jukebox-mc/jukebox/api/command"
    jukeboxEventAPI "github.com/jukebox-mc/jukebox/api/event"
)

func main() {
    log.Println("Jukebox is preparing to juke..")

    // Create Minecraft server
    conf, err := dragonflyServer.DefaultConfig().Config(nil)

    if err != nil {
        panic(err)
    }

    server := conf.New()

    // Create the plugo server
    plugo := plugo.New("jukebox")

    log.Println(plugo.Ln.Addr())

    // Provide dragonfly server to functions that get called remotely
    jukeboxServerAPI.DragonflyServer = server
    jukeboxPlayerAPI.DragonflyServer = server
    jukeboxEventAPI.DragonflyServer = server

    // Provide plugo to functions that get called remotely
    jukeboxCommandAPI.Plugo = &plugo
    jukeboxEventAPI.Plugo = &plugo

    // Expose server functions
    plugo.Expose("IsPlayerOnline", jukeboxServerAPI.IsPlayerOnline)
    plugo.Expose("Players", jukeboxServerAPI.Players)

    // Expose player functions
    plugo.Expose("AbortBreaking", jukeboxPlayerAPI.AbortBreaking)
    plugo.Expose("Absorption", jukeboxPlayerAPI.Absorption)
    plugo.Expose("AddEffect", jukeboxPlayerAPI.AddEffect)
    plugo.Expose("AddExperience", jukeboxPlayerAPI.AddExperience)
    plugo.Expose("AddFood", jukeboxPlayerAPI.AddFood)
    plugo.Expose("AirSupply", jukeboxPlayerAPI.AirSupply)

    // Expose event functions
    plugo.Expose("SubscribeEvent", jukeboxEventAPI.SubscribeEvent)

    // Expose command functions
    plugo.Expose("AddCommand", jukeboxCommandAPI.AddCommand)

    // Load plugins or if first time running, create plugins folder
    path := "plugins"
    
    _, err = os.Stat(path)

    if errors.Is(err, os.ErrNotExist) {
        os.Mkdir(path, os.ModePerm)
    }

    plugins, _ := os.ReadDir(path)

    for _, plugin := range plugins {
        plugo.StartChild(path+"/"+plugin.Name())
        plugo.Call(plugin.Name(), "onEnable")
    }

    log.Printf("Jukebox detected %v discs", len(plugins))

    server.Listen()

    for server.Accept(func(p *dragonflyPlayer.Player) {
        p.Handle(jukeboxEventAPI.NewEventHandler(p))
    }) {}
}
