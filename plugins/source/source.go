package main

import (
	"github.com/Xe/xj9/bot"
	"github.com/matrix-org/gomatrix"
)

func source(c *gomatrix.Client, m *gomatrix.Event, parv []string) error {
	c.SendText(m.RoomID, "Source code: https://github.com/Xe/xj9")
	return nil
}

func init() {
	go bot.DefaultCommandSet.AddCmd("source", "link to bot source code", bot.NoPermissions, source)
}
