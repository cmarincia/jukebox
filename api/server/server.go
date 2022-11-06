package server

import (
    dragonfly "github.com/df-mc/dragonfly/server"
)

var server *dragonfly.Server

func SetName(args ...any) []any {
    server.SetName(args...)
    return nil
}
