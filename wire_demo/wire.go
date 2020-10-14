//+build wireinject

package main

import (
	"fmt"
	"github.com/google/wire"
)

func InitEvent() Event {
	build := wire.Build(NewEvent, NewGreeter, NewMessage)
	fmt.Println(build)
	return Event{}
}
