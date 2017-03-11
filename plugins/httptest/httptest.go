package main

import (
	"net/http"

	"github.com/Xe/xj9/bot"
	"github.com/matrix-org/gomatrix"
)

type httpTestCmd struct{}

func (h httpTestCmd) Verb() string {
	return "httptest"
}

func (h httpTestCmd) Helptext() string {
	return "a simple test for HTTP URLs"
}

func (h httpTestCmd) Handler(c *gomatrix.Client, ev *gomatrix.Event, parv []string) error {
	if len(parv) == 2 {
		loc := parv[1]

		resp, err := http.Get(loc)
		if err != nil {
			c.SendText(ev.RoomID, err.Error())
			return err
		}

		c.SendText(ev.RoomID, "GET "+loc+": "+resp.Status)

		return nil
	}

	return bot.ErrParvCountMismatch
}

func (h httpTestCmd) Permissions(c *gomatrix.Client, ev *gomatrix.Event, parv []string) error {
	return bot.NoPermissions(c, ev, parv)
}

func init() {
	bot.DefaultCommandSet.Add(&httpTestCmd{})
}

func Okay() bool {
	return true
}
