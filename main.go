package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Xe/xj9/common"
	_ "github.com/joho/godotenv/autoload"
	"github.com/matrix-org/gomatrix"
	"github.com/sbinet/plugs"
	"github.com/tchap/go-exchange/exchange"
)

var (
	matrixHomeserver = mustEnv("MATRIX_HOMESERVER")
	matrixUsername   = mustEnv("MATRIX_USERNAME")
	matrixPassword   = mustEnv("MATRIX_PASSWORD")
	adminName        = mustEnv("ADMIN")
	autoloadPlugins  = mustEnv("AUTOLOAD")
)

func mustEnv(key string) string {
	result := os.Getenv(key)

	if result == "" {
		log.Fatalf("%s is not defined", key)
	}

	return result
}

func main() {
	for _, name := range strings.Split(autoloadPlugins, ",") {
		_, err := plugs.Open("github.com/Xe/xj9/plugins/" + name)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("loaded %s", name)
	}

	cli, _ := gomatrix.NewClient(matrixHomeserver, "", "")
	resp, err := cli.Login(&gomatrix.ReqLogin{
		Type:     "m.login.password",
		User:     matrixUsername,
		Password: matrixPassword,
	})
	if err != nil {
		panic(err)
	}
	cli.SetCredentials(resp.UserID, resp.AccessToken)

	syncer := cli.Syncer.(*gomatrix.DefaultSyncer)
	syncer.OnEventType("m.room.message", func(ev *gomatrix.Event) {
		body, ok := ev.Body()
		if !ok {
			return
		}

		log.Printf("[%s] %s: %s", ev.RoomID, ev.Sender, body)

		if ev.Sender == cli.UserID {
			return
		}

		err := common.Exc.Publish(exchange.Topic("incoming"), common.Event{
			Client: cli,
			Event:  ev,
		})
		if err != nil {
			log.Printf("failed to publish message: %#v: %v", ev, err)
			return
		}
	})

	// Blocking version
	if err := cli.Sync(); err != nil {
		fmt.Println("Sync() returned ", err)
	}
}
