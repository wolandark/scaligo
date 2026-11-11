package keys

import (
	"scaligo/scales"
	"time"

	"github.com/gdamore/tcell/v2"
)

var input string
var timer *time.Timer

var byAbbrev = invert(scales.ScaleMap)

func invert(m map[string][]string) map[string]string {
	out := make(map[string]string)
	for full, abbrevs := range m {
		for _, a := range abbrevs {
			out[a] = full
		}
	}
	return out
}

func ReadKeys(s tcell.Screen) ([]string, string, bool, bool, bool) {
	switch ev := s.PollEvent().(type) {
	case *tcell.EventKey:
		if ev.Rune() == 'q' || ev.Key() == tcell.KeyCtrlC {
			return nil, "", false, false, false
		}

		if ev.Rune() == 'h' {
			return nil, "", true, false, true
		}

		if ev.Rune() == 's' {
			return nil, "", false, true, true
		}

		input += string(ev.Rune())

		if timer != nil {
			timer.Stop()
		}

		timer = time.AfterFunc(1000*time.Millisecond, func() {
			s.PostEvent(tcell.NewEventInterrupt(nil))
		})

	case *tcell.EventInterrupt:
		full, ok := byAbbrev[input]
		input = ""
		if ok {
			scale, _, title := scales.GetScale(full)
			return scale, title, false, false, true
		}

	case *tcell.EventResize:
		s.Sync()
	}
	return nil, "", false, false, true
}
