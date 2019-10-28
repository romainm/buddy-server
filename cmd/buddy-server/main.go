package main

import (
    "log"
    "github.com/romainm/buddy-server/pkg/server"
    "github.com/romainm/buddy-server/internal/orm"
)

func main() {
    orm, err := orm.Factory()
    if err != nil {
        log.Panic(err)
    }
    server.Run(orm)
}

