package main

import (
    "os"
    "log"
    "errors"
    "github.com/antiquark/plugo"
    dragonflyServer "github.com/df-mc/dragonfly/server"
    jukeboxServerAPI "github.com/jukebox-mc/jukebox/api/server"
)

func main() {
    log.Println("Jukebox is preparing to juke..")

    // Create Minecraft server
    server = dragonflyServer.New(nil, nil)

    // Create the plugo server
    plugo := plugo.New("jukebox")

    // Provide actual server to functions that get called remotely
    jukeboxServerAPI.server = server

    // Expose server functions
    plugo.Expose("SetName", jukeboxServer.SetName)

    // Load plugins or if first time running, create plugins folder
    path := "plugins"
    
    _, err := os.Stat(path)

    if errors.Is(err, os.ErrNotExist) {
        os.Mkdir(path, os.ModePerm)
    }

    plugins, _ := os.ReadDir(path)

    for _, plugin := range plugins {
        plugo.StartChild(path+"/"+plugin.Name())
    }

    log.Printf("Jukebox detected %v discs", len(plugins))

    server.Start()

    for server.Accept() {
    }
}
