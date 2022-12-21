package main

import (
    "log"
    "github.com/curzodo/plugo"
    "github.com/jukebox-mc/jukebox/global"
    dfPlayer "github.com/df-mc/dragonfly/server/player"
    dfServer "github.com/df-mc/dragonfly/server"
    jukeboxServerAPI "github.com/jukebox-mc/jukebox/api/server"
    jukeboxPlayerAPI "github.com/jukebox-mc/jukebox/api/player"
    jukeboxCommandAPI "github.com/jukebox-mc/jukebox/api/command"
    jukeboxEventAPI "github.com/jukebox-mc/jukebox/api/event"
)

func main() {
    log.Println("Jukebox is preparing to juke..")

    // Create Minecraft server
    conf, err := dfServer.DefaultConfig().Config(nil)

    if err != nil {
        panic(err)
    }

    server := conf.New()
    
    // Create the plugo server
    pluginServer, _ := plugo.New("jukebox")

    // Set global variables plugo and server.
    global.PluginServer = pluginServer
    global.Server = server

    // Expose server functions
    pluginServer.Expose(jukeboxServerAPI.IsPlayerOnline)
    pluginServer.Expose(jukeboxServerAPI.Players)

    // Expose player functions
    pluginServer.Expose(jukeboxPlayerAPI.AbortBreaking)
    pluginServer.Expose(jukeboxPlayerAPI.Absorption)
    pluginServer.Expose(jukeboxPlayerAPI.AddEffect)
    pluginServer.Expose(jukeboxPlayerAPI.AddExperience)
    pluginServer.Expose(jukeboxPlayerAPI.AddFood)
    pluginServer.Expose(jukeboxPlayerAPI.AirSupply)

    // Expose event functions
    pluginServer.Expose(jukeboxEventAPI.SubscribeEvent)

    // Expose command functions
    pluginServer.Expose(jukeboxCommandAPI.AddCommand)

    pluginServer.StartChildren("plugins")
    
    server.Listen()

    for server.Accept(func(p *dfPlayer.Player) {
        p.Handle(jukeboxEventAPI.NewEventHandler(p))
    }) {}
}
