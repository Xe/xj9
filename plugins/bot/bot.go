package main

import (
	"log"

	"github.com/Xe/xj9/bot"
	"github.com/Xe/xj9/common"
	"github.com/tchap/go-exchange/exchange"
)

func init() {
	common.Exc.Subscribe(exchange.Topic("incoming"), func(topic exchange.Topic, event exchange.Event) {
		ev, ok := event.(common.Event)
		if !ok {
			return
		}

		err := bot.DefaultCommandSet.Run(ev.Client, ev.Event)
		if err != nil {
			log.Printf("while running %v: %v", ev.Event, err)
			return
		}
	})
}
