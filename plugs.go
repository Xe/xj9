package main

import (
	"errors"
	"log"
	"os"
	"plugin"

	"github.com/Xe/xj9/bot"
	"github.com/matrix-org/gomatrix"
)

var errNotAdmin = errors.New("plugs: you're not an admin, sorry.")

func perms(c *gomatrix.Client, ev *gomatrix.Event, parv []string) error {
	if ev.Sender != os.Getenv("ADMIN") {
		return errNotAdmin
	}

	return nil
}

func command(c *gomatrix.Client, ev *gomatrix.Event, parv []string) error {
	if len(parv) != 2 {
		return bot.ErrParvCountMismatch
	}

	name := parv[1]

	_, err := plugin.Open("./plugins/" + name + ".so")
	if err != nil {
		return err
	}

	log.Printf("%s loaded %s", ev.Sender, name)
	c.SendText(ev.RoomID, "loaded "+name)

	return nil
}

func init() {
	go bot.DefaultCommandSet.AddCmd("plugload", "loads a plugin (admin only)", perms, command)
}
