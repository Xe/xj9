package main

import (
	"github.com/Xe/xj9/common"
	"github.com/tchap/go-exchange/exchange"
)

func init() {
	common.Exc.Subscribe(exchange.Topic("incoming"), func(topic exchange.Topic, event exchange.Event) {
		cev, ok := event.(common.Event)
		if !ok {
			return
		}

		c := cev.Client
		ev := cev.Event

		body, ok := ev.Body()
		if !ok {
			return
		}

		if body == "h" {
			c.SendText(ev.RoomID, "h")
		}
	})
}
