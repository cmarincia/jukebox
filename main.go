package main

import (
    "os"
    "log"
    "errors"
    "github.com/antiquark/plugo"
    dragonflyServer "github.com/df-mc/dragonfly/server"
    dragonflyPlayer "github.com/df-mc/dragonfly/server/player"
    jukeboxServerAPI "github.com/jukebox-mc/jukebox/api/server"
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

    // Provide actual server to functions that get called remotely
    jukeboxServerAPI.Server = server

    // Expose server functions
    plugo.Expose("IsPlayerOnline", jukeboxServerAPI.IsPlayerOnline)
    plugo.Expose("Players", jukeboxServerAPI.Players)

    // Load plugins or if first time running, create plugins folder
    path := "plugins"
    
    _, err = os.Stat(path)

    if errors.Is(err, os.ErrNotExist) {
        os.Mkdir(path, os.ModePerm)
    }

    plugins, _ := os.ReadDir(path)

    for _, plugin := range plugins {
        plugo.StartChild(path+"/"+plugin.Name())
    }

    log.Printf("Jukebox detected %v discs", len(plugins))

    server.Listen()

    for server.Accept(func(_ *dragonflyPlayer.Player) {

    }) {
    }
}
