package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	"github.com/matrix-org/gomatrix"
)

var (
	matrixHomeserver = mustEnv("MATRIX_HOMESERVER")
	matrixUsername   = mustEnv("MATRIX_USERNAME")
	matrixPassword   = mustEnv("MATRIX_PASSWORD")
)

func mustEnv(key string) string {
	result := os.Getenv(key)

	if result == "" {
		log.Fatalf("%s is not defined", key)
	}

	return result
}

func main() {
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

		if body[0] == '.' {
			// handle commands
			parts := strings.Fields(body)
			verb := strings.ToLower(parts[0][1:])

			switch verb {
			case "httptest":
				loc := parts[1]

				resp, err := http.Get(loc)
				if err != nil {
					cli.SendText(ev.RoomID, err.Error())
					return
				}

				cli.SendText(ev.RoomID, "GET "+loc+": "+resp.Status)
			}
		}

		switch body {
		case "h":
			cli.SendText(ev.RoomID, "h")
		}
	})

	// Blocking version
	if err := cli.Sync(); err != nil {
		fmt.Println("Sync() returned ", err)
	}
}
