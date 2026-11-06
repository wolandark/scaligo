package keys

import (
	"github.com/gdamore/tcell/v2"
	"gs/scales"
	"time"
)

var input string
var title string
var timer *time.Timer
var scale []string

func ReadKeys(s tcell.Screen) ([]string, string, bool) {
	switch ev := s.PollEvent().(type) {
	case *tcell.EventKey:
		if ev.Rune() == 'q' {
			return nil, "", false
		}

		input += string(ev.Rune())

		if timer != nil {
			timer.Stop()
		}

		timer = time.AfterFunc(500*time.Millisecond, func() {
			s.PostEvent(tcell.NewEventInterrupt(nil))
		})

	case *tcell.EventInterrupt:
		switch input {

		// C
		case "C", "cM":
			scale, _, title = scales.GetScale("c_maj")
			input = ""
		case "cm":
			scale, _, title = scales.GetScale("c_min")
			input = ""
		case "C#":
			scale, _, title = scales.GetScale("c_sharp_maj")
			input = ""
		case "c#m":
			scale, _, title = scales.GetScale("c_sharp_min")
			input = ""

			// D
		case "D", "dM":
			scale, _, title = scales.GetScale("d_maj")
			input = ""
		case "dm":
			scale, _, title = scales.GetScale("d_min")
			input = ""
		case "D#":
			scale, _, title = scales.GetScale("d_sharp_maj")
			input = ""
		case "d#m":
			scale, _, title = scales.GetScale("d_sharp_min")
			input = ""

			// E
		case "E", "eM":
			scale, _, title = scales.GetScale("e_maj")
			input = ""
		case "em":
			scale, _, title = scales.GetScale("e_min")
			input = ""

			// F
		case "F", "fM":
			scale, _, title = scales.GetScale("f_maj")
			input = ""
		case "fm":
			scale, _, title = scales.GetScale("f_min")
			input = ""
		case "F#":
			scale, _, title = scales.GetScale("f_sharp_maj")
			input = ""
		case "f#m":
			scale, _, title = scales.GetScale("f_sharp_min")
			input = ""

			// G
		case "G", "gM":
			scale, _, title = scales.GetScale("g_maj")
			input = ""
		case "gm":
			scale, _, title = scales.GetScale("g_min")
			input = ""
		case "G#":
			scale, _, title = scales.GetScale("g_sharp_maj")
			input = ""
		case "g#m":
			scale, _, title = scales.GetScale("g_sharp_min")
			input = ""

			// A
		case "A", "aM":
			scale, _, title = scales.GetScale("a_maj")
			input = ""
		case "am":
			scale, _, title = scales.GetScale("a_min")
			input = ""
		case "A#":
			scale, _, title = scales.GetScale("a_sharp_maj")
			input = ""
		case "a#m":
			scale, _, title = scales.GetScale("a_sharp_min")
			input = ""

			// B
		case "B", "bM":
			scale, _, title = scales.GetScale("b_maj")
			input = ""
		case "bm":
			scale, _, title = scales.GetScale("b_min")
			input = ""
		}

		if title != "" {
			return scale, title, true
		}

	case *tcell.EventResize:
		s.Sync()
	}
	return nil, "", true
}
