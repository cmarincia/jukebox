package command

import (
    "github.com/df-mc/dragonfly/server/cmd"
    "github.com/antiquark/plugo"
)

var Plugo *plugo.Plugo

func AddCommand(args ...any) []any {
    ch := commandHandler {
        args[0].(string),
        args[1].(string),
    }

    cmd.Register(cmd.New(args[2].(string), args[3].(string), nil, &ch))
    return nil
}

type commandHandler struct {
    plugoName string // Name of the plugo that called AddCommand
    functionName string // Name of the remote function to call
}

func (ch commandHandler) Run(src cmd.Source, o *cmd.Output) {
    Plugo.Call(ch.plugoName, ch.functionName, src.Name())
}
