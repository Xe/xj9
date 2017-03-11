package bot

import (
	"fmt"

	"github.com/matrix-org/gomatrix"
)

func (cs *CommandSet) help(s *gomatrix.Client, m *gomatrix.Event, parv []string) error {
	switch len(parv) {
	case 1:
		// print all help on all commands
		result := "Bot commands: \n"

		for verb, cmd := range cs.cmds {
			result += fmt.Sprintf("%s%s: %s\n", cs.Prefix, verb, cmd.Helptext())
		}

		result += "If there's any problems please don't hesitate to ask someone for help."

		_, err := s.SendText(m.RoomID, result)
		return err

	default:
		return ErrParvCountMismatch
	}
}
